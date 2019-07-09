package input

import (
	"fmt"
	"strings"
)

// 从命令行输入函数
func InputString(s string) string {
	var in string
	fmt.Print(s)
	fmt.Scan(&in)
	return strings.TrimSpace(in)
}

// 用户信息
func InputUser() map[string]string {
	user := map[string]string{}

	user["name"] = InputString("请输入名字:")
	user["birthday"] = InputString("请输入出生日期(2019-07-07):")
	user["tel"] = InputString("请输入联系方式:")
	user["addr"] = InputString("请输入地址:")
	fmt.Println("*******************************")
	return user
}

// 显示选择的用户
func ListUser(pk int, user map[string]string) {
	fmt.Println("================================")
	fmt.Println("ID:", pk)
	fmt.Println("Name:", user["name"])
	fmt.Println("出生日期:", user["birthday"])
	fmt.Println("联系方式:", user["tel"])
	fmt.Println("地址:", user["addr"])
}

// 打印所有的用户信息
func PrintUser(users map[int]map[string]string) {
	for id, user := range users {
		ListUser(id, user)
	}
	fmt.Println("================================")
}
