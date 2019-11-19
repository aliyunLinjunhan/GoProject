package main
import (
	"fmt"
	"time"
	_ "errors"
	"net"
	"go_project/chatroom/server/model"
)


// 处理和客户端的通讯
func process1(conn net.Conn) {
	// 先延时关闭
	defer conn.Close()

	// 调用总控
	processor := &Processor{
		Conn : conn,
	}
	err := processor.Process2()
	if err != nil {
		fmt.Println("客户端和服务器通讯携程err", err)
	}
}

// 编写一个函数完成对UserDao的初始化任务
func initUserDao(){
	
	model.MyUserDao = model.NewUserDao(pool)
}
	

func main() {
	
	// 初始化redis链接池
	initPool("localhost:6379", 16, 0, 300 *time.Second)
	// 先初始化redis链接池，再初始化UserDao
	initUserDao()
	// 提示信息
	fmt.Println("服务器在88889端口监听.........")
	listen, err := net.Listen("tcp", "0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.listen err ", err)
		return
	}
	// 一旦监听成功，就等待客户端来连接
	for {
		fmt.Println("等待客户端来链接服务器.........")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen accept err ", err)
		}
		// 一旦链接成功，则启动一个携程与客户端保持通讯
		go process1(conn)
	}
}

