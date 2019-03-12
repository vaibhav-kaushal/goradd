//** This file was code generated by got. DO NOT EDIT. ***

package controls

import (
	"bytes"
	"context"

	"github.com/goradd/goradd/pkg/page"
)

func (form *ControlsForm) AddHeadTags() {
	form.FormBase.AddHeadTags()
	if "Goradd Bootstrap Examples" != "" {
		form.Page().SetTitle("Goradd Bootstrap Examples")
	}

	// double up to deal with body attributes if they exist
	form.Page().BodyAttributes = `
`
}

func (form *ControlsForm) DrawTemplate(ctx context.Context, buf *bytes.Buffer) (err error) {

	buf.WriteString(`
`)
	path := page.GetContext(ctx).HttpContext.URL.Path
	buf.WriteString(`
<nav class="navbar navbar-expand-lg navbar-light bg-light">
  <a class="navbar-brand" href="#">Goradd Bootstrap Examples</a>
  <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
    <span class="navbar-toggler-icon"></span>
  </button>
  <div class="collapse navbar-collapse" id="navbarNav">
    <ul class="navbar-nav">
        <li class="nav-item"><a class="nav-link" href="`)

	buf.WriteString(path)

	buf.WriteString(`">Home</a></li>
        <li class="nav-item"><a class="nav-link" href="`)

	buf.WriteString(path)

	buf.WriteString(`?control=forms1">Standard Forms</a></li>
    </ul>
  </div>
</nav>

`)

	buf.WriteString(`
`)

	{
		err := form.detail.AddClass("container").Draw(ctx, buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`
`)

	buf.WriteString(`
`)

	return
}
