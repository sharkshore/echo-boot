package echodemo

import (
	"io"
	"github.com/labstack/echo"
	"html/template"
	"net/http"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer,name string ,data interface{} ,c echo.Context) error  {
	return t.templates.ExecuteTemplate(w,name,data)
}

func Hello(c echo.Context) error{
	return c.Render(http.StatusOK,"hello","will")
}
func main() {
	t:=&Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e:=echo.New()
	e.Renderer=t
	e.GET("/hello",Hello)
	e.Logger.Fatal(e.Start(":1323"))
}
