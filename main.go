package main

import (
	"fmt"
	"net"
)

type Message struct {
	sender string
	text   []byte
}

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
	msgch      chan Message
}

func newServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
		msgch:      make(chan Message, 10),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)

	if err != nil {
		return err
	}

	defer ln.Close()
	s.ln = ln

	go s.acceptLoop()

	<-s.quitch
	close(s.msgch)

	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()

		if err != nil {
			fmt.Println("An Error Occured While Making a New Connection: ", err)
			continue
		}

		fmt.Println("New connection established: ", conn.RemoteAddr().String())

		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 2048)

	for {
		n, err := conn.Read(buf)

		if err != nil {
			fmt.Println("Error reading connection: ", err)
			continue
		}

		s.msgch <- Message{
			sender: conn.RemoteAddr().String(),
			text:   buf[:n],
		}
	}
}

func main() {
	server := newServer(":3000")
	go func() {
		for msg := range server.msgch {
			fmt.Printf("%s: %s\n", msg.sender, string(msg.text))
		}
	}()

	server.Start()

}
