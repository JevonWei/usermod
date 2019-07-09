package operate

import (
	"fmt"

	"github.com/JevonWei/user/getid"
	"github.com/JevonWei/user/input"
)

// 添加函数
func Add(users map[int]map[string]string) {
	id := getid.GetId(users)

	// 调用用户函数，新增用户
	user := input.InputUser()

	users[id] = user
	fmt.Printf("ID为%d的用户已添加\n", id)

	// name := inputString("请输入名字: ")
	// birthday := inputString("请输入出生日期(2019-07-07)：")
	// tel := inputString("请输入联系方式: ")
	// addr := inputString("请输入地址: ")

	// users[id] = {
	// 	"name" := name,
	// 	"birthday" := birthday,
	// 	"tel" := tel,
	// 	"addr" := addr,
	// }
}
