package main

import (
	"2022summer/initialize"
)

func main() {
	err := initialize.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer initialize.Close()

	r := initialize.SetupRouter()

	/*user1 := model.User{Username: "wxy", Password: "123qweasd", RealName: "王新元", Email: "wangxinyuan4869@163.com"}
	err = service.CreateUser(&user1)
	if err != nil {
		panic(err)
	}*/
	/*fmt.Println(service.QueryUserByUsername("wxy"))
	fmt.Println(service.QueryUserByUserID(1))
	fmt.Println(service.QueryUserByEmail("wangxinyuan4869@163.com"))
	fmt.Println(service.QueryUserByUsername("orz"))
	fmt.Println(service.QueryUserByUserID(2))
	fmt.Println(service.QueryUserByEmail("orz@163.com"))*/
	/*user2, _ := service.QueryUserByUsername("wxy")
	user2.Username = "orz"
	err = service.UpdateUser(&user2)
	if err != nil {
		panic(err)
	}*/

	if err := r.Run("localhost:8000"); err != nil {
		panic(err)
	}
}
