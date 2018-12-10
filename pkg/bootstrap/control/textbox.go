package control

import (
	"encoding/gob"
	"github.com/spekary/goradd/pkg/html"
	"github.com/spekary/goradd/pkg/page"
	"github.com/spekary/goradd/pkg/page/control"
)

type Textbox struct {
	control.Textbox
}

func NewTextbox(parent page.ControlI, id string) *Textbox {
	t := new (Textbox)
	t.Init(t, parent, id)
	return t
}


func (t *Textbox) DrawingAttributes() *html.Attributes {
	a := t.Textbox.DrawingAttributes()
	a.AddClass("form-control")
	return a
}


func init () {
	gob.RegisterName("bootstrap.textbox", new(Textbox))
}


