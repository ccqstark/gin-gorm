package main

import (
	"fmt"
	"github.com/ccqstark/gin_learn/tools"
)

func main5() {

	db:=tools.GetDB()
	//user:=User{}
	var users []User

	//等价于: SELECT * FROM `user`  WHERE (id = '9') LIMIT 1
	//这里问号(?), 在执行的时候会被10替代
	//db.Where("id = ?", 9).Take(&user)


	// in 语句
	//等价于: SELECT * FROM `user`  WHERE (id in ('1','2','5','8'))
	//args参数传递的是数组
	//db.Where("id in (?)", []int{1,2,5,8}).Find(&users)


	//等价于: SELECT * FROM `user`  WHERE (id >= 10 and id <= 30)
	//这里使用了两个问号(?)占位符，后面传递了两个参数替换两个问号。
	//db.Where("id >= ? and id <= ?", 10, 30).Find(&users)


	//like语句
	//等价于: SELECT * FROM `user`  WHERE (name like '%cc%')
	db.Where("name like ?", "%cc%").Find(&users)

	//fmt.Println(user)

	for _,data := range users{
		fmt.Println(data)
	}
}
