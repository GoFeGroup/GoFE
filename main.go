package main

import (
	"fmt"

	"github.com/GoFeGroup/GoFE/layout/router"

	"github.com/gin-gonic/gin"

	"github.com/GoFeGroup/GoFE/admin"
	"github.com/GoFeGroup/GoFE/layout/logo"
	"github.com/GoFeGroup/GoFE/layout/menu"
)

func main() {
	//r := &router.Routers{R: []router.Router{
	//	{Path: "/", Target: a, SubRouters: nil},
	//	{Path: "/", Target: b, SubRouters: nil},
	//}}
	m := menu.NewMenu([]router.Router{
		{Path: "/", Target: a, SubRouters: nil},
		{Path: "/", Target: b, SubRouters: nil},
	})

	adm := admin.NewAdmin().SetMenu(m).SetLogo(logo.NewLogo("/images/logo-white.png",
		"/images/logo.png", "/images/logo-black.png").SetLightMode(true))
	fmt.Println(adm.Generate())

	//_ = r.Static().Run(":80")
}

func b(ctx *gin.Context) {

}

func a(ctx *gin.Context) {

}
