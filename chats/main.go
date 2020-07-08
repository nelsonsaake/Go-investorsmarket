package main

import "projects/investorsmarket/chats/service"

func main() {

	server := service.NewServer()
	server.Run(":3004")
}
