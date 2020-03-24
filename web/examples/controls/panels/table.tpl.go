//** This file was code generated by got. DO NOT EDIT. ***

package panels

import (
	"bytes"
	"context"
)

func (ctrl *TablePanel) DrawTemplate(ctx context.Context, buf *bytes.Buffer) (err error) {

	buf.WriteString(`
`)

	buf.WriteString(`
<style>
table {
  font-family: "Trebuchet MS", Arial, Helvetica, sans-serif;
  border-collapse: collapse;
  width: 100%;
}

table td, table th {
  border: 1px solid #ddd;
  padding: 8px;
}

table tr:nth-child(even){background-color: #f2f2f2;}

table tr:hover {background-color: #ddd;}

table th {
  padding-top: 12px;
  padding-bottom: 12px;
  text-align: left;
  background-color: #4CAF50;
  color: white;
}
</style>
<h1>Tables</h1>
<p>
The Table control creates html tables from various forms of data. After creating a table, you
add TableColumns to the table, which link the data to the display. Tables can also have pagers
to allow the user to page through data when it is too much to display all at once.
</p>
`)

	buf.WriteString(`
`)
	if `` == "" {
		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("pager1").Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	} else {

		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("pager1").ProcessAttributeString(``).Draw(ctx, buf)
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
			err := ctrl.Page().GetControl("table1").Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	} else {

		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("table1").ProcessAttributeString(``).Draw(ctx, buf)
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
			err := ctrl.Page().GetControl("table2").Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	} else {

		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("table2").ProcessAttributeString(``).Draw(ctx, buf)
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
