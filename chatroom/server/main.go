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

