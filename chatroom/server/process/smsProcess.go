package process

// 编写ServerProcessMes 函数根据不同的消息种类，决定调用那个函数
func ServerProcessMes(conn net.Conn, mes *message.Message) (err error) {

	switch mes.Type {
		case message.LoginMesType:
			// 处理登陆逻辑
			err = serverProcessLogin(conn, mes)
		case message.RegisterMesType:
			//
		default: 
			fmt.Println("消息类型不存在.......................")
	}
	return
}