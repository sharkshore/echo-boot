package main

import (
	"github.com/labstack/echo"
	"net/http"
)

type User struct{
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func main() {
	e:=echo.New()
	e.GET("/", func(context echo.Context) error {
		u:=new(User)
		u.Name="will"
		u.Email="will@163.com"
		return context.JSON(http.StatusOK,u)
	})
	e.Logger.Fatal(e.Start(":1323"))
}