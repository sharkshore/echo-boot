package main

import (
	"github.com/labstack/echo"
)

func main() {
	e:=echo.New()
	e.File("/","public/index.html")
	e.Static("/static","assets")
	e.Logger.Fatal(e.Start(":1323"))
}