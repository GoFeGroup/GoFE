package GoFE

import (
	"fmt"
	"testing"

	"github.com/GoFeGroup/GoFE/components"
	"github.com/GoFeGroup/GoFE/static"
)

func TestGenerateStyles(t *testing.T) {
	fmt.Println(static.GenerateStyles())
}

func TestNewButton(t *testing.T) {
	fmt.Println(components.NewButton().SetTitle("Press").Generate())
}

func TestNewTable(t *testing.T) {
	header := components.NewTableHeader().
		AddCell("th1").AddCell("th2").AddCell("th3").AddCell("th4")
	row1 := components.NewTableRow().AddCell(1).AddCell(2).AddCell(3).AddCell(4)
	row2 := components.NewTableRow().AddCell(1).AddCell(2).AddCell(3).AddCell(4).SetDataId("999988")
	row3 := components.NewTableRow().AddCell(1).AddCell(2).AddCell(3).AddCell(4)
	row4 := components.NewTableRow().AddCell(1).AddCell(2).AddCell(3).AddCell(4)
	table := components.NewTable().SetCheckbox(true).
		AddHeader(header).AddRow(row1).AddRow(row2).AddRow(row3).AddRow(row4)
	fmt.Println(table.Generate())
}
