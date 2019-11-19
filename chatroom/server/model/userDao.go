package model

import (
	"fmt"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
)

// 在服务器启动后，就初始化一个userDao实例
// 把它做成全局变量，需要和redis操作时直接使用即可
var (
	MyUserDao *UserDao
)

// 定义一个UserDao 结构体
// 完成对结构体的各种操作

type UserDao struct {
	Pool *redis.Pool
}

// 使用工厂模式
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {

	userDao = &UserDao{
		pool,
	}
	return 
}

// 1. 根据用户id返回一个User实例+err
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {

	// 通过id去调用redis查询用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		// 错误
		if err == redis.ErrNil{
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{}

	// 这里需要把res反序列化为User实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json unmarshal err ", err)
		return
	}
	return
}

// 2.完成登陆的校验
// 1.login 完成对用户的验证
// 2.如果用户的id或pwd都正确，则返回一个user实例和一个error
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {

	// 先从UserDao的连接池中取出一个链接
	conn := this.Pool.Get()
	defer conn.Close()

	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	// 这时证明这个用户是获取到
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

