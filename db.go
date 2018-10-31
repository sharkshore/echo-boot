package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"database/sql"
	"fmt"
	"net/http"
)

type Excuse struct {
	Error string `json:"error"`
	Id    int `json:"id"`
	Quote string `json:"quote"`
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	//数据库连接字符串
	db, err := sql.Open("mysql", "xinyan:xinyan@123@tcp(10.0.19.41:3306)/CREDIT_PUSH_ENGINE?charset=utf8")
	if err != nil {
		println("服务器连接失败")
	}
	defer db.Close()

	//添加
	e.POST("/insert", func(c echo.Context) error {
		ex:=new(Excuse)
		if binderror:=c.Bind(ex);binderror !=nil{
			return c.JSON(500,Excuse{Id: 0, Error: "true", Quote: "数据绑定失败"})
		}
		_, _ = db.Exec("insert into excuses (id, quote) VALUES (?, ?)", ex.Id, ex.Quote)
		return c.JSON(200,Excuse{Id: 0, Error: "true", Quote: "数据插入成功"})
	})

	//删除
	e.POST("/delete", func(c echo.Context) error {
		ex:=new(Excuse)
		if binderror:=c.Bind(ex);binderror !=nil{
			return c.JSON(500,Excuse{Id: 0, Error: "true", Quote: "数据绑定失败"})
		}
		_,err:=db.Exec("delete from excuses where id=?",ex.Id)

		if err !=nil{
			println(err.Error())
		}
		return c.JSON(200,Excuse{Id:ex.Id,Error:"true",Quote:"删除成功"})
	}	)

	//修改
	e.POST("/update", func(c echo.Context) error {
		ex:=new(Excuse)
		if binderror:=c.Bind(ex);binderror !=nil{
			return c.JSON(500,Excuse{Id: 0, Error: "true", Quote: "数据绑定失败"})
		}
		_,err:=db.Exec("update excuses set quote=? where id=?",ex.Quote,ex.Id)
		if err !=nil{
			println(err.Error())
		}
		return c.JSON(200,Excuse{Id:ex.Id,Error:"true",Quote:"修改成功"})

	})

	//查询
	e.GET("/", func(c echo.Context) error {

		var(
			arr []Excuse
		)
		rows,err:=db.Query("select id ,quote from excuses ")

		if err != nil {
			println(err)
		}
		defer rows.Close()

		for rows.Next(){
			var (
				id int
				quote  string
			)
			rows.Scan(&id,&quote)
			arr=append(arr, Excuse{Id:id,Quote:quote})
		}

		return c.JSON(http.StatusOK, arr)
	})

	//指定查询
	e.GET("/id/:id", func(c echo.Context) error {

		requested_id := c.Param("id")
		fmt.Println(requested_id)

		var quote string
		var id int
		err:= db.QueryRow("select id,quote from excuses where id=?", requested_id).Scan(&id, &quote)
		if err != nil {
			println(err)
		}
		response := Excuse{Id: id, Error: "false", Quote: quote}
		return c.JSON(200, response)
	})

	e.Logger.Fatal(e.Start(":4000"))

}
