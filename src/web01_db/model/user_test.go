package model

import (
	"fmt"
	"testing"
)

// testmain函数可以在测试函数执行之前做一些其他操作
func TestMain(m *testing.M) {
	fmt.Println("测试开始：")
	//通过m.Run()来执行子测试函数
	m.Run()
}
func TestUser(t *testing.T) {
	fmt.Println("开始测试User中的相关方法")
	// 通过t.run()来执行子测试函数
	t.Run("测试添加用户：", testAddUser)
}

func testAddUser(t *testing.T) {

	fmt.Println("子测试添加用户：")
	user := &User{}
	// 调用添加用户的方法
	user.AddUser()
	user.AddUser2()
}
