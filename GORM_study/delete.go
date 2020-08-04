package main

import "github.com/ccqstark/gin_learn/tools"

func main() {

	db:=tools.GetDB()
	db.LogMode(true)

	user:=User{}
	var users[]User

	//db.Where("id=?",12).Take(&user)
	//db.Delete(&user)

	db.Where("name LIKE ?", "%xi%").Delete(User{})
	// DELETE from user where name LIKE "%xi%";

	db.Delete(User{}, "name LIKE ?", "%new%")
	// DELETE from user where name LIKE "%new%";

	db.Delete(&user)
	//// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE id = 111;

	// 批量删除
	db.Where("age = ?", 20).Delete(&User{})
	//// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE age = 20;

	// 在查询记录时，软删除记录会被忽略
	db.Where("age = 20").Find(&user)
	//// SELECT * FROM users WHERE age = 20 AND deleted_at IS NULL;

	// 使用 Unscoped 方法查找软删除记录
	db.Unscoped().Where("age = 20").Find(&users)
	//// SELECT * FROM users WHERE age = 20;

	// 使用 Unscoped 方法永久删除记录
	db.Unscoped().Delete(&user)
	//// DELETE FROM user WHERE id=10;

}
