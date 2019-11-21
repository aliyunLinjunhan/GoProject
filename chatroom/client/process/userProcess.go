package process
import (
	"fmt"
	"encoding/json"
	"net"
	"go_project/chatroom/common/message"
	"go_project/chatroom/server/utils"
)

type UserProcess struct {
	// 字段........

}
// 写一个登陆函数
func (this *UserProcess) Login(userId int, userPwd string) (err error) {

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
	tf := &utils.Transfer{
		Conn: conn,
	}

	// 7. 已经拿到了要送的data
	_ = tf.WriteRkg(data)

	// 这里需要处理服务器的反馈信息
	mes, err = tf.ReadPkg()	// mes 就是
	if err != nil {
		fmt.Println("readPkg 出错了！！", err)
	}

	// 将mess反序列化为LoginRes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		// 初始化CurUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline

		// 显示当前在线列表
		fmt.Println("当前在线的用户id")
		for _, v := range loginResMes.UserIds {
			// 显示自己
			if v == userId {
				continue
			}
			fmt.Println("用户id:\t", v)
			// 完成客户端的初始化
			user := &message.User{
				UserId: v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Print("\n\n")
		// 这里启动一个协程
		// 该协程保持和服务端的通讯，如果服务器有数据推送给客户端
		// 接收并显示在终端
		go serverProcessMes(conn)

		for {
			ShowMenu()
		}
	}else{
		fmt.Println(loginResMes.Error)
	}

	return 
}

func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error){

	// 1. 链接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("register net dial err ", err)
		return
	}
	// 延时关闭
	defer conn.Close()

	// 2. 准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.RegisterMesType
	// 3.创建了一个LoginMes 结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	// 5.将用户信息进行序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err ", err)
		return
	}

	// 6.将用户序列化信息字符串化
	mes.Data = string(data)

	// 7.将发送信息进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json marshal err ", err)
		return
	}

	// 7.创建一个Transfer实例
	tf := &utils.Transfer{
		Conn: conn,
	}

	// 发送数据给服务器
	err = tf.WriteRkg(data)
	if err != nil {
		fmt.Println("register write pkg err ", err)
	}

	mes, err = tf.ReadPkg()

	if err != nil {
		fmt.Println("register read pkg err ", err)
		return
	}

	// 将mess反序列化为LoginRes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，你可以登陆了。")
	}else{
		fmt.Println(registerResMes.Error)
	}
	return
}