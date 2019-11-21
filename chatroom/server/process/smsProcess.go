package process

import (
	"encoding/json"
	"fmt"
	"go_project/chatroom/client/utils"
	"go_project/chatroom/common/message"
	"net"
)

type SmsProcess struct{
	// 暂时不需要
}

// 写方法转发消息
func (this *SmsProcess) SendGroupMes(mes *message.Message) {

	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json marshal err ", err)
	}
	// 遍历服务端的onlineUser map[int]*UserProcess
	for id, up := range userMgr.onlineUsers {
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
}

func (this *SmsProcess) SendMesToEachOnlineUser(info []byte, conn net.Conn) {

	// 创建一个transfer
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WriteRkg(info)
	if err != nil {
		fmt.Println("write pkg err", err)
	}
}


