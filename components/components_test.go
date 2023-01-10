package components

import (
	"fmt"
	"html/template"
	"testing"

	"github.com/GoFeGroup/GoFE/components/tag"

	"github.com/GoFeGroup/GoFE/components/table"

	"github.com/GoFeGroup/GoFE/components/button"

	"github.com/GoFeGroup/GoFE/components/icon"
)

func TestNewButton(t *testing.T) {
	fmt.Println(button.New().SetTitle("Press").Generate())
	fmt.Println(button.New().SetTitle("Press").SetType(button.Primary).Generate())
	fmt.Println(button.New().SetTitle("Press").SetType(button.Success).Generate())
	fmt.Println(button.New().SetTitle("Press").SetType(button.Info).Generate())
	fmt.Println(button.New().SetTitle("Press").SetType(button.Warning).Generate())
	fmt.Println(button.New().SetTitle("Press").SetType(button.Danger).Generate())
}

func TestNewTable(t *testing.T) {
	header := table.NewHeader().
		AddCell("th1").AddCell("th2").AddCell("th3").AddCell("th4")
	row1 := table.NewRow().AddCell(1).AddCell(2).AddCell(3).AddCell(4)
	row2 := table.NewRow().AddCell(1).AddCell(2).AddCell(3).AddCell(4).SetDataId("999988")
	row3 := table.NewRow().AddCell(1).AddCell(2).AddCell(3).AddCell(4)
	row4 := table.NewRow().AddCell(1).AddCell(2).AddCell(3).AddCell(4)
	tab := table.New().SetCheckbox(true).
		AddHeader(header).AddRow(row1).AddRow(row2).AddRow(row3).AddRow(row4)
	fmt.Println(tab.Generate())
}

func TestNewIcon(t *testing.T) {
	fmt.Println(icon.New(icon.Home).Generate())
	fmt.Println(icon.New(icon.User).Generate())
	fmt.Println(icon.New(icon.Data).Generate())
	fmt.Println(icon.New(icon.Message).Generate())
	fmt.Println(icon.New(icon.Setting).Generate())
}

func TestNewButtonWithIcon(t *testing.T) {
	i := icon.New(icon.Refresh).SetStyle("color:red").Generate()
	btn := button.New().SetTitle(
		template.HTML(fmt.Sprintf("%v%v", i, "刷新"))).Generate()
	fmt.Println(btn)
}

func TestNewButtonWithBackground(t *testing.T) {
	fmt.Println(button.New().SetStyle("background:blue;color:red").Generate())
}

func TestNewTag(t *testing.T) {
	fmt.Println(tag.New("Tag").Generate())
	fmt.Println(tag.New("失败").SetType(tag.Primary).Generate())
	fmt.Println(tag.New("成功").SetType(tag.Success).Generate())
	fmt.Println(tag.New("警告").SetType(tag.Warning).Generate())
	fmt.Println(tag.New("失败").SetType(tag.Info).Generate())
	fmt.Println(tag.New("失败").SetType(tag.Danger).Generate())
}
