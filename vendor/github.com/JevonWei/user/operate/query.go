package operate

import (
	"fmt"
	"strings"

	"github.com/JevonWei/user/input"
)

// 查询函数
func Query(users map[int]map[string]string) {
	q := input.InputString("请输入查询的信息:")

	title := fmt.Sprintf("%-5s|%-10s|%-5s|%-10s|%-15s", "ID", "Name", "birthday", "Tel", "Addr")
	fmt.Println(title)
	fmt.Println((strings.Repeat("-", len(title))))
	for k, user := range users {
		if strings.Contains(user["name"], q) || strings.Contains(user["tel"], q) || strings.Contains(user["addr"], q) {
			fmt.Printf("%-5d|%-10s|%-5s|%-10s|%-15s\n\n", k, user["name"], user["birthday"], user["tel"], user["addr"])

		}
	}
}
