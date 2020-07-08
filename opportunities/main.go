package main

import "projects/investorsmarket/opportunities/service"

func main() {

	server := service.NewServer()
	server.Run(":3002")
}
