package process

import "fmt"

// 定义一个全局变量跟踪在线用户
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

// 对userMgr进行初始化
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// 完成对onlineUser添加
func (this *UserMgr) AddOnlineUser(up *UserProcess) {

	this.onlineUsers[up.UserId] = up
}

// 删除
func (this *UserMgr) DelOnLineUser(userId int) {

	delete(this.onlineUsers, userId)
}

// 获取制定id的链接
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {

	// 如何从map取出一个值，带检测方式
	up, ok := this.onlineUsers[userId]
	if !ok {
		// 要查找的用户不在线
		err = fmt.Errorf("用户%d 不存在", userId)
		return
	}
	return
}

// 返回在线列表
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {

	return this.onlineUsers
}
