package logo

import (
	"fmt"
	"html/template"

	"github.com/GoFeGroup/GoFE/components/gen"
	"github.com/GoFeGroup/GoFE/utils"
)

type Logo struct {
	Id              string
	LightMode       bool
	Image           string // 图片
	ImageInCollapse string // Collapse模式的图片
	ImageInLight    string // Light模式的图片
}

func NewLogo(image string, images ...string) *Logo {
	logo := &Logo{
		Id:              fmt.Sprintf("logo-%s", utils.RandomString(8)),
		Image:           image,
		ImageInLight:    image,
		ImageInCollapse: image,
	}
	if len(images) >= 2 {
		logo.ImageInCollapse = images[0]
		logo.ImageInLight = images[1]
	} else if len(images) >= 1 {
		logo.ImageInCollapse = images[0]
	}
	return logo
}

func (logo *Logo) template() string {
	return `<style>
#{{.Id}} {
{{ if .LightMode -}}
 --logo-image: url("{{.ImageInLight}}");
{{ else -}}
 --logo-image: url("{{.Image}}");
{{ end -}}
 --logo-sm-image: url("{{.ImageInCollapse}}");
}
</style>
<div id="{{.Id}}" class="gfe-logo"></div>
`
}

func (logo *Logo) Generate() template.HTML {
	return gen.GenerateHtml(logo.template(), logo)
}
