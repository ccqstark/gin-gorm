package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Person struct {
	Name string `form:"name"` //打标签。对应字段
	Address string `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
}


func main() {
	r:=gin.Default()
	r.GET("/testing",testing)//绑定到testing函数中
	r.POST("/testing",testing)
	r.Run()
}

func testing(c *gin.Context) {
	var person Person
	//根据请求格式类型(content-type)的不同自动解析到绑定的结构体中
	if err := c.ShouldBind(&person); err==nil {
		c.String(200,"%v",person)
	}else {
		c.String(200,"person bind error:%v",err)
	}
}