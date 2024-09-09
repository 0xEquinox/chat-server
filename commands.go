package main

type CommandType int

const (
	Capabilities CommandType = iota
	CreateRoom
	Join
	Leave
	Send
	Subscribe
	Unsubscribe
	RoomInfo
	Delete
	Edit
	Create
	Acknowledge
)

type Command struct {
	ctype   CommandType
	payload []byte
}
