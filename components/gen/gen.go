package gen

import (
	"bytes"
	"html/template"

	"github.com/GoFeGroup/GoFE/utils"
)

func GenerateHtmlByTemplate(tpl *template.Template, obj any) template.HTML {
	buffer := new(bytes.Buffer)
	utils.ThrowIfError(tpl.Execute(buffer, obj))
	return template.HTML(buffer.String())
}

func GenerateHtml(tplStr string, obj any) template.HTML {
	tpl := template.Must(template.New("tpl").Parse(tplStr))
	return GenerateHtmlByTemplate(tpl, obj)
}
