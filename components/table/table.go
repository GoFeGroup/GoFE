package table

import (
	"fmt"
	"html/template"

	"github.com/GoFeGroup/GoFE/components/gen"

	"github.com/GoFeGroup/GoFE/utils"
)

type Cell any

type Header struct {
	Id   string
	Cell []Cell
}

func NewHeader() *Header {
	return &Header{
		Id: fmt.Sprintf("th-%s", utils.RandomString(8)),
	}
}

func (th *Header) AddCell(cell Cell) *Header {
	th.Cell = append(th.Cell, cell)
	return th
}

type Row struct {
	Id     string
	DataId string
	Cell   []Cell
}

func NewRow() *Row {
	return &Row{
		Id: fmt.Sprintf("tr-%s", utils.RandomString(8)),
	}
}

func (tr *Row) SetDataId(id string) *Row {
	tr.DataId = id
	return tr
}

func (tr *Row) AddCell(cell Cell) *Row {
	tr.Cell = append(tr.Cell, cell)
	return tr
}

type Table struct {
	Id       string
	Header   *Header
	Rows     []*Row
	Checkbox bool
}

func New() *Table {
	return &Table{
		Id: fmt.Sprintf("tab-%s", utils.RandomString(8)),
	}
}

func (tab *Table) SetCheckbox(checkbox bool) *Table {
	tab.Checkbox = checkbox
	return tab
}

func (tab *Table) AddHeader(header *Header) *Table {
	tab.Header = header
	return tab
}

func (tab *Table) AddRow(row *Row) *Table {
	tab.Rows = append(tab.Rows, row)
	return tab
}

func (tab *Table) template() string {
	return `
{{- $checkbox := .Checkbox -}}
{{- $id := .Id -}}

<table id="{{ $id }}" class="tab">
{{- with .Header }}
	<tr id="{{.Id}}">
		{{- if $checkbox -}} <th><input type="checkbox" id="checkbox-all-{{ $id }}"></th> {{- end }}
		{{- range .Cell }} <th>{{ . }}</th> {{- end -}}
	</tr>
{{- end }}

{{- range .Rows }}
	<tr id="{{.Id}}" {{- if .DataId }} data-id="{{.DataId}}" {{- end -}} >
		{{- if $checkbox -}} <td><input type="checkbox"></td> {{- end }}
		{{- range .Cell }} <td>{{.}}</td> {{- end -}}
	</tr>
{{- end }}
</table>

{{- if $checkbox }}
<script>
  $("#checkbox-all-{{ $id }}").change(() => {
    if ($("#checkbox-all-{{$id}}").prop("checked")) {
      $("table#{{$id}} tr> td> input").prop({ checked: true });
    } else {
      $("table#{{$id}} tr> td> input").prop({ checked: false });
    }
  });
</script>
{{- end }}
`
}

func (tab *Table) Generate() template.HTML {
	return gen.GenerateHtml(tab.template(), tab)
}
