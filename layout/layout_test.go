package layout

import (
	"fmt"
	"testing"

	"github.com/GoFeGroup/GoFE/layout/logo"
)

func TestNewLogo(t *testing.T) {
	fmt.Println(logo.NewLogo("/images/logo-white.png",
		"/images/logo.png", "/images/logo-black.png").Generate())
}
