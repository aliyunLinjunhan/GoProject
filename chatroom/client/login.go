package main
import (
	"fmt"
	"encoding/json"
	"net"
	"encoding/binary"
	"go_code/chatroom/common/message"
)

// 写一个登陆函数
func login(userId int, userPwd string) (err error) {

	// 下一个就要开始定协议
	// 1. 链接到服务器
	conn, err := net.Dail("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net dial err ", err)
		return 
	}
	defer conn
}