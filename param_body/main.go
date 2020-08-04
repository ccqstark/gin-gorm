package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	r:=gin.Default()
	r.POST("/test", func(c *gin.Context){
		bodyByte,err := ioutil.ReadAll(c.Request.Body)
		c.Request.Body=ioutil.NopCloser(bytes.NewBuffer(bodyByte))
		if err != nil {
			c.String(http.StatusBadRequest,err.Error())
			c.Abort()
		}
		firstName:=c.PostForm("first_name")
		lastName:=c.DefaultPostForm("last_name","ccqstark")
		c.String(http.StatusOK,"%s %s %s",firstName,lastName,string(bodyByte))
	})

	r.Run()
}
