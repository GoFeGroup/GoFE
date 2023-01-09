package components

import (
	"bytes"

	"github.com/GoFeGroup/GoFE/utils"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
)

func MinifyCSS(cssStr string) string {
	m := minify.New()
	r := bytes.NewBufferString(cssStr)
	w := &bytes.Buffer{}
	utils.ThrowIfError(css.Minify(m, w, r, nil))
	return w.String()
}
