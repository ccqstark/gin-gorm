package main

import (
	"fmt"
	"github.com/ccqstark/gin_learn/tools"
)

func main7() {

	db:=tools.GetDB()
	db.LogMode(true)

	//user:=User{}
	var users []User

	//等价于: SELECT * FROM `user`  WHERE (id >= 10) ORDER BY id desc
	//db.Where("id >= ?", 10).Order("id desc").Find(&users)

	//等价于: SELECT * FROM `user` ORDER BY id desc LIMIT 10 OFFSET 0
	//offset 表示从表的第几行开始(从0开始算)
	//db.Order("id desc").Limit(5).Offset(0).Find(&users)


	//total := 0
	////等价于: SELECT count(*) FROM `user`
	////这里也需要通过model设置模型，让gorm可以提取模型对应的表名
	//db.Model(User{}).Count(&total)
	//fmt.Println(total)


	//定一个Result结构体类型，用来保存查询结果
	type Result struct {
		Birthday string
		Total int
	}

	var results []Result
	////等价于: SELECT type, count(*) as total FROM `user` GROUP BY type HAVING (total > 0)
	//db.Model(User{}).Select("birthday, count(*) as  total").Group("birthday").Having("total >= 9").Scan(&results)
	//
	////scan类似Find都是用于执行查询语句，然后把查询结果赋值给结构体变量，区别在于scan不会从传递进来的结构体变量提取表名.
	////这里因为我们重新定义了一个结构体用于保存结果，但是这个结构体并没有绑定user表，所以这里只能使用scan查询函数。
	//
	//fmt.Println(results)


	sql := "SELECT birthday, count(*) as  total FROM `user` where id > ? GROUP BY birthday HAVING (total > 0)"
	//因为sql语句使用了一个问号(?)作为绑定参数, 所以需要传递一个绑定参数(Raw第二个参数).
	//Raw函数支持绑定多个参数
	db.Raw(sql, 25).Scan(&results)
	fmt.Println(results)


	//fmt.Println(user)

	for _,data := range users{
		fmt.Println(data)
	}

}
