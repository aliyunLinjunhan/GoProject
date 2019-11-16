package main
import(
	"fmt"
	"net"
)

func process(conn net.Conn){

	// 这里循环接受客户端发送的数据
	defer conn.Close()	// 关闭conn

	for{
		// 创建一个新的切片
		buf := make([]byte, 1024)
		// 1. 等待客户端通过conn发送信息
		// 2. 如果客户端没有write[发送],那么携程就阻塞在这里
		// fmt.Printf("等待客户端的信息......， 客户端IP为 %v\n", conn.RemoteAddr())
		n, err := conn.Read(buf)	// 从conn读取数据
		if err != nil {
			fmt.Println("服务器端的read err=", err)
			return 
		}
		// 3.显示客户端发送的内容到服务器
		fmt.Println(string(buf[:n]))
	}

}

func main() {
	
	fmt.Println("服务器开始监听.......")
	// tcp表示使用的网络协议，并且在本地的8888端口
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil{
		fmt.Println("listen err", err)
		return
	}
	defer listen.Close()  //延时关闭

	for {
		// 等待客户端连接
		fmt.Println("等待客户端来连接......")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accpet err=", err)
		}else{
			fmt.Printf("Accept() suc con=%v 客户端ip为%v \n", conn, conn.RemoteAddr())
		}
		go process(conn)
	}

	fmt.Printf("listen suc=%v \n ", listen)
}