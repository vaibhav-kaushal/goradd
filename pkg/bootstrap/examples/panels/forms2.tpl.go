//** This file was code generated by got. DO NOT EDIT. ***

package panels

import (
	"bytes"
	"context"

	. "github.com/goradd/goradd/pkg/bootstrap/control"
	"github.com/goradd/goradd/pkg/page/control"
)

func (ctrl *Forms2Panel) DrawTemplate(ctx context.Context, buf *bytes.Buffer) (err error) {

	buf.WriteString(`

<h2>Inline Form Layout</h2>
<p>
This is an example of a typical form with inline labels.
</p>

`)

	buf.WriteString(`
`)
	if ctrl.Page().HasControl("nameText-fg") {
		ctrl.Page().GetControl("nameText-fg").(control.LabelAttributer).LabelAttributes().Merge(`class="col-form-label col-2"`)
	}

	buf.WriteString(`
`)

	buf.WriteString(`
`)
	if ctrl.Page().HasControl("nameText-fg") {
		ctrl.Page().GetControl("nameText-fg").(FormGroupI).InnerDivAttributes().Merge(`class="col-10"`)
	}

	buf.WriteString(`
`)

	buf.WriteString(`
`)
	if `class="form-row"` == "" {
		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("nameText-fg").Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	} else {

		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("nameText-fg").ProcessAttributeString(`class="form-row"`).Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	}

	buf.WriteString(`

`)

	buf.WriteString(`
`)
	if ctrl.Page().HasControl("streetText-fg") {
		ctrl.Page().GetControl("streetText-fg").(control.LabelAttributer).LabelAttributes().Merge(`class="col-form-label col-2"`)
	}

	buf.WriteString(`
`)

	buf.WriteString(`
`)
	if ctrl.Page().HasControl("streetText-fg") {
		ctrl.Page().GetControl("streetText-fg").(FormGroupI).InnerDivAttributes().Merge(`class="col-10"`)
	}

	buf.WriteString(`
`)

	buf.WriteString(`
`)
	if `class="form-row"` == "" {
		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("streetText-fg").Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	} else {

		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("streetText-fg").ProcessAttributeString(`class="form-row"`).Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	}

	buf.WriteString(`

`)

	buf.WriteString(`
`)

	buf.WriteString(`
<div class="form-row">
`)

	buf.WriteString(` `)

	buf.WriteString(`
<div class="col-6">
    `)

	buf.WriteString(`
`)
	if ctrl.Page().HasControl("cityText-fg") {
		ctrl.Page().GetControl("cityText-fg").(control.LabelAttributer).LabelAttributes().Merge(`class="col-form-label col-4"`)
	}

	buf.WriteString(`
    `)

	buf.WriteString(`
`)
	if ctrl.Page().HasControl("cityText-fg") {
		ctrl.Page().GetControl("cityText-fg").(FormGroupI).InnerDivAttributes().Merge(`class="col-8"`)
	}

	buf.WriteString(`
    `)

	buf.WriteString(`
`)
	if `class="form-row"` == "" {
		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("cityText-fg").Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	} else {

		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("cityText-fg").ProcessAttributeString(`class="form-row"`).Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	}

	buf.WriteString(`
</div>
<div class="col-2">
    `)

	buf.WriteString(`
`)
	if ctrl.Page().HasControl("stateText-fg") {
		ctrl.Page().GetControl("stateText-fg").(control.LabelAttributer).LabelAttributes().Merge(`class="col-form-label col-6"`)
	}

	buf.WriteString(`
    `)

	buf.WriteString(`
`)
	if ctrl.Page().HasControl("stateText-fg") {
		ctrl.Page().GetControl("stateText-fg").(FormGroupI).InnerDivAttributes().Merge(`class="col-6"`)
	}

	buf.WriteString(`
    `)

	buf.WriteString(`
`)
	if `class="form-row"` == "" {
		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("stateText-fg").Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	} else {

		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("stateText-fg").ProcessAttributeString(`class="form-row"`).Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	}

	buf.WriteString(`
</div>
<div class="col-4">
    `)

	buf.WriteString(`
`)
	if ctrl.Page().HasControl("zipText-fg") {
		ctrl.Page().GetControl("zipText-fg").(control.LabelAttributer).LabelAttributes().Merge(`class="col-form-label col-4"`)
	}

	buf.WriteString(`
    `)

	buf.WriteString(`
`)
	if ctrl.Page().HasControl("zipText-fg") {
		ctrl.Page().GetControl("zipText-fg").(FormGroupI).InnerDivAttributes().Merge(`class="col-8"`)
	}

	buf.WriteString(`
    `)

	buf.WriteString(`
`)
	if `class="form-row"` == "" {
		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("zipText-fg").Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	} else {

		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("zipText-fg").ProcessAttributeString(`class="form-row"`).Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	}

	buf.WriteString(`
</div>

`)

	buf.WriteString(`
</div>
<div class="form-row">
`)

	buf.WriteString(` `)

	buf.WriteString(`

<div class="col-3 offset-2">
`)

	buf.WriteString(`
`)
	if `` == "" {
		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("ajaxButton").Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	} else {

		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("ajaxButton").ProcessAttributeString(``).Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	}

	buf.WriteString(`
`)

	buf.WriteString(`
</div>
`)
	if "3" == "" {
		buf.WriteString(`<div class="col">
`)
	} else {

		buf.WriteString(`<div class="col-3">
`)
	}

	buf.WriteString(`
`)

	buf.WriteString(`
`)
	if `` == "" {
		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("serverButton").Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	} else {

		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("serverButton").ProcessAttributeString(``).Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	}

	buf.WriteString(`
`)

	buf.WriteString(`
</div>
`)

	buf.WriteString(`

`)

	buf.WriteString(`
</div>
`)

	buf.WriteString(` `)

	buf.WriteString(`

`)

	buf.WriteString(`
`)

	return
}