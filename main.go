package main

import (
	"fmt"

	"github.com/GoFeGroup/GoFE/layout/router"

	menu2 "github.com/GoFeGroup/GoFE/layout/menu"

	"github.com/GoFeGroup/GoFE/layout/logo"

	"github.com/GoFeGroup/GoFE/admin"
)

func main() {

	menu := menu2.NewMenu([]router.Router{
		{Path: "/", Target: "/", SubRouters: nil},
		{Path: "/", Target: "/", SubRouters: nil},
	})

	adm := admin.NewAdmin().SetLogo(logo.NewLogo("/images/logo-white.png",
		"/images/logo.png", "/images/logo-black.png")).SetMenu(menu)
	fmt.Println(adm.Generate())
}
