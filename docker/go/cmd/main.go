package main

import (
	"log"
	"os"
	"sample-docker-gin/internal/server"
)

func main() {

	err := server.Start()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
