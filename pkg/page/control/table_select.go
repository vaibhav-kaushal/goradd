package control

import (
	"fmt"
	"github.com/goradd/gengen/pkg/maps"
	"github.com/goradd/goradd/pkg/config"
	"github.com/goradd/goradd/pkg/html"
	"github.com/goradd/goradd/pkg/page"
)

// PrimaryKeyer is an interface that is often implemented by model objects.
type PrimaryKeyer interface {
	PrimaryKey() string
}

type SelectTableI interface {
	TableI
}

// SelectTable is a table that is row selectable. To detect a row selection, trigger on event.RowSelected
type SelectTable struct {
	Table
	selectedID string
}

func NewSelectTable(parent page.ControlI, id string) *SelectTable {
	t := &SelectTable{}
	t.Init(t, parent, id)
	return t
}

func (t *SelectTable) Init(self page.ControlI, parent page.ControlI, id string) {
	t.Table.Init(self, parent, id)
	t.ParentForm().AddJQueryUI()
	t.ParentForm().AddJavaScriptFile(config.GoraddAssets() + "/js/goradd-scrollIntoView.js", false, nil)
	t.ParentForm().AddJavaScriptFile(config.GoraddAssets() + "/js/table-select.js", false, nil)
	t.AddClass("gr-clickable-rows")
}

func (t *SelectTable) this() SelectTableI {
	return t.Self.(SelectTableI)
}

func (t *SelectTable) GetRowAttributes(row int, data interface{}) (a *html.Attributes) {
	var id string

	if t.RowStyler() != nil {
		a = t.RowStyler().TableRowAttributes(row, data)
		id = a.Get("id") // styler might be giving us an id
	} else {
		a = html.NewAttributes()
	}

	// try to guess the id from the data
	if id == "" {
		switch obj := data.(type) {
		case IDer:
			id = obj.ID()
		case PrimaryKeyer:
			id = obj.PrimaryKey()
		case map[string]string:
			id, _ = obj["id"]
		case maps.StringGetter:
			id = obj.Get("id")
		}
	}
	if id != "" {
		// TODO: If configured, encrypt the id so its not publicly showing database ids
		a.SetDataAttribute("id", id)
		// We need an actual id for aria features
		a.SetID(t.ID() + "_" + id)
	} else {
		a.AddClass("nosel")
	}

	a.Set("role", "option")
	return a
}

func (t *SelectTable) ΩDrawingAttributes() *html.Attributes {
	a := t.Table.ΩDrawingAttributes()
	a.SetDataAttribute("grctl", "selecttable")
	a.Set("role", "listbox")
	a.SetDataAttribute("grWidget", "goradd.selectTable")
	if t.selectedID != "" {
		a.SetDataAttribute("grOptSelectedId", t.selectedID)
	}
	return a
}

func (t *SelectTable) ΩUpdateFormValues(ctx *page.Context) {
	if data := ctx.CustomControlValue(t.ID(), "selectedId"); data != nil {
		t.selectedID = fmt.Sprintf("%v", data)
	}
}

func (t *SelectTable) SelectedID() string {
	return t.selectedID
}

func (t *SelectTable) SetSelectedID(id string) {
	t.selectedID = id
	t.ExecuteWidgetFunction("option", "selectedId", id)
}

func (t *SelectTable) ΩMarshalState(m maps.Setter) {
	m.Set("selId", t.selectedID)
}

func (t *SelectTable) ΩUnmarshalState(m maps.Loader) {
	if v, ok := m.Load("selId"); ok {
		if id, ok := v.(string); ok {
			t.selectedID = id
		}
	}
}
