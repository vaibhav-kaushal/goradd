// Package page is the user-interface layer of goradd, and implements state management and rendering
// of an html page, as well as the framework for rendering controls.
//
// To use the page package, you start by creating a form object, and then adding controls to that form.
// You also should add a drawing template to define additional html for the form.
package page

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"github.com/goradd/goradd/pkg/config"
	"github.com/goradd/goradd/pkg/goradd"
	"github.com/goradd/goradd/pkg/html"
	"github.com/goradd/goradd/pkg/i18n"
	reflect2 "github.com/goradd/goradd/pkg/reflect"
	"github.com/goradd/goradd/pkg/session"
	strings2 "github.com/goradd/goradd/pkg/strings"
	"reflect"
	"strconv"
	"strings"
)

// PageRenderStatus keeps track of whether we are rendering the page or not
type PageRenderStatus int

// Future note. Below is for general information but should NOT be used to synchronize multiple drawing routines.
// An architecture using channels to synchronize page changes and drawing would be better.
// For now, except for testing, we should not get in a situation where multiple copies of a form
// are being used.
const (
	PageIsNotRendering PageRenderStatus = iota // FormBase has started rendering but has not finished
	PageIsRendering
)

// PageCacheVersion helps us keep track of when a change to the application changes the pagecache format. It is only needed
// when serializing the pagecache. Some page cache stores may be difficult to invalidate the whole thing, so this lets
// lets us invalidate old pagecaches individually. Feel free to bump this as needed, though you should use
// a number after UserPageCacheVersion so there is no conflict with the goradd default.
var PageCacheVersion int32 = 1

// This value is used to generate unique ids in the control registry. However, if the control registry
// detects a collision, you will need to change this value and restart your app. If you have a running
// page cache, you should change the PageCacheVersion above as well to invalidate it.
var ControlRegistrySalt = "goradd"

// If you want to bump the page cache version yourself, you can use this as a starting point so there is no
// conflict with goradd itself
const UserPageCacheVersion = 10000


// PageDrawFunc is the type of the page drawing function. This is implemented by the page drawing template.
type PageDrawFunc func(context.Context, *Page, *bytes.Buffer) error

// DrawI is the interface for items that draw into the draw buffer
type DrawI interface {
	Draw(context.Context, *bytes.Buffer) error
}

// A code we use during serialization to indicate that we just unserialized a control id
const controlCode = "**grc**"


// The Page object is the top level drawing object, and is essentially a wrapper for the form. The Page draws the
// html, head and body tags, and includes the one Form object on the page. The page also maintains a record of all
// the controls included on the form.
type Page struct {
	// BodyAttributes contains the attributes that will be output with the body tag. It should be set before the
	// form draws, like in the AddHeadTags function.
	BodyAttributes string

	stateId      string // Id in cache of the pagestate. Needs to be output by form.
	renderStatus PageRenderStatus
	idPrefix     string // For creating unique ids for the app

	controlRegistry map[string]ControlI
	form            FormI
	idCounter       int
	title           string // page title to draw in head tag
	htmlHeaderTags  []html.VoidTag
	responseHeader  map[string]string // queues up anything to be sent in the response header
	responseError   int

	language int // Don't serialize this. This is a cached version of what the session holds.
}

// Init initializes the page. Should be called by a form just after creating Page.
func (p *Page) Init() {
}

// Restore is called immediately after the page has been unserialized, to fix up decoded controls.
func (p *Page) Restore() {
	for _,c := range p.controlRegistry {
		c.Restore()
	}
}

