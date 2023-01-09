package tag

import (
	"fmt"
	"html/template"

	"github.com/GoFeGroup/GoFE/components/gen"
	"github.com/GoFeGroup/GoFE/utils"
)

type Tag struct {
	Id    string
	Title template.HTML
	Style template.CSS
}

func New(title template.HTML) *Tag {
	return &Tag{
		Id:    fmt.Sprintf("tag-%s", utils.RandomString(8)),
		Title: title,
	}
}

func (t *Tag) SetStyle(style template.CSS) *Tag {
	t.Style = style
	return t
}

func (t *Tag) template() string {
	return `<span id="{{.Id}}" class="tag" {{- with .Style }} style="{{.}}" {{- end }}>{{.Title}}</span>`
}

func (t *Tag) Generate() template.HTML {
	return gen.GenerateHtml(t.template(), t)
}
