package main

import "projects/investorsmarket/users/service"

func main() {

	server := service.NewServer()
	server.Run(":3001")
}
