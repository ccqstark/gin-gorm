package main

import (
	"github.com/ccqstark/gin_learn/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID int  `gorm:"column:id"`
	Name string  `gorm:"column:name"`
	Birthday string  `gorm:"column:birthday`
}

func main() {
	db:=tools.GetDB()

	r:=gin.Default()

	r.POST("/api", func(c *gin.Context){
		u:=User{}
		db.First(&u)
		c.JSON(http.StatusOK, gin.H{
			"id":u.ID,
			"name":u.Name,
			"birthday":u.Birthday,
		})
	})
	r.Run()

}
