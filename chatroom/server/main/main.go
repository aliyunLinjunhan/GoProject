package main
import (
	"fmt"
	_ "errors"
	"net"
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
	

func main() {
	
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

