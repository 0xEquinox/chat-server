package main

import (
	"fmt"
	"net"

	"github.com/vmihailenco/msgpack/v5"
)

type Room struct {
	id int
}

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
	msgch      chan Message
	rooms      []Room
}

func newServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
		msgch:      make(chan Message, 10),
		rooms:      make([]Room, 2048),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)

	if err != nil {
		return err
	}

	fmt.Println("Starting Server")
	defer ln.Close()
	s.ln = ln

	go s.acceptLoop()

	<-s.quitch
	close(s.msgch)

	return nil
}

func (s *Server) acceptLoop() {
	fmt.Println("Listening for connections")
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

		message := buf[:n]

		command, err := parseCommand(message)
		if err != nil {
			fmt.Println("Error parsing command: ", err)
			continue
		}

		s.msgch <- Message{
			sender:  conn.RemoteAddr().String(),
			command: command,
		}
	}
}

func (s *Server) handeCommands() {
	for msg := range s.msgch {
		switch msg.command.ctype {
		case Capabilities:
			// Handle the Capabilities command
			fmt.Println("Handling Capabilities command from", msg.sender)
			// Implement capabilities logic here

		case CreateRoom:
			// Handle the CreateRoom command
			roomName := msgpack.RawMessage(msg.command.payload)
			fmt.Printf("Creating room: %s by %s\n", roomName, msg.sender)
			// Implement room creation logic here

		case Join:
			// Handle the Join command
			roomName := string(msg.command.payload)
			fmt.Printf("%s is joining room: %s\n", msg.sender, roomName)
			// Implement join logic here

		case Leave:
			// Handle the Leave command
			roomName := string(msg.command.payload)
			fmt.Printf("%s is leaving room: %s\n", msg.sender, roomName)
			// Implement leave logic here

		case Send:
			// Handle the Send command
			messageText := string(msg.command.payload)
			fmt.Printf("Message from %s: %s\n", msg.sender, messageText)
			// Implement message sending logic here

		case Subscribe:
			// Handle the Subscribe command
			channelName := string(msg.command.payload)
			fmt.Printf("%s is subscribing to channel: %s\n", msg.sender, channelName)
			// Implement subscribe logic here

		case Unsubscribe:
			// Handle the Unsubscribe command
			channelName := string(msg.command.payload)
			fmt.Printf("%s is unsubscribing from channel: %s\n", msg.sender, channelName)
			// Implement unsubscribe logic here

		case RoomInfo:
			// Handle the RoomInfo command
			roomName := string(msg.command.payload)
			fmt.Printf("Request for info about room: %s from %s\n", roomName, msg.sender)
			// Implement room info logic here

		case Delete:
			// Handle the Delete command
			itemID := string(msg.command.payload)
			fmt.Printf("%s is deleting item: %s\n", msg.sender, itemID)
			// Implement delete logic here

		case Edit:
			// Handle the Edit command
			editInfo := string(msg.command.payload)
			fmt.Printf("%s is editing item: %s\n", msg.sender, editInfo)
			// Implement edit logic here

		case Create:
			// Handle the Create command
			createInfo := string(msg.command.payload)
			fmt.Printf("%s is creating item: %s\n", msg.sender, createInfo)
			// Implement create logic here

		case Acknowledge:
			// Handle the Acknowledge command
			ackInfo := string(msg.command.payload)
			fmt.Printf("%s is acknowledging: %s\n", msg.sender, ackInfo)
			// Implement acknowledge logic here

		default:
			fmt.Printf("Unknown command from %s\n", msg.sender)
		}
	}
}
