package main

import (
	"fmt"
	"github.com/ccqstark/gin_learn/tools"
)


func main4() {

	db:=tools.GetDB()

	user:=User{}
	//var users []User
	//
	////获取第一条记录，按主键排序
	//db.First(&user)
	//// SELECT * FROM users ORDER BY id LIMIT 1;
	//
	////获取一条记录，不指定排序
	//db.Take(&user)
	//// SELECT * FROM users LIMIT 1;
	//
	////获取最后一条记录，按主键排序
	//db.Last(&user)
	////SELECT * FROM users ORDER BY id DESC LIMIT 1;
	//
	////获取所有的记录
	//db.Find(&users)
	////SELECT * FROM users;
	//
	////通过主键进行查询 (仅适用于主键是数字类型)
	//db.First(&user, 8)
	//// SELECT * FROM users WHERE id = 10;

	//for _,data := range users{
	//	fmt.Println(data)
	//}
	//fmt.Println(user)
	//
	//var name[]string
	//db.Model(User{}).Pluck("name",&name)
	//
	//for _,data := range name{
	//	fmt.Println(data)
	//}

	//err := db.First(&user,100).Error
	//if gorm.IsRecordNotFoundError(err) {
	//	fmt.Println("no data")
	//} else if err != nil {
	//	//如果err不等于record not found错误，又不等于nil，那说明sql执行失败了。
	//	fmt.Println("query failed", err)
	//}

	//链式操作，先查询，然后检测查询结果
	if db.First(&user,100).RecordNotFound() {
		fmt.Println("no data")
	}

}
