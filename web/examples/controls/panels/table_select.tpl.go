//** This file was code generated by got. DO NOT EDIT. ***

package panels

import (
	"bytes"
	"context"
)

func (control *TableSelectPanel) DrawTemplate(ctx context.Context, buf *bytes.Buffer) (err error) {

	buf.WriteString(`
`)

	buf.WriteString(`
<style>
#scroller {
    border: 1px solid gray;
    overflow-x: hidden;
    overflow-y: scroll;
    height: 200px;
    width: 50%;
    margin-bottom: 20px;
    scroll-behavior: smooth;
}
</style>

<h1>Tables - Selectable Rows</h1>
<p>
The SelectTable is a table that is used to select an item from a list of items. It uses css and javascript to demonstrate
to the user that the table is selectable, and it will remember the selection and report back to the go code
through a RowSelected event when an item is selected, and what that item was.
</p>
<p>
The table is wrapped in a box that allows the table to scroll in the box.
The SelectTable will show the selected item at startup, and when its javascript <i>showSelectedItem</i> function
is called, which you can do from a Javascript action.
Some things to try as a demonstration of its capabilities:
</p>
<ul>
<li>Select an item, scroll the table so the selected item is not showing, and then click the Show Selected Item button.</li>
<li>Select an item, scroll the table so the selected item is not showing, and then refresh the page.</li>
</ul>
<p>
In each case, you should see the table scrolled so that the item is visible.
</p>
<div id="scroller">
`)

	buf.WriteString(`
`)

	{
		err := control.Table1.Draw(ctx, buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`
</div>

<div class="boxed" style="width:50%">
`)

	buf.WriteString(`
`)

	{
		err := control.InfoPanel.Draw(ctx, buf)
		if err != nil {
			return err
		}
	}

	buf.WriteString(`
</div>
`)

	buf.WriteString(`
`)

	{
		err := control.ShowButton.Draw(ctx, buf)
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
