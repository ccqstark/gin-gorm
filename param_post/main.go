package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r:=gin.Default()
	r.POST("/test", func(c *gin.Context){
		firstName:=c.PostForm("first_name")
		lastName:=c.DefaultPostForm("last_name","ccqstark")
		c.String(http.StatusOK,"%s %s",firstName,lastName)
	})

	r.Run()
}
