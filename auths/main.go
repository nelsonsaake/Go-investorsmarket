package main

import "projects/investorsmarket/auths/service"

func main() {

	server := service.NewServer()
	server.Run(":3007")
}