func (p *Page) runPage(ctx context.Context, buf *bytes.Buffer, isNew bool) (err error) {
	grCtx := GetContext(ctx)
	p.ClearResponseHeaders()

	if grCtx.err != nil {
		panic(grCtx.err) // An error occurred during unpacking of the context, so report that now
	}

	if err = p.Form().Run(ctx); err != nil {
		return err
	}

	// cache the language tags so we only need to look them up once for every call
	p.language = i18n.SetDefaultLanguage(ctx, grCtx.Header.Get("accept-language"))

	if isNew {
		p.Form().AddHeadTags()
		p.Form().CreateControls(ctx)
		p.Form().LoadControls(ctx)

	} else {
		// Test for a CSRF attack
		csrf := session.Get(ctx, goradd.SessionCsrf)
		csrf2, found := grCtx.FormValue(htmlCsrfToken)
		if !found || csrf != csrf2 {
			return fmt.Errorf("CSRF error: %s", p.stateId)
		}

		p.Form().updateValues(ctx) // Tell all the controls to update their values.
		// if this is an event response, do the actions associated with the event
		if p.HasControl(grCtx.actionControlID) {
			p.GetControl(grCtx.actionControlID).control().doAction(ctx)
		}

		// Redraw controls that requested a redraw, probably through the watcher mechanism
		for _,id := range grCtx.refreshIDs {
			if p.HasControl(id) {
				p.GetControl(id).Refresh()
			}
		}
	}

	if grCtx.RequestMode() == Ajax {
		err = p.DrawAjax(ctx, buf)
		p.SetResponseHeader("Content-Type", "application/json")
	} else if grCtx.RequestMode() == Server || grCtx.RequestMode() == Http {
		//p.url = grCtx.HttpContext.URL. We might want a record of the original URL to be used during ajax calls someday. Until we have a reason, this will remain commented out.
		err = p.Draw(ctx, buf)
	} else {
		// TODO: Implement a hook for the CustomAjax call and/or Rest API calls?
	}

	p.Form().writeAllStates(ctx)
	p.Form().Exit(ctx, err)

	pageCache.Set(p.stateId, p)

	return
}

// Returns the form for the page
func (p *Page) Form() FormI {
	return p.form
}

// Draw draws the page.
func (p *Page) Draw(ctx context.Context, buf *bytes.Buffer) (err error) {
	f := p.form.PageDrawingFunction()
	return f(ctx, p, buf)
}

// DrawHeaderTags draws all the inner html for the head tag
func (p *Page) DrawHeaderTags(ctx context.Context, buf *bytes.Buffer) {
	if p.title != "" {
		buf.WriteString("  <title>")
		buf.WriteString(p.title)
		buf.WriteString("  </title>\n")
	}

	// draw things like additional meta tags, etc
	if p.htmlHeaderTags != nil {
		for _, tag := range p.htmlHeaderTags {
			buf.WriteString(tag.Render())
		}
	}

	p.Form().DrawHeaderTags(ctx, buf)
}

// SetControlIdPrefix sets the prefix for control ids. Some javascript frameworks (i.e. jQueryMobile) require that control ids
// be unique across the application, vs just in the page, because they create internal caches of control ids. This
// allows you to set a per page prefix that will be added to all control ids to make them unique across the whole
// application. However, its up to you to make sure the names are unique per page.
func (p *Page) SetControlIdPrefix(prefix string) *Page {
	p.idPrefix = prefix
	return p
}

// GenerateControlID generates unique control ids. If you want to do your own id generation, or modifying of given ids, implement that
// in an override to the control.Init function. The given id is one that the user supplies. User provided ids and
// generated ids can be further munged by providing an id prefix through SetControlIdPrefix().
func (p *Page) GenerateControlID(id string) string {
	if id != "" {
		if strings.Contains(id, "_") {
			// underscores are used by the action system to route actions to sub items of the control.
			panic("You cannot add a control with an underscore in the name. Use a hyphen instead.")
		}
		if p.idPrefix != "" {
			if !strings.HasPrefix(id, p.idPrefix) { // subcontrols might already have this prefix
				id = p.idPrefix + id
			}
		}
		if p.HasControl(id) {
			panic(fmt.Sprintf(`A control with id "%s" is being added a second time to the page. Ids must be unique on the page.`, id))
		} else {
			return id
		}
	} else {
		var trialid string
		for trialid == "" || p.HasControl(trialid) { // checks to make sure user did not previously add a control that might match our generation pattern
			p.idCounter++
			trialid = p.idPrefix + "c" + strconv.Itoa(p.idCounter)
		}
		return trialid
	}
}

