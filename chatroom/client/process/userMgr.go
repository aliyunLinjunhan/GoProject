package process

import (
	"go_project/chatroom/common/message"
	"fmt"
)

// 客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)

// 在客户端显示在线的用户
func outputOnlineUser() {
	// 遍历一把 onlineUsers
	fmt.Println("当前用户列表:")
	for id, _ := range onlineUsers{
		fmt.Println("用户id:\t", id)
	}
}

func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {

	// 适当优化
	user, ok := onlineUsers[notifyUserStatusMes.UserId]

	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user

	outputOnlineUser()
}

