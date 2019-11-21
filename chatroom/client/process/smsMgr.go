package process

import (
	"encoding/json"
	"fmt"
	"go_project/chatroom/common/message"
)

func outputGroupMes(mes *message.Message) {

	// 显示即可
	// 1. 反序列化mesData
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json unmarshal err ", err)
		return
	}

	// 显示信息
	fmt.Printf("\n用户id：%d \n对大家说：\t%s", smsMes.UserId, smsMes.Content)
	fmt.Println()

}

