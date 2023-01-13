package menu

import (
	"fmt"
	"html/template"

	"github.com/GoFeGroup/GoFE/layout/router"
	"github.com/GoFeGroup/GoFE/utils"
)

type Menu struct {
	Id      string
	Routers []router.Router
}

func NewMenu(routers []router.Router) *Menu {
	return &Menu{
		Id:      fmt.Sprintf("menu-%v", utils.RandomString(8)),
		Routers: routers,
	}
}

func (menu *Menu) Generate() template.HTML {
	return ``
}
