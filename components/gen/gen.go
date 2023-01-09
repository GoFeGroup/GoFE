package gen

import (
	"bytes"
	"html/template"

	"github.com/GoFeGroup/GoFE/utils"
)

func GenerateHtml(tplStr string, obj any) template.HTML {
	tpl := template.Must(template.New("button").Parse(tplStr))

	buffer := new(bytes.Buffer)
	err := tpl.Execute(buffer, obj)
	utils.ThrowIfError(err)
	return template.HTML(buffer.String())
}
