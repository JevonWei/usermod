package operate

import (
	"fmt"
	"strconv"

	"github.com/JevonWei/user/input"
)

// 修改函数
func Modify(users map[int]map[string]string) {
	input.PrintUser(users)

	idString := input.InputString("请输入修改用户ID:")

	if id, err := strconv.Atoi(idString); err == nil {
		if user, ok := users[id]; ok {
			fmt.Println("")
			fmt.Println("将要修改的用户信息为:")
			input.ListUser(id, user)
			in := input.InputString("是否确定修改(Y/N)?: ")
			if in == "Y" || in == "y" {
				user := input.InputUser()
				users[id] = user
				fmt.Printf("ID为%d的用户已修改\n", id)
			}
		} else {
			fmt.Println("输入的用户ID不存在")
		}
	} else {
		fmt.Println("输入的ID不正确")
	}
}
