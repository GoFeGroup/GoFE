package button

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/GoFeGroup/GoFE/utils"
)

type Button struct {
	Id     string
	Title  template.HTML
	Action template.JS
	Style  template.CSS
}

func New() *Button {
	return &Button{
		Id:     fmt.Sprintf("btn-%s", utils.RandomString(8)),
		Title:  "Button",
		Action: "",
	}
}

func (btn *Button) SetTitle(title template.HTML) *Button {
	btn.Title = title
	return btn
}

func (btn *Button) SetStyle(style template.CSS) *Button {
	btn.Style = style
	return btn
}

func (btn *Button) template() string {
	return `<button id="{{.Id}}" class="btn" {{- with .Style }} style="{{.}}" {{- end }}>{{.Title}}</button>`
}

func (btn *Button) Generate() template.HTML {
	tpl := template.Must(template.New("button").Parse(btn.template()))

	buffer := new(bytes.Buffer)
	err := tpl.Execute(buffer, btn)
	utils.ThrowIfError(err)
	return template.HTML(buffer.String())
}
