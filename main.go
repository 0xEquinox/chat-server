package main

type Message struct {
	sender  string
	command Command
}

func main() {
	server := newServer(":3000")
	go server.handeCommands()
	server.Start()
}
