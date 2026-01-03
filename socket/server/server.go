package server

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
)

type Server struct {
	ip, port string
}

func NewServer(ip string, port string) *Server {
	return &Server{ip: ip, port: port}
}

func (s *Server) Run() {

	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.ip, s.port))
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept err : ", err)
			continue
		}

		go s.Handle(conn)

	}

}

func (s *Server) Handle(conn net.Conn) {
	defer conn.Close()
	for {

		buff := make([]byte, 1024)
		reader := bufio.NewReader(conn)
		readSize, err := reader.Read(buff)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Println("读数据错误", err)
		}

		go s.Reader(buff, readSize)
		go s.Writer(conn)

	}
}

func (s *Server) Reader(buff []byte, size int) {

	fmt.Println(string(buff), size)
}

func (s *Server) Writer(conn net.Conn) {
	conn.Write([]byte("1"))
}
