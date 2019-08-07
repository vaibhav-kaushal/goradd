//** This file was code generated by got. DO NOT EDIT. ***

package browsertest

import (
	"bytes"
	"context"
)

func (form *TestForm) AddHeadTags() {
	form.FormBase.AddHeadTags()
	if "Tests" != "" {
		form.Page().SetTitle("Tests")
	}

	// double up to deal with body attributes if they exist
	form.Page().BodyAttributes = ``
}

func (form *TestForm) DrawTemplate(ctx context.Context, buf *bytes.Buffer) (err error) {

	buf.WriteString(`
<h1>Browser Based Tests</h1>
`)

	buf.WriteString(`
`)

	{
		err := form.TestList.Draw(ctx, buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`
`)

	buf.WriteString(`
`)

	{
		err := form.RunButton.Draw(ctx, buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`
`)

	buf.WriteString(`
`)

	{
		err := form.RunAllButton.Draw(ctx, buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`
<div>
Currently running:
`)

	buf.WriteString(`
`)

	{
		err := form.RunningLabel.Draw(ctx, buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`
`)

	buf.WriteString(`
`)

	{
		err := form.Controller.Draw(ctx, buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`
</div>
`)

	buf.WriteString(`
`)

	return
}
