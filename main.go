package main

import (
	"log"
	"os"
	"myprojet/server"
)

func main() {
	arg := os.Args
	if len(arg) != 1 {
		log.Fatal("Please provide one argument")
	}
	
	server.Server()
}
