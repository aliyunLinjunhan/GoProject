package main
import (

	"fmt"
	"net"
)

// 处理和客户端的通讯
func process(conn net.Conn) {
	// 先延时关闭
	defer conn.Close()

	// 循环的客户端发送的消息
	for {
		buf := make([]byte, 8096)
		fmt.Print("读取客户端发送的数据......")
		n, err := conn.Read(buf[:4])
		if n != 4 || err != nil {
			fmt.Println("conn read err", err)
			return
		}
		fmt.Println("读到的buf=", buf[:4])
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
		go process(conn)
	}
}

