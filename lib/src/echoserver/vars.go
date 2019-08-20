package echoserver

import (
	"github.com/trhhosting/keyword/lib/src/config"
	"gitlab.com/labstack/echo"
)

var (
	path = currentPath()
	cfg config.Cockroachdatastore
)
var e = echo.New()

func init() {
	e.HideBanner = false
	e.Debug = true
	e.Logger.Info()
}
