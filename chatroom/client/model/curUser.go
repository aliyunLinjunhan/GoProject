package model

import (
	"net"
	"go_project/chatroom/common/message"
)

//
type CurUser struct {

	Conn net.Conn
	message.User
}
