package main

import "projects/investorsmarket/posts/service"

func main() {

	server := service.NewServer()
	server.Run(":3006")
}