// GetControl returns the control with the given id. If not found, it panics. Use HasControl to check for existence.
func (p *Page) GetControl(id string) ControlI {
	if id == "" {
		panic("attempting to get a control with a blank id")
	}
	if p.controlRegistry == nil {
		panic("control registry is not initialized")
	}
	if c,ok := p.controlRegistry[id]; !ok {
		panic("control with id " + id + " was not found")
	} else {
		return c
	}
}

func (p *Page) HasControl(id string) bool {
	if id == "" {
		return false
	}
	_,ok := p.controlRegistry[id]
	return ok
}

// addControl adds the given control to the controlRegistry. It is called by the control code whenever a control is created.
func (p *Page) addControl(control ControlI) {
	id := control.ID()

	if id == "" {
		panic("ControlBase must have an id before being added.")
	}

	if p.controlRegistry == nil {
		p.controlRegistry = make (map[string]ControlI)
	}

	if p.HasControl(id) {
		panic("ControlBase id already exists. ControlBase must have a unique id on the page before being added.")
	}

	p.controlRegistry[id] = control

	if control.Parent() == nil {
		if f, ok := control.(FormI); ok {
			if p.form != nil {
				panic("The Form object for the page has already been set.")
			} else {
				p.form = f
			}
		} else {
			panic("Controls must have a parent.")
		}
	}
}

/* Remove?
func (p *Page) changeControlID(oldId string, newId string) {
	if p.GetControl(newId) != nil {
		panic(fmt.Errorf("this control id is already defined on the page: %s", newId))
	}
	ctrl := p.GetControl(oldId)
	p.controlRegistry.Delete(oldId)
	p.controlRegistry.Set(newId, ctrl)
}
*/

func (p *Page) removeControl(id string) {
	// Execute the javascript to remove the control from the dom if we are in ajax mode
	// TODO: Application::ExecuteSelectorFunction('#' . $objControl->getWrapperID(), 'remove');
	// TODO: Make This a direct command in the ajax renderer

	delete (p.controlRegistry,id)
}

// Title returns the content of the <title> tag that will be output in the head of the page.
func (p *Page) Title() string {
	return p.title
}

// Call SetTitle to set the content of the <title> tag to be output in the head of the page.
func (p *Page) SetTitle(title string) {
	p.title = title
}

// StateID returns the page state id. This is output by the form so that we can recover the saved state of the page
// each time we call into the application.
func (p *Page) StateID() string {
	return p.stateId
}

// DrawAjax renders the page during an ajax call. Since the page itself is already rendered, it simply hands off this
// responsibility to the form.
func (p *Page) DrawAjax(ctx context.Context, buf *bytes.Buffer) (err error) {
	err = p.Form().renderAjax(ctx, buf)
	return
}

/* Serialize and Deserialize are now called directly
// GobEncode here is implemented to intercept the GobSerializer to only encode an empty structure. We use this as part
// of our overall serialization strategy for forms. Controls still need to be registered with gob.
func (p *Page) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err = p.Serialize(enc)
	return buf.Bytes(), err
}

func (p *Page) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err = p.Deserialize(dec)
	return err
}
*/

func (p *Page) MarshalJSON() (data []byte, err error) {
	return
}

func (p *Page) UnmarshalJSON(data []byte) (err error) {
	return
}

