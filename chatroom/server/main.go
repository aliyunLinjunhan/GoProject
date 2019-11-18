package main
import (

	"fmt"
	_ "errors"
	"encoding/binary"
	"encoding/json"
	"go_project/chatroom/common/message"
	"net"
	"io"
)

// 处理和客户端的通讯
func process(conn net.Conn) {
	// 先延时关闭
	defer conn.Close()
	
	// 循环的客户端发送的消息
	for {

		// 这里将读取数据包， 直接进行封装
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端推出了，服务端也正常退出...")
				return
			} else {
				fmt.Println("readPkg err =", err)
				return
			}
		}
		fmt.Println("mes=", mes)

	}
}

func readPkg(conn net.Conn) (mes message.Message, err error) {

	buf := make([]byte, 8096)
	fmt.Print("读取客户端发送的数据......")
	_, err = conn.Read(buf[:4])
	if err != nil {
		// err = errors.New("read pkg header error")
		return
	}
	
	// 根据buf[:4] 转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])

	// 根据pkgLen读取消息内容 
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		// err = errors.New("read pkg body error")
		return
	}

	// 把pkgLen 反序列化为 -》message.Message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("unmarshall err ", err)
		return
	}

	return 
	

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

