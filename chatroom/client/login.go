package main
import (
	"fmt"
	"encoding/json"
	"net"
	"encoding/binary"
	"go_project/chatroom/common/message"
)

// 写一个登陆函数
func login(userId int, userPwd string) (err error) {

	// 下一个就要开始定协议
	// 1. 链接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net dial err ", err)
		return 
	}
	// 延时关闭
	defer conn.Close()

	// 2. 准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.LoginMesType
	// 3.创建了一个LoginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	// 4.将loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json marshal err ", err)
		return
	}

	// 5.把data赋给mes.Data字段
	mes.Data = string(data)

	// 6.将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes json marshall err ", err)
		return
	}

	// 7. 已经拿到了要送的data
	// 7.1 先把data长度发送给服务器
	// 先获取data的长度 -》 转成一个表示长度的bytes的切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	// 发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn write(len) err ", err)
		return 
	}
	fmt.Printf("客户端发送消息的长度 %d, 内容 %s", len(data), string(data))

	// 发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn write(content) err ", err)
		return 
	}

	// 这里需要处理服务器的反馈信息
	mes, err = readPkg(conn)	// mes 就是
	if err != nil {
		fmt.Println("readPkg 出错了！！", err)
	}

	// 将mess反序列化为LoginRes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登陆成功!!!")
	}else{
		fmt.Println(loginResMes.Error)
	}

	return 
}