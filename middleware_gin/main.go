package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	r:=gin.New()
	f,_ := os.Create("gin.log")
	gin.DefaultWriter=io.MultiWriter(f)
	gin.DefaultErrorWriter=io.MultiWriter(f)

	r.Use(gin.Logger(),gin.Recovery())

	r.GET("/test",func(c *gin.Context){
		name:=c.DefaultQuery("name","ccq")
		panic("test panic")
		c.String(200,"%s",name)
	})
	r.Run()
}