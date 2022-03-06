package main

import (
	"log"
	"os"
	"sample-docker-gin/internal/server"
)

type TableSample struct {
	Id   int    `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

func main() {

	err := server.Start()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
