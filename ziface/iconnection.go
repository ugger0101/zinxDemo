package ziface

import "net"

// 定义连接接⼝口
type IConnection interface {
	//启动连接，让当前连接开始⼯工作
	Start()
	//停⽌止连接，结束当前连接状态M
	Stop()
	//从当前连接获取原始的socket TCPConn GetTCPConnection() *net.TCPConn

	//获取当前连接ID
	GetConnID() uint32

	//获取当前连接绑定socket conn
	GetTCPConnection() *net.TCPConn

	// 获取远程客户端的TCP状态 IP port
	RemoterAddr() net.Addr

	// 发送数据，将数据发送给远程的客户端
	Send(data []byte) error
}

type HandFunc func(*net.TCPConn, []byte, int) error
