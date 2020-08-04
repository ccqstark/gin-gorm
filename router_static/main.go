package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r:=gin.Default()
	r.Static("/assets","./assets") //（路由，资源）
	r.StaticFS("/static",http.Dir("static")) //另一种写法 （路由，资源）
	r.StaticFile("/favicon.ico","./favicon.ico")//单个文件
	r.Run()

}
