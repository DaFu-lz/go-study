package usedb

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	fmt.Println("测试函数")
	t.Run("验证用户名和密码", testLogin)
	t.Run("验证用户名和密码", testRegist)
	t.Run("验证用户名和密码", testSave)
}

func testLogin(t *testing.T) {
	user, _ := CheckUserNameAndPassWord("admin", "123456")
	fmt.Println("获取的用户信息是：", user)
}

func testRegist(t *testing.T) {
	user, _ := CheckUserName("admin")
	fmt.Println("获取的用户信息是：", user)
}
func testSave(t *testing.T) {
	SaveUser("admin1", "123456")
}
