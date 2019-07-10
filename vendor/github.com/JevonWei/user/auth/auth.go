package auth

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/howeyc/gopass"
)

const (
	maxAuth = 3
)

// 将系统密码进行md5加密
var passWord = md5.Sum([]byte("danran"))

// 认证函数，从命令行输入密码，并进行验证
// 通过返回值验证是否成功
func Auth() bool {
	for i := 0; i < maxAuth; i++ {
		fmt.Print("请输入本系统登录密码:")

		// 输入密码不回显
		if bytes, err := gopass.GetPasswd(); err == nil {

			if passWord == md5.Sum(bytes) {
				//if string(bytes) == passWord {
				return true
			} else {
				fmt.Println("密码输入错误，请重新输入...")
				fmt.Println(strings.Repeat("-", 30))
			}
		}
	}
	fmt.Printf("密码输入%d次错误，程序退出\n", maxAuth)
	return false

}
