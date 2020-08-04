package main
//导入tools包
import (
	"fmt"
	"github.com/ccqstark/gin_learn/tools"
)

//定义这个结构体对应某一个表的字段结构，打的标签代表对应的字段，关闭表名复数的话对应的表就是user
type User struct {
	ID int  `gorm:"primary_key"`
	Name string  `gorm:"column:name"`
	Birthday string  `gorm:"column:birthday`
}

func main2() {
	//获取DB
	db := tools.GetDB()
	db.LogMode(true)

	//执行数据库查询操作前，用模型生成一个来存储记录的对象（也对应了要查询的表）
	u := User{}

	//自动生成sql： SELECT * FROM `user`  WHERE (name = 'ccq') LIMIT 1
	db.Where("name = ?", "ccq").First(&u)

	fmt.Println(u)
}