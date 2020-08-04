package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"reflect"
	"strings"
	"time"
)

type Booking struct {
	CheckIn time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

//自定义验证器
func bookableDate(v *validator.Validate, topStruct reflect.Value,
	currentStructOrField reflect.Value, field reflect.Value,
	fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if date,ok:=field.Interface().(time.Time);ok {
		today:=time.Now()
		if date.Unix()>today.Unix(){
			return true
		}
	}
	return false
}


func main() {
	r:=gin.Default()

	if v,ok := binding.Validator.Engine().(*validator.Validate);ok{
		v.RegisterValidation("bookabledate",bookableDate)
	}

	r.GET("/booking", func(c *gin.Context){
		var b Booking
		if err := c.ShouldBind(&b);err!=nil {
			fmt.Println(err.Error())
			fmt.Println("--------")
			errArray:=strings.Split(err.Error(),"\n")
			firstErr:=errArray[0]
			switch firstErr{
			case "Key: 'Booking.CheckIn' Error:Field validation for 'CheckIn' failed on the 'bookabledate' tag":
				c.JSON(500, gin.H{"error":1})
			case "Key: 'Booking.CheckOut' Error:Field validation for 'CheckOut' failed on the 'gtfield' tag":
				c.JSON(500, gin.H{"error":2})
			}
			return
		}
		c.JSON(200,gin.H{"message":"ok!","booking":b})
	})
	r.Run()
}
