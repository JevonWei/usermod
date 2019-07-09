package operate

import (
	"fmt"
	"strconv"

	"github.com/JevonWei/user/input"
)

// 删除用户
func DelUser(users map[int]map[string]string) {
	input.PrintUser(users)

	idString := input.InputString("请输入删除用户ID:")
	if id, err := strconv.Atoi(idString); err == nil {
		if user, ok := users[id]; ok {
			fmt.Println("将要删除的用户信息为:")
			input.ListUser(id, user)
			in := input.InputString("是否确定删除(Y/N)?")
			if in == "Y" || in == "y" {
				delete(users, id)
				fmt.Printf("ID为%d的用户已删除\n", id)
			}
		} else {
			fmt.Println("输入的用户ID不存在")
		}
	} else {
		fmt.Println("输入的ID不正确")
	}
}
