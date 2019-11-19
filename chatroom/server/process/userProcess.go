package process
import (
	"fmt"
	"net"
	"go_project/chatroom/common/message"
	"go_project/chatroom/server/utils"
	"encoding/json"

)

type UserProcess struct {
	// 字段?
	Conn net.Conn

}

// 编写一个ServerProcessLogin函数，专门处理登陆请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {

	// 1. 先从mes中取出mes.Data，并直接反序列化LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json umarshal fail err ", err)
		return
	}

	//1.先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//2.再声明一个 LoginResMes 
	var LoginResMes message.LoginResMes


	// 如果用户id=100，密码=123456， 认为合法， 否则不合法
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		// 合法
		LoginResMes.Code = 200
	}else{
		// 不合法
		LoginResMes.Code = 500
		LoginResMes.Error = "该用户不存在，请确认账号密码是否正确！！"
	}

	//3.将loginResMes序列化
	data, err := json.Marshal(LoginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail ", err)
		return
	}

	//4. 将data赋值给resMes
	resMes.Data = string(data)

	//5. 对resMes进行序列化，准备发送
	data, err = json.Marshal(resMes)

	//6. 发送包(进行封装)
	// 因为使用了分层
	tf := &utils.Transfer{
		Conn : this.Conn,
	}
	err = tf.WriteRkg(data)
	return 

}
