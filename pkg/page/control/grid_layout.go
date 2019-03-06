package control

import "github.com/goradd/goradd/pkg/html"

// NextItemPlacement controls whether items are layed out with the axis, or across an axis
// This only applies if you are using columns.
type NextItemPlacement int

const (
	// NextItemWithAxis puts the next item within the prevous item's container
	NextItemWithAxis NextItemPlacement = iota
	// NextItemCrossAxis puts the next item in to the next container
	NextItemCrossAxis
)

// GridLayoutBuilder is a helper that will allow a slice of items to be layed out in a table like
// pattern. It will compute the number of rows required, and then wrap the rows in
// row html, and the cells in cell html. You can have the items flow with the rows, or flow
// across the row axis. You can use this to build a table or a table-like structure.
type GridLayoutBuilder struct {
	items   []string
	columnCount int
	direction   NextItemPlacement
	rowTag string
	rowAttributes *html.Attributes
}

// Items sets the html for each item to display.
func (g *GridLayoutBuilder) Items(items []string) *GridLayoutBuilder {
	g.items = items
	return g
}

// ColumnCount sets the number of columns.
func (g *GridLayoutBuilder) ColumnCount(count int) *GridLayoutBuilder {
	g.columnCount = count
	return g
}

// Direction indicates how items are placed, whether they should fill up rows first, or fill up columns.
func (g *GridLayoutBuilder) Direction (placement NextItemPlacement) *GridLayoutBuilder {
	g.direction = placement
	return g
}



func (g *GridLayoutBuilder) RowTag (t string) *GridLayoutBuilder {
	g.rowTag = t
	return g
}

func (g *GridLayoutBuilder) RowClass (t string) *GridLayoutBuilder {
	g.getRowAttributes().SetClass(t)
	return g
}

func (g *GridLayoutBuilder) getRowAttributes() *html.Attributes {
	if g.rowAttributes == nil {
		g.rowAttributes = html.NewAttributes()
	}
	return g.rowAttributes
}


func (g *GridLayoutBuilder) Build () string {
	if len(g.items) == 0 {
		return ""
	}
	if g.rowTag == "" {
		g.rowTag = "div"
	}
	if g.columnCount == 0 {
		g.columnCount = 1
	}

	if g.direction == NextItemWithAxis {
		return g.wrapRows()
	} else {
		return g.wrapColumns()
	}
}


func (g *GridLayoutBuilder) wrapRows() string {
	var rows string
	var row string
	for i,_ := range g.items {
		row += g.items[i]
		if (i + 1) % g.columnCount == 0 {
			rows += html.RenderTag(g.rowTag, g.rowAttributes, row)
			row = ""
		}
	}
	if row != "" {
		// partial row
		rows += html.RenderTag(g.rowTag, g.rowAttributes, row)
	}
	return rows
}

func (g *GridLayoutBuilder) wrapColumns() string {
	l := len(g.items)
	rowCount := (l - 1) / g.columnCount + 1

	var row string
	var rows string

	for r := 0; r < rowCount; r++ {
		for c := 0; c < g.columnCount; c++ {
			i := c * rowCount + r
			if i < l {
				row += g.items[i]
			}
		}
		rows += html.RenderTag(g.rowTag, g.rowAttributes, row)
		row = ""
	}
	if row != "" {
		// partial row
		rows += html.RenderTag(g.rowTag, g.rowAttributes, row)
	}
	return rows
}






