package admin

import (
	"fmt"
	"html/template"

	"github.com/GoFeGroup/GoFE/layout/content"
	"github.com/GoFeGroup/GoFE/layout/footer"
	"github.com/GoFeGroup/GoFE/layout/header"

	"github.com/GoFeGroup/GoFE/components/gen"
	"github.com/GoFeGroup/GoFE/layout/logo"
	"github.com/GoFeGroup/GoFE/layout/menu"
	"github.com/GoFeGroup/GoFE/utils"
)

type Admin struct {
	Id      string
	Logo    *logo.Logo
	Menu    menu.Menu
	Header  header.Header
	Content content.Content
	Footer  footer.Footer
}

func NewAdmin() *Admin {
	return &Admin{
		Id: fmt.Sprintf("admin-%v", utils.RandomString(8)),
	}
}

func (adm *Admin) SetLogo(logo *logo.Logo) *Admin {
	adm.Logo = logo
	return adm
}

func (adm *Admin) layout() string {
	return `
<div id={{.Id}} class="gfe-admin">
	<div id="sidebar-{{.Id}}" class="gfe-sidebar">
		{{ block "logo" . }} GoFE - logo {{ end }}
		{{ block "menu" . }} GoFE - menu {{ end }}
	</div>
	<div id="main-{{.Id}}" class="gfe-main">
		{{ block "header" . }} GoFE - header {{ end }}
		{{ block "content" . }} GoFE - content {{ end }}
		{{ block "footer" . }} GoFE - footer {{ end }}
	</div>
</div>`
}

func (adm *Admin) wrap(tplName string, tpl template.HTML) string {
	return fmt.Sprintf(`{{ define "%v" }} %v {{ end }}`, tplName, tpl)
}

func (adm *Admin) Generate() template.HTML {
	tpl := template.Must(template.New("tpl").Parse(adm.layout()))
	tpl = template.Must(tpl.Parse(adm.wrap("logo", adm.Logo.Generate())))
	tpl = template.Must(tpl.Parse(adm.wrap("menu", adm.Menu.Generate())))

	return gen.GenerateHtmlByTemplate(tpl, adm)
}
