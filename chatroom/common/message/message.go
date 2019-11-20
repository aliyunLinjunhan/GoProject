package message

const (
	LoginMesType = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
	RegisterResMesType = "RegisterResMes"
)

type Message struct {

	Type string `json:"type"`  // 消息类型
	Data string `json:"data"`   // 消息的数据
}

// 定义俩个消息

type LoginMes struct {

	UserId int `json:"userId"` // 用户id
	UserPwd string `json:"userPwd"`  // 用户密码
	UserName string `json:"userName"` // 用户名
}

type LoginResMes struct {

	Code int `json:"code"`  // 返回状态码，500表示该用户未注册 200表示登陆成功
	Error string `json:"error"`  // 返回错误信息
	UserIds []int		// 增加字段，保存用户id的切片
}

type RegisterMes struct {

	User User `json:"user"`
}

type RegisterResMes struct {

	Code int `json:"code"`  // 返回码500表示该用户未注册， 200表示注册成功
	Error string `json:"error"`  // 返回错误信息
}
