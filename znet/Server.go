package znet

import "zinxDemo/ziface"

type Server struct {
	// 服务器name
	Name string

	IPVersion string

	IP string

	Port int
}

func (s *Server) Start() {

}
func (s *Server) Stop() {

}
func (s *Server) Serve() {

}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}
