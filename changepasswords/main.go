package main

import "projects/investorsmarket/changepasswords/service"

func main() {

	server := service.NewServer()
	server.Run(":3005")
}
