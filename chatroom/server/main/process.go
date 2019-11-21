package main
import (
	"net"
	"go_project/chatroom/common/message"
	"go_project/chatroom/server/utils"
	"go_project/chatroom/server/process"
	"fmt"
	"io"
)

type Processor struct {
	Conn net.Conn

}

// 编写ServerProcessMes 函数根据不同的消息种类，决定调用那个函数
func (this *Processor) ServerProcessMes(mes *message.Message) (err error) {

	fmt.Println("读到对方发送的包..", mes.Data)

	switch mes.Type {
		case message.LoginMesType:
			// 处理登陆逻辑
			up := &process.UserProcess{
				Conn : this.Conn,
			}
			err = up.ServerProcessLogin(mes)
		case message.RegisterMesType:
			// 处理注册逻辑
			up := &process.UserProcess{
				Conn:this.Conn,
			}
			err = up.ServerProcessRegister(mes)
		case message.SmsMesType:
			// 去转发群聊消息
			smsProcess := &process.SmsProcess{}
			smsProcess.SendGroupMes(mes)
	default:
			fmt.Println("消息类型不存在.......................")
	}
	return
}

func (this *Processor) Process2() (err error){
	// 循环的客户端发送的消息
	for {

		// 这里将读取数据包， 直接进行封装
		tf := &utils.Transfer{
			Conn : this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端推出了，服务端也正常退出...")
				return err
			} else {
				fmt.Println("readPkg err =", err)
				return err
			}
		}
		// fmt.Println("mes=", mes)
		err = this.ServerProcessMes(&mes)
		if err != nil {
			return err
		}

	}
}