package main

import (
	"fmt"

	"github.com/GoFeGroup/GoFE/layout/logo"

	"github.com/GoFeGroup/GoFE/admin"
)

func main() {

	adm := admin.NewAdmin().SetLogo(logo.NewLogo("/images/logo-white.png",
		"/images/logo.png", "/images/logo-black.png"))
	fmt.Println(adm.Generate())
}
