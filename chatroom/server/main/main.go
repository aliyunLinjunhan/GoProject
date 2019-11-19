package main
import (

	"fmt"
	_ "errors"
	"encoding/binary"
	"encoding/json"
	"go_project/chatroom/common/message"
	"net"
	"io"
	"go_project/chatroom/server/process"
)

// // 编写一个ServerProcessLogin函数，专门处理登陆请求
// func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {

// 	// 1. 先从mes中取出mes.Data，并直接反序列化LoginMes
// 	var loginMes message.LoginMes
// 	err = json.Unmarshal([]byte(mes.Data), &loginMes)
// 	if err != nil {
// 		fmt.Println("json umarshal fail err ", err)
// 		return
// 	}

// 	//1.先声明一个resMes
// 	var resMes message.Message
// 	resMes.Type = message.LoginResMesType

// 	//2.再声明一个 LoginResMes 
// 	var LoginResMes message.LoginResMes


// 	// 如果用户id=100，密码=123456， 认为合法， 否则不合法
// 	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
// 		// 合法
// 		LoginResMes.Code = 200
// 	}else{
// 		// 不合法
// 		LoginResMes.Code = 500
// 		LoginResMes.Error = "该用户不存在，请确认账号密码是否正确！！"
// 	}

// 	//3.将loginResMes序列化
// 	data, err := json.Marshal(LoginResMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal fail ", err)
// 		return
// 	}

// 	//4. 将data赋值给resMes
// 	resMes.Data = string(data)

// 	//5. 对resMes进行序列化，准备发送
// 	data, err = json.Marshal(resMes)

// 	//6. 发送包(进行封装)
// 	err = writeRkg(conn, data)
// 	return 

// }

// // 编写ServerProcessMes 函数根据不同的消息种类，决定调用那个函数
// func ServerProcessMes(conn net.Conn, mes *message.Message) (err error) {

// 	switch mes.Type {
// 		case message.LoginMesType:
// 			// 处理登陆逻辑
// 			err = serverProcessLogin(conn, mes)
// 		case message.RegisterMesType:
// 			//
// 		default: 
// 			fmt.Println("消息类型不存在.......................")
// 	}
// 	return
// }

// 处理和客户端的通讯
func process(conn net.Conn) {
	// 先延时关闭
	defer conn.Close()

	// 调用总控
	processor := &processor{
		Conn : conn,
	}
	err = process.Process2()
	if err != nil {
		fmt.Println("客户端和服务器通讯携程err", err)
	}
	
	// // 循环的客户端发送的消息
	// for {

	// 	// 这里将读取数据包， 直接进行封装
	// 	mes, err := readPkg(conn)
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			fmt.Println("客户端推出了，服务端也正常退出...")
	// 			return
	// 		} else {
	// 			fmt.Println("readPkg err =", err)
	// 			return
	// 		}
	// 	}
	// 	// fmt.Println("mes=", mes)
	// 	err = ServerProcessMes(conn, &mes)
	// 	if err != nil {
	// 		return
	// 	}

	// }
}

// func readPkg(conn net.Conn) (mes message.Message, err error) {

// 	buf := make([]byte, 8096)
// 	fmt.Print("读取客户端发送的数据......")
// 	_, err = conn.Read(buf[:4])
// 	if err != nil {
// 		// err = errors.New("read pkg header error")
// 		return
// 	}
	
// 	// 根据buf[:4] 转成一个uint32类型
// 	var pkgLen uint32
// 	pkgLen = binary.BigEndian.Uint32(buf[0:4])

// 	// 根据pkgLen读取消息内容 
// 	n, err := conn.Read(buf[:pkgLen])
// 	if n != int(pkgLen) || err != nil {
// 		// err = errors.New("read pkg body error")
// 		return
// 	}

// 	// 把pkgLen 反序列化为 -》message.Message
// 	err = json.Unmarshal(buf[:pkgLen], &mes)
// 	if err != nil {
// 		fmt.Println("unmarshall err ", err)
// 		return
// 	}

// 	return 
// }

// func writeRkg(conn net.Conn, data []byte) (err error) {

// 	// 先发送一个包的长度，发给对方
// 	var pkgLen uint32
// 	pkgLen = uint32(len(data))
// 	var buf [4]byte
// 	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
// 	// 发送长度
// 	n, err := conn.Write(buf[:4])
// 	if n != 4 || err != nil {
// 		fmt.Println("conn write(len) err ", err)
// 		return 
// 	}
// 	fmt.Printf("客户端发送消息的长度 %d, 内容 %s", len(data), string(data))

// 	// 发送消息本身
// 	n, err = conn.Write(data)
// 	if err != nil || n != int(pkgLen) {
// 		fmt.Println("conn write(content) err ", err)
// 		return 
// 	}
// 	return 
// }

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

