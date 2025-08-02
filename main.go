package main

import (
	"log"
	"myprojet/server"
	"os"
)

func main() {
	arg := os.Args
	if len(arg) != 1 {
		log.Fatal("Please provide one argument")
	}
	server.Server()
}
