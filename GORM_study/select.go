package main

import (
	"fmt"
	"github.com/ccqstark/gin_learn/tools"
)


func main6() {
	db:=tools.GetDB()
	db.LogMode(true)

	user:=User{}

	//等价于: SELECT id,name FROM `user`  WHERE `user`.`id` = '1' AND ((id = '1')) LIMIT 1
	db.Select("id,name").Where("id = ?", 1).Take(&user)

	//这种写法是直接往Select函数传递数组，数组元素代表需要选择的字段名
	db.Select([]string{"id", "name"}).Where("id = ?", 1).Take(&user)


	//可以直接书写聚合语句
	//等价于: SELECT count(*) as total FROM `user`
	total := []int{}
	//Pluck接收的一定要是slice

	//Model函数，用于指定绑定的模型，这里生成了一个User{}变量。目的是从模型变量里面提取表名，Pluck函数我们没有直接传递绑定表名的结构体变量，gorm库不知道表名是什么，所以这里需要指定表名
	//Pluck函数，主要用于查询一列值
	db.Model(User{}).Select("count(*) as total").Pluck("total", &total)

	fmt.Println(total[0])
	
	//fmt.Println(user)
}
