package main

import "github.com/gin-gonic/gin"

func main() {
	r:=gin.Default()
	r.POST("/user/*action",func(c *gin.Context){
		c.String(200,"hello ccq")
	})

	r.Run()
}
