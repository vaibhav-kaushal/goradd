package control

import (
	"context"
	"github.com/goradd/goradd/pkg/html"
	"github.com/goradd/goradd/pkg/page"
)

type SpanI interface {
	PanelI
}

// Span is a Goradd control that is a basic "span" wrapper. Use it to style and listen to events on a span. It
// can also be used as the basis for more advanced javascript controls.
type Span struct {
	Panel
}

func NewSpan(parent page.ControlI, id string) *Span {
	p := &Span{}
	p.Self = p
	p.Init(parent, id)
	return p
}

func (c *Span) Init(parent page.ControlI, id string) {
	c.Panel.Init(parent, id)
	c.Tag = "span"
}

func (c *Span) DrawingAttributes(ctx context.Context) html.Attributes {
	a := c.ControlBase.DrawingAttributes(ctx)
	a.SetDataAttribute("grctl", "span")
	return a
}

type SpanCreator struct {
	ID string
	Text string
	TextIsHtml bool
	Children []page.Creator
	page.ControlOptions
}

func (c SpanCreator) Create(ctx context.Context, parent page.ControlI) page.ControlI {
	ctrl := NewSpan(parent, c.ID)
	if c.Text != "" {
		ctrl.SetText(c.Text)
	}
	ctrl.SetTextIsHtml(c.TextIsHtml)
	ctrl.ApplyOptions(ctx, c.ControlOptions)
	ctrl.AddControls(ctx, c.Children...)
	return ctrl
}

// GetSpan is a convenience method to return the button with the given id from the page.
func GetSpan(c page.ControlI, id string) *Span {
	return c.Page().GetControl(id).(*Span)
}

func init() {
	page.RegisterControl(&Span{})
}