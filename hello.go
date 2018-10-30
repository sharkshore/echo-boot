package echodemo

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e:=echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK,"Hello,world")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