// MarshalBinary is called by the framework to serialize the page state.
func (p *Page) MarshalBinary() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)
	if err = e.Encode(PageCacheVersion); err != nil {
		return
	}
	if err = e.Encode(p.stateId); err != nil {
		return
	}
	if err = e.Encode(p.idPrefix); err != nil {
		return
	}
	if err = e.Encode(p.title); err != nil {
		return
	}
	if err = e.Encode(p.htmlHeaderTags); err != nil {
		return
	}
	if err = e.Encode(p.BodyAttributes); err != nil {
		return
	}
	if err = e.Encode(p.form.ID()); err != nil {
		return
	}

	if err = p.encodeControlRegistry(e); err != nil {
		return
	}
	data = buf.Bytes()
	return
}

func (p *Page) encodeControlRegistry(e *gob.Encoder) (err error) {

	// We encode the control registry bottom up so that there is a high likelihood that child controls
	// will be available to parent controls when parent controls get deserialized. This make it possible
	// for us to deserialize forms and custom controls that save a pointer to a control, as long as
	// that pointer is exported.

	// make a copy of the ids
	ids := make(map[string]ControlI)
	for k, v := range p.controlRegistry {
		ids[k] = v
	}

	var l int = len(p.controlRegistry)

	if err = e.Encode(l); err != nil {
		return
	}
	p.form.RangeSelfAndAllChildren(func(ctrl ControlI) {
		_ = p.encodeControl(ctrl, e)
		delete(ids, ctrl.ID())
	})

	// encode controls not attached to the form, like dialogs
	for len(ids) != 0 {
		// process one item out of map at a time
		// we need to do it this way because these unattached items might have children, and we must
		// ensure that all children get serialized first
		for _,c := range ids {
			c.RangeSelfAndAllChildren(
				func(ctrl ControlI) {
					if _,ok := ids[ctrl.ID()]; ok { // we didn't yet process it
						_ = p.encodeControl(ctrl, e)
						delete(ids, ctrl.ID())
					}
				})
			break
		}
	}
	return
}

func (p *Page) encodeControl(ctrl ControlI, e *gob.Encoder) (err error){
	if err = e.Encode(ctrl.ID()); err != nil {
		return
	}
	if err = e.Encode(controlRegistryID(ctrl)); err != nil {
		return
	}

	if err = p.serializeControl(ctrl, e); err != nil {
		return
	}
	return
}

// Users can create exported items on their objects and they will be serialized and restored automatically
// Alternatively they can implement their own Serialize method.
func (p *Page) serializeControl(c ControlI, e Encoder) error {
	v := reflect.Indirect(reflect.ValueOf(c))
	fieldCount := v.NumField()
	_ = fieldCount
	exportedFields := reflect2.FieldValues(c)

	// convert all embedded controls to the id of the control
	for name,val := range exportedFields {
		if ctrl,ok := val.(ControlI); ok {
			exportedFields[name] = controlCode + ctrl.ID()
		}
	}
	if err := c.Serialize(e); err != nil {
		if config.Debug {
			panic("Error serializing control " + c.ID() + ": " + err.Error())
		}
		return err
	}
	if err := e.Encode(exportedFields); err != nil {
		if config.Debug {
			panic ("Error serializing exported fields of " + c.ID() + ": " + err.Error())
		}
		return err
	}
	return nil
}

func (p *Page) UnmarshalBinary(data []byte) (err error) {
	b := bytes.NewBuffer(data)
	dec := gob.NewDecoder(b)

	var pageCacheVersion int32
	if err = dec.Decode(&pageCacheVersion); err != nil {
		panic(err)
	}
	if pageCacheVersion != PageCacheVersion {
		return fmt.Errorf("stale data in cache") // This is a soft error indicating that the system should create a new page state
	}

	if err = dec.Decode(&p.stateId); err != nil {
		panic(err)

	}
	if err = dec.Decode(&p.idPrefix); err != nil {
		panic(err)
	}
	if err = dec.Decode(&p.title); err != nil {
		panic(err)
	}
	if err = dec.Decode(&p.htmlHeaderTags); err != nil {
		panic(err)
	}
	if err = dec.Decode(&p.BodyAttributes); err != nil {
		panic(err)
	}
	var formID string
	if err = dec.Decode(&formID); err != nil {
		panic(err)
	}

	if err = p.decodeControlRegistry(dec); err != nil {
		return
	}

	p.form = p.controlRegistry[formID].(FormI)
	return
}

