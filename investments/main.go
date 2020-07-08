package main

import "projects/investorsmarket/investments/service"

func main() {

	server := service.NewServer()
	server.Run(":3003")
}
