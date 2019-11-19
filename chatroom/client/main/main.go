package main

import (
	"fmt"
	"go_project/chatroom/client/process"
)

// 定义俩个变量，一个表示id，一个表示password
var (
	UserId int
	UserPassword string
)

func main() {

	// 接受用户选择
	var key int
	// 判断是否还继续现实菜单
	var loop = true

	for loop {
		fmt.Println("---------------------------欢迎登陆多人聊天系统-------------------------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 4 清选择(1-3):")

		fmt.Scanf("%d\n", &key)
		switch key {

			case 1: 
				fmt.Println("登陆聊天室")
				//说明用户要登陆
				fmt.Println("请输入您的ID")
				fmt.Scanf("%d\n", &UserId)
				fmt.Println("请输入您的密码")
				fmt.Scanf("%s\n", &UserPassword)
				// 使用登陆函数
				up := &process.UserProcess{

				}
				up.Login(UserId, UserPassword)
				loop = false
			case 2: 
				fmt.Println("注册用户")
				loop = false
			case 3: 
				fmt.Println("退出系统")
				loop = false
			default : 
				fmt.Println("您的输入有误！！！")
		}
	}
}