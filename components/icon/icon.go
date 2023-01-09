package icon

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/GoFeGroup/GoFE/utils"
)

type Icon struct {
	Id    string
	Class string
	Style template.CSS
}

func New(class string) *Icon {
	return &Icon{
		Id:    fmt.Sprintf("i-%s", utils.RandomString(8)),
		Class: class,
	}
}

func (i *Icon) SetStyle(style template.CSS) *Icon {
	i.Style = style
	return i
}

func (i *Icon) template() string {
	return `<i id="{{.Id}}" class="bx {{.Class}}" {{- with .Style }} style="{{.}}" {{- end }}></i>`
}

func (i *Icon) Generate() template.HTML {
	tpl := template.Must(template.New("button").Parse(i.template()))

	buffer := new(bytes.Buffer)
	err := tpl.Execute(buffer, i)
	utils.ThrowIfError(err)
	return template.HTML(buffer.String())
}
