package model

import (
	"fmt"
	"webapp/src/web01_db/utils"
)

// User 结构体
type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

// AddUser 添加User方法一
func (user *User) AddUser() error {

	//1.写sql语句
	sqlStr := "insert into user(username,password,email) values(?,?,?)"
	// 2.预编译
	inStmt, err := utils.Db.Prepare(sqlStr)

	if err != nil {
		fmt.Println("预编译出现异常：", err)
		return err
	}
	// 3.执行
	_, err2 := inStmt.Exec("admin", "password", "abc@mail.com")
	if err2 != nil {
		fmt.Println("执行出现异常：", err2)
		return err2
	}
	return nil
}

// AddUser 添加User方法二
func (user *User) AddUser2() error {

	//1.写sql语句
	sqlStr := "insert into user(username,password,email) values(?,?,?)"

	// 3.执行
	_, err := utils.Db.Exec(sqlStr, "admin2", "password", "abc2@mail.com")
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}
