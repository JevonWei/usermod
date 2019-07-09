package getid

func GetId(users map[int]map[string]string) int {
	var id int
	for k := range users {
		if id < k {
			id = k
		}
	}
	return id + 1
}
