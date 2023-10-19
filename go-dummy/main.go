package main

import "go-dummy-project/go-dummy/server"

func main() {

	// channel := make(chan struct{})
	// go config.GetRedisClient(channel)
	// <-channel

	server.StartServer()
}