func (p *Page) decodeControlRegistry(d *gob.Decoder) (err error) {
	p.controlRegistry = make(map[string]ControlI)
	var l int
	if err = d.Decode(&l); err != nil {
		panic(err)
	}

	for i := 0; i < l; i++ {
		if err = p.decodeControl(d); err != nil {
			return
		}
	}
	return
}

func (p *Page) decodeControl(d *gob.Decoder) (err error) {
	var id string
	var registryID uint64
	if err = d.Decode(&id); err != nil {
		panic(err)
	}
	if err = d.Decode(&registryID); err != nil {
		return
	}

	c := createRegisteredControl(registryID, p)
	p.controlRegistry[id] = c
	if err = p.deserializeControl(c, d); err != nil {
		return
	}
	return
}

func (p *Page) deserializeControl(c ControlI, d Decoder) error {
	if err := c.Deserialize(d); err != nil {
		return err
	}
	var exportedFields map[string]interface{}
	if err := d.Decode(&exportedFields); err != nil {
		return err
	}
	// Substitute embedded control ids for the actual control
	for name,val := range exportedFields {
		if s,ok := val.(string); ok && strings2.StartsWith(s, controlCode) {
			id := s[len(controlCode):]
			if ctrl, ok2 := p.controlRegistry[id]; ok2 {
				exportedFields[name] = ctrl
			}
		}
	}

	return reflect2.SetFieldValues(c, exportedFields)
}

// AddHtmlHeaderTag adds the given tag to the head section of the page.
func (p *Page) AddHtmlHeaderTag(t html.VoidTag) {
	p.htmlHeaderTags = append(p.htmlHeaderTags, t)
}

func (p *Page) HasMetaTag(name string) bool {
	for _,t := range p.htmlHeaderTags {
		if t.Tag == "meta" &&
			t.Attr["name"] == name {
			return true
		}
	}
	return false
}

// SetResponseHeader sets a value in the html response header. You generally would only need to do this if your are outputting
// custom content, like a pdf file.
func (p *Page) SetResponseHeader(key, value string) {
	if p.responseHeader == nil {
		p.responseHeader = map[string]string{}
	}
	p.responseHeader[key] = value
}

// ClearResponseHeaders removes all the current response headers.
func (p *Page) ClearResponseHeaders() {
	p.responseHeader = nil
}

// PushRedraw will cause the form to refresh in between events. This will cause the client to pull
// the ajax response. Its possible that this will happen while drawing. We avoid the race condition
// by sending the message anyways, and allowing the client to send an event back to us, essentially
// using the javascript event mechanism to synchronize us. We might get an unnecessary redraw, but
// that is not a big deal.
/*
func (p *Page) PushRedraw() {
	channel := "form-" + p.stateId
	if ws.HasChannel(channel) { // If we call this while launching a page, the channel isn't created yet, but the page is going to be drawn, so its ok.
		ws.SendMessage(channel, map[string]interface{}{"grup": true})
	} else {
		log.FrameworkDebug("Pushing redraw with no channel.")
	}
}
*/
// LanguageCode returns the language code that will be put in the lang attribute of the html tag.
// It is taken from the i18n package.
func (p *Page) LanguageCode() string {
	return i18n.CanonicalValue(p.language)
}

// Cleanup is called by the page cache when the page is removed from memory.
func (p *Page) Cleanup() {
	p.Form().RangeSelfAndAllChildren(func(ctrl ControlI) {
		ctrl.Cleanup()
	})
}

