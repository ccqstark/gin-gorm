package main

import (
	"fmt"
	"github.com/ccqstark/gin_learn/tools"
)


//type User struct {
//	ID int  `gorm:"column:id"`
//	Name string  `gorm:"column:name"`
//	Birthday string  `gorm:"column:birthday`
//}

func main3() {
	//获取DB
	db := tools.GetDB()
	db.LogMode(true)


	u := User{
		Name:"newRecord115",
		Birthday:"987",
	}

	db.Create(&u) //一定要加&啊
	//db.Create(&u)
	//fmt.Println(db.NewRecord(&u))
	//这里会再运行插入操作一次！
	//if err := db.Create(&u).Error;err!=nil {
	//	fmt.Println("err occur",err)
	//	return
	//}

	//获取插入记录的Id
	var id []int
	db.Raw("select LAST_INSERT_ID() as id").Pluck("id", &id)

	//因为Pluck函数返回的是一列值，返回结果是slice类型，我们这里只有一个值，所以取第一个值即可。
	fmt.Println(id[0])

}

