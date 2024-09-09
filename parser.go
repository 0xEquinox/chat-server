package main

import "fmt"

func parseCommand(message []byte) (Command, error) {
	if len(message) < 1 {
		return Command{}, fmt.Errorf("Message too short")
	}

	// First byte of the message is the type
	ctype := CommandType(message[0])
	payload := message[1:]

	switch ctype {
	case Capabilities:
		return Command{ctype: Capabilities, payload: payload}, nil

	case CreateRoom:
		return Command{ctype: CreateRoom, payload: payload}, nil

	case Join:
		return Command{ctype: Join, payload: payload}, nil

	case Leave:
		return Command{ctype: Leave, payload: payload}, nil

	case Send:
		return Command{ctype: Send, payload: payload}, nil

	case Subscribe:
		return Command{ctype: Subscribe, payload: payload}, nil

	case Unsubscribe:
		return Command{ctype: Unsubscribe, payload: payload}, nil

	case RoomInfo:
		return Command{ctype: RoomInfo, payload: payload}, nil

	case Delete:
		return Command{ctype: Delete, payload: payload}, nil

	case Edit:
		return Command{ctype: Edit, payload: payload}, nil

	case Create:
		return Command{ctype: Create, payload: payload}, nil

	case Acknowledge:
		return Command{ctype: Acknowledge, payload: payload}, nil

	default:
		return Command{}, fmt.Errorf("unknown command type: %d", ctype)
	}
}
