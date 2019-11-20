package process
import (
	"fmt"
	"net"
	"go_project/chatroom/common/message"
	"go_project/chatroom/server/utils"
	"go_project/chatroom/server/model"
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
		fmt.Println("json unmarshal fail err ", err)
		return
	}

	// 1.先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	// 2.再声明一个 LoginResMes 
	var LoginResMes message.LoginResMes
	// 到redis数据库完成用户验证
	// 使用MyUserDao.Login到redis中去验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			LoginResMes.Code = 500
			LoginResMes.Error = err.Error()
		}else if err == model.ERROR_USER_PWD {
			LoginResMes.Code = 403
			LoginResMes.Error = err.Error()
		}else {
			LoginResMes.Code = 300
			LoginResMes.Error = "服务器内部错误"
		}
		//测试成功
	} else {
		LoginResMes.Code = 200
		fmt.Println(user, "登陆成功")
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

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error){

	// 1. 先从mes中取出mes.Data，并直接反序列化LoginMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json unmarshal fail err ", err)
		return
	}

	// 1.先声明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	// 2.再声明一个 registerResMes
	var registerResMes message.RegisterResMes

	// 到redis数据库完成用户注册验证
	// 使用MyUserDao.Register到redis中去验证
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		}else {
			registerResMes.Code = 506
			registerResMes.Error = "发生未知错误"
		}
		//测试成功
	} else {
		registerResMes.Code = 200
		fmt.Println("注册成功")
	}

	//3.将registerResMes序列化
	data, err := json.Marshal(registerResMes)
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
