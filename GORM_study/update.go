package main

import (
	"github.com/ccqstark/gin_learn/tools"
	"github.com/jinzhu/gorm"
)

func main8() {
	db:=tools.GetDB()
	db.LogMode(true)
	user:=User{}
	//user:=User{
	//	ID: 8,
	//}
	//
	//db.Model(&user).Update("name","jjjj")

	//db.Where("id=?",9).Take(&user)
	//user.Birthday="77777"
	//db.Save(&user)

	//fmt.Println(db.Model(User{}).Updates(User{Birthday: "2000"}).RowsAffected)

	db.Where("id=?",2).Take(&user)

	db.Model(&user).Update("birthday", gorm.Expr("birthday * ? + ?", 2, 100))
	// UPDATE "products" SET "price" = price * '2' + '100', "updated_at" = '2013-11-17 21:34:10' WHERE "id" = '2';

	db.Model(&user).UpdateColumn("birthday", gorm.Expr("birthday - ?", 1))
	// UPDATE "products" SET "quantity" = quantity - 1 WHERE "id" = '2';



}
