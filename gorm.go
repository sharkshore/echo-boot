package main

import (
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"time"
)

type (
	UserAction struct {
		Id int `gorm:"column:ID;AUTO_INCREMENT" json:"id"`
		MerchantNo string `gorm:"column:MERCHANT_NO;not null" json:"merchantNo"`
		AppId int `gorm:"column:APP_ID" json:"appId"`
		ChannelOrderNo string `gorm:"column:CHANNEL_ORDER_NO;not null" json:"channelOrderNo"`
		MessageContent string `gorm:"column:MESSAGE_CONTENT;not null" json:"messageContent"`
		Os string `gorm:"column:OS;not null" json:"os"`
		RegId string `gorm:"column:REG_ID" json:"regId"`
		Action string `gorm:"column:ACTION" json:"action"`
		CreatedAt *time.Time `gorm:"column:CREATED_AT" json:"createdAt"`
		UpdatedAt *time.Time `gorm:"column:UPDATED_AT" json:"updatedAt"`
	}

)

type jsonTime time.Time

func (this jsonTime) MarshalJSON()([]byte,error)  {

	
}


func main() {
	db, err := gorm.Open("mysql", "xinyan:xinyan@123@tcp(10.0.19.41:3306)/CREDIT_PUSH_ENGINE?charset=utf8&parseTime=True&loc=Local")

	if err!=nil{
		panic("failed to connect database")
	}
	defer db.Close()

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	e:=echo.New()
	e.Use(middleware.Logger())

	//添加
	e.POST("/", func(c echo.Context) error {
        ua:=new(UserAction)
		if bindError:=c.Bind(ua);bindError!=nil{
			println(bindError.Error())
			return c.JSON(500, "数据绑定失败")
		}
		db.Exec("insert into T_USER_ACTION values(?,?,?,?,?,?,?,?,?)",ua.MerchantNo,ua.AppId,ua.ChannelOrderNo,ua.MessageContent,ua.Os,ua.RegId,ua.Action,time.Now(),time.Now())
		return c.JSON(200,"成功")
	})

	e.Logger.Fatal(e.Start(":3399"))



}
