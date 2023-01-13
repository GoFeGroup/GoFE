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
 --logo-image: url("{{.Image}}");
 --logo-sm-image: url("{{.ImageInCollapse}}");
}
${{.Id}}.light {
 --logo-image: url("{{.ImageInLight}}");
}
</style>
<div id="{{.Id}}" class="gfe-logo {{- if .LightMode }} light {{- end }}"></div>
`
}

func (logo *Logo) SetLightMode(light bool) *Logo {
	logo.LightMode = light
	return logo
}

func (logo *Logo) SetLightModeJS(light bool) template.JS {
	return utils.If(light,
		template.JS(fmt.Sprintf(`${"#%v"}.addClass("light");`, logo.Id)),
		template.JS(fmt.Sprintf(`${"#%v"}.removeClass("light");`, logo.Id)))
}

func (logo *Logo) Generate() template.HTML {
	return gen.GenerateHtml(logo.template(), logo)
}
