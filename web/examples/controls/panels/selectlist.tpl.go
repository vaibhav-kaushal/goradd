//** This file was code generated by got. DO NOT EDIT. ***

package panels

import (
	"bytes"
	"context"
)

func (ctrl *SelectListPanel) DrawTemplate(ctx context.Context, buf *bytes.Buffer) (err error) {

	buf.WriteString(`
<h1>Selection Lists</h1>
<p>
Selection lists let you select a zero, one, or multiple items from a group of items.
Below are the supported selection lists
that represent what is available through standard html controls. Many css/javascript frameworks use
these standard html widgets as the basis for their more elaborate selection mechanisms.
</p>
<p>
Selection lists all use the ItemList mixin, so you have a single kind of interface to add labels and
associated values to the list. Individual items can be styled as well.
</p>
<p>
These lists expect to have all of the items inserted and styled before being drawn. If you change items
after they are drawn, in response to a button click for example, then be sure to call Refresh() on the control
to redraw the list of items to select from.
While some of the list widgets can allow items to scroll in and out of view,
all the items in a list are rendered to html (In other words, the html is sent to the browser, but
items that are scrolled out of view are simply hidden from the user by the browser).
There is no current mechanism for paging in or filtering items not shown.
Therefore, these selection lists are not great for selecting from a very large list of items.
</p>
<p>
Also see the types of columns available in the html table control which can serve as selections from a list.
Tables are better at managing a very large list of items to select from, as the items in a
table can be paged in, and can also be filtered.
</p>
<h2>Single Selection Lists</h2>
<p>
These lists let the user select zero or one item from a list. If you call SetRequired(true), then
the user will be forced to select an item.
</p>
<h3>SelectList</h3>
<p>
A SelectList is a typical dropdown list with a single selection. You can also set the size attribute
to display it as a scrolling list rather than a dropdown list.
</p>
`)

	buf.WriteString(`
`)
	if `` == "" {
		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("singleSelectList-ff").Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	} else {

		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("singleSelectList-ff").ProcessAttributeString(``).Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	}

	buf.WriteString(`
`)

	buf.WriteString(`
`)
	if `` == "" {
		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("selectListWithSize-ff").Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	} else {

		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("selectListWithSize-ff").ProcessAttributeString(``).Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	}

	buf.WriteString(`
<h3>RadioList</h3>
<p>
A radio list presents a list of radio buttons in a table format. The Value of the control is the value
of the item the user selects.
</p>
<p>
RadioLists can be scrolling lists if you wish. To see the scrolling effect, you have to somehow limit the
size of the control itself using CSS.
</p>

`)

	buf.WriteString(`
`)
	if `` == "" {
		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("radioList1-ff").Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	} else {

		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("radioList1-ff").ProcessAttributeString(``).Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	}

	buf.WriteString(`
`)

	buf.WriteString(`
`)
	if `` == "" {
		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("radioList2-ff").Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	} else {

		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("radioList2-ff").ProcessAttributeString(``).Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	}

	buf.WriteString(`
`)

	buf.WriteString(`
`)
	if `` == "" {
		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("radioList3-ff").Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	} else {

		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("radioList3-ff").ProcessAttributeString(``).Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	}

	buf.WriteString(`

<h2>Multiple Selection Lists</h2>
<p>
These lists let the user select zero or multiple items from a list. If you call SetRequired(true), then
the user will be forced to select at least one item.
</p>
<h3>HTML Multiple Selection List</h3>
<p>
`)

	buf.WriteString(`This list allows multiple selections if you hold down the shift key or the command key. While this is
a standard html widget, the look and behavior of this widget often depends on the browser used and
the operating system, so its not a great widget for multiple selections. It is here for completeness, and
also in case a javascript widget can build a better interface on top of it. Also, notice that it
sets the &#34;required&#34; attribute, which lets the browser check to see if an item is selected. But again,
the user interface is dependent on the browser, and sometimes it is very confusing to a user. You can
turn off browser validity checking by setting the &#34;novalidate&#34; attribute on the form.
`)

	buf.WriteString(`
</p>
`)

	buf.WriteString(`
`)
	if `` == "" {
		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("multiselectList-ff").Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	} else {

		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("multiselectList-ff").ProcessAttributeString(``).Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	}

	buf.WriteString(`
<h3>Checkbox List</h3>
<p>
This list is similar to the radio list above, but uses checkboxes so you can select multiple
items. It is a much better way to present to the user a selection of multiple items, and is the default
control that gets generated by the code generator to let you edit one-to-many relationships between
data objects.
</p>
`)

	buf.WriteString(`
`)
	if `` == "" {
		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("checklist1-ff").Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	} else {

		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("checklist1-ff").ProcessAttributeString(``).Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	}

	buf.WriteString(`

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
`)

	return
}
