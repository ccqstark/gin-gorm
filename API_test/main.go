package main

import (
	"github.com/ccqstark/gin_learn/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID       int    `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	Birthday string `gorm:"column:birthday`
}

type Person struct {
	ID   int    `form:"id"`
	Name string `form:"name"`
}

func main() {
	db := tools.GetDB()

	r := gin.Default()
	r.Use(CorsMiddleware()) //使用跨域中间件

	r.GET("/get_test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "success",
			"code": 1,
		})
	})

	r.GET("/url/:name/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": c.Param("name"),
			"id":   c.Param("id"),
		})
	})

	r.POST("/post_test", func(c *gin.Context) {
		u := User{}
		db.First(&u)
		c.JSON(http.StatusOK, gin.H{
			"id":       u.ID,
			"name":     u.Name,
			"birthday": u.Birthday,
		})
	})

	r.GET("get_args", func(c *gin.Context) {
		id := c.Query("id")
		code := c.Query("code")
		c.JSON(http.StatusOK, gin.H{
			"id":   id,
			"code": code,
		})
	})

	r.POST("post_args", func(c *gin.Context) {
		firstName := c.PostForm("first_name")
		lastName := c.PostForm("last_name")
		c.JSON(http.StatusOK, gin.H{
			"first_name": firstName,
			"last_name":  lastName,
		})
	})

	r.POST("/json_request", jsonFunc)

	//r.PUT()
	//
	//r.DELETE()

	r.Run()
}

func jsonFunc(c *gin.Context) {
	var person Person
	if err := c.ShouldBind(&person); err != nil {
		c.String(http.StatusBadRequest, "%v", err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id":   person.ID,
			"name": person.Name,
		})
	}
}

//CORS跨域中间件
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		// 核心处理方式
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
		c.Set("content-type", "application/json")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()
	}
}
