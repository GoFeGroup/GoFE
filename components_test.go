package GoFE

import (
	"fmt"
	"html/template"
	"testing"

	"github.com/GoFeGroup/GoFE/components/icon"

	"github.com/GoFeGroup/GoFE/components"
)

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

func TestNewIcon(t *testing.T) {
	fmt.Println(icon.NewIcon(icon.Home).Generate())
	fmt.Println(icon.NewIcon(icon.User).Generate())
	fmt.Println(icon.NewIcon(icon.Data).Generate())
	fmt.Println(icon.NewIcon(icon.Message).Generate())
	fmt.Println(icon.NewIcon(icon.Setting).Generate())
}

func TestNewButtonWithIcon(t *testing.T) {
	i := icon.NewIcon(icon.Refresh).SetStyle("color:red").Generate()
	btn := components.NewButton().SetTitle(
		template.HTML(fmt.Sprintf("%v%v", i, "刷新"))).Generate()
	fmt.Println(btn)
}

func TestNewButtonWithBackground(t *testing.T) {
	fmt.Println(components.NewButton().SetStyle("background:blue;color:red").Generate())
}
