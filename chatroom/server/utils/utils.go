package utils

import (
	"net"
	"go_project/chatroom/common/message"
	"fmt"
	"encoding/binary"
	"encoding/json"
)

// 这里将这些方法关联到结构体
type Transfer struct {
	// 分析它应该有哪些字段
	Conn net.Conn
	Buf [8064]byte  //传输时使用的缓冲

}

func (this *Transfer) readPkg(conn net.Conn) (mes message.Message, err error) {

	buf := make([]byte, 8096)
	fmt.Print("读取客户端发送的数据......")
	_, err = conn.Read(buf[:4])
	if err != nil {
		// err = errors.New("read pkg header error")
		return
	}
	
	// 根据buf[:4] 转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])

	// 根据pkgLen读取消息内容 
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		// err = errors.New("read pkg body error")
		return
	}

	// 把pkgLen 反序列化为 -》message.Message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("unmarshall err ", err)
		return
	}

	return 
}

func (this *Transfer) writeRkg(conn net.Conn, data []byte) (err error) {

	// 先发送一个包的长度，发给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	// 发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn write(len) err ", err)
		return 
	}
	fmt.Printf("客户端发送消息的长度 %d, 内容 %s", len(data), string(data))

	// 发送消息本身
	n, err = conn.Write(data)
	if err != nil || n != int(pkgLen) {
		fmt.Println("conn write(content) err ", err)
		return 
	}
	return 
}