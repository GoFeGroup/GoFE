package components

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/GoFeGroup/GoFE/static"
	"github.com/GoFeGroup/GoFE/utils"
)

type Button struct {
	Id     string
	Title  string
	Action template.JS
}

func NewButton() *Button {
	return &Button{
		Id:     fmt.Sprintf("btn-%s", utils.RandomString(8)),
		Title:  "Button",
		Action: "",
	}
}

func (btn *Button) SetTitle(title string) *Button {
	btn.Title = title
	return btn
}

func (btn *Button) template() string {
	return `<button id="{{.Id}}" class="btn">{{.Title}}</button>`
}

func (btn *Button) Generate() template.HTML {
	tpl := template.Must(template.New("button").Parse(btn.template()))

	buffer := new(bytes.Buffer)
	err := tpl.Execute(buffer, btn)
	utils.ThrowIfError(err)
	return template.HTML(buffer.String())
}

var defaultButtonStyle = template.CSS(`
.btn {
    position: relative;
    background: #2e82ff;
    color: #fff;
    border: 1px solid transparent;
    border-radius: 8px;
    min-height: 26px;
    min-width: 48px;
    font-size: 14px;
    padding: 0 8px;
}

.btn:focus,
.btn:hover {
    opacity: 0.8;
    transition: 0.3s;
}`)

func init() {
	static.RegisterStyle(defaultButtonStyle)
}
