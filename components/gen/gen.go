package gen

import (
	"bytes"
	"html/template"

	"github.com/GoFeGroup/GoFE/utils"
)

func GenerateHtml(tplStr string, obj any) template.HTML {
	tpl := template.Must(template.New("tpl").Parse(tplStr))

	buffer := new(bytes.Buffer)
	utils.ThrowIfError(tpl.Execute(buffer, obj))
	return template.HTML(buffer.String())
}
