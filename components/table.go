package components

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/GoFeGroup/GoFE/utils"
)

type TableCell any

type TableHeader struct {
	Id   string
	Cell []TableCell
}

func NewTableHeader() *TableHeader {
	return &TableHeader{
		Id: fmt.Sprintf("th-%s", utils.RandomString(8)),
	}
}

func (th *TableHeader) AddCell(cell TableCell) *TableHeader {
	th.Cell = append(th.Cell, cell)
	return th
}

type TableRow struct {
	Id     string
	DataId string
	Cell   []TableCell
}

func NewTableRow() *TableRow {
	return &TableRow{
		Id: fmt.Sprintf("tr-%s", utils.RandomString(8)),
	}
}

func (tr *TableRow) SetDataId(id string) *TableRow {
	tr.DataId = id
	return tr
}

func (tr *TableRow) AddCell(cell TableCell) *TableRow {
	tr.Cell = append(tr.Cell, cell)
	return tr
}

type Table struct {
	Id       string
	Header   *TableHeader
	Rows     []*TableRow
	Checkbox bool
}

func NewTable() *Table {
	return &Table{
		Id: fmt.Sprintf("tab-%s", utils.RandomString(8)),
	}
}

func (tab *Table) SetCheckbox(checkbox bool) *Table {
	tab.Checkbox = checkbox
	return tab
}

func (tab *Table) AddHeader(header *TableHeader) *Table {
	tab.Header = header
	return tab
}

func (tab *Table) AddRow(row *TableRow) *Table {
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
	tpl := template.Must(template.New("table").Parse(tab.template()))

	buffer := new(bytes.Buffer)
	err := tpl.Execute(buffer, tab)
	utils.ThrowIfError(err)
	return template.HTML(buffer.String())
}
