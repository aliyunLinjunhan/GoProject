package process

import (
	"encoding/json"
	"fmt"
	"go_project/chatroom/client/utils"
	"go_project/chatroom/common/message"
)

type SmsProcess struct {

}

func (this *SmsProcess) SendGroupMes(content string) (err error) {

	// 1. 创建一个消息体
	var mes message.Message
	mes.Type = message.SmsMesType

	// 2. 创建消息内容
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	// 3. 序列化
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail = ", err.Error())
	}
	mes.Data = string(data)

	// 4.对mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("sendGroupMes json.Marshal fail = ", err.Error())
		return
	}

	// 5. 将mes发送给服务器
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}

	// 6.发送
	err = tf.WriteRkg(data)
	if err != nil {
		fmt.Println("write pkg err ", err)
		return
	}
	return
}