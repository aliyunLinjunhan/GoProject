package main
import (
	"fmt"
	"net"
	"go_project/chatroom/common/message"
	"encoding/binary"
	"encoding/json"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {

	buf := make([]byte, 8096)
	fmt.Print("读取服务端发送的数据......")
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

func writeRkg(conn net.Conn, data []byte) (err error) {

	// 7.1 先把data长度发送给服务器
	// 先获取data的长度 -》 转成一个表示长度的bytes的切片

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