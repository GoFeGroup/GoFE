package menu

import "github.com/GoFeGroup/GoFE/layout/router"

type Menu struct {
	Id      string
	Routers []router.Router
}

func NewMenu() *Menu {

	return &Menu{}
}
