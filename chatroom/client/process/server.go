package process

import (
	"encoding/json"
	"fmt"
	"go_project/chatroom/common/message"
	"go_project/chatroom/server/utils"
	"net"
	"os"
)

// 显示登陆成功后的界面.....................
func ShowMenu() {
	fmt.Println("-------------恭喜登陆成功----------------------")
	fmt.Println("-------------1、显示在线用户列表----------------")
	fmt.Println("-------------2、发送消息----------------------")
	fmt.Println("-------------3、信息列表-----------------------")
	fmt.Println("-------------4、退出系统-----------------------")
	fmt.Println("清选择(1-4):")
	var key int
	var content string
	fmt.Scanf("%d\n", &key)
	smsProcess := &SmsProcess{}
	switch key {
		case 1:
			outputOnlineUser()
		case 2:
			fmt.Println("请输入你要对大家说的话:")
			fmt.Scanf("%s\n", &content)
			_ = smsProcess.SendGroupMes(content)
		case 4:
			fmt.Println("你选择了退出系统.........")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误！！！！")
	}
}

// 和服务器保持通讯
func serverProcessMes(conn net.Conn) {

	// 创建一个transfer实例不停的读取服务器发送的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("客户端持续读取服务端的链接出错了", err)
			return
		}
		//如果读到消息
		switch mes.Type {
			case message.NotifyUserStatusMesType :
			// 1.取出NotifyUserStatusMes
				var notifyUserStatusMes message.NotifyUserStatusMes
				_ = json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			// 2. 把这个用户的消息，状态保存到客户Map中
				updateUserStatus(&notifyUserStatusMes)
			// 处理
			case message.SmsMesType:
				outputGroupMes(&mes)

			default: fmt.Println("服务器返回未知的消息")
		}
		//fmt.Printf("mes=%v", mes)
	}
}