package static

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/GoFeGroup/GoFE/utils"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
)

var defaultStyles []string

func MinifyCSS(cssStr string) string {
	m := minify.New()
	r := bytes.NewBufferString(cssStr)
	w := &bytes.Buffer{}
	utils.ThrowIfError(css.Minify(m, w, r, nil))
	return w.String()
}

func RegisterStyle(style template.CSS) {
	defaultStyles = append(defaultStyles, MinifyCSS(string(style)))
}

var mainStyle = `
* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

html, body {
    width: 100%;
    height: 100%;
}`

func _generateStyles() string {
	m := MinifyCSS(mainStyle)
	return m + strings.Join(defaultStyles, "")
}

func GenerateStyles() template.CSS {
	return template.CSS(_generateStyles())
}
