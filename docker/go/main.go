package main

import (
	"os"
	"sample-docker-gin/internal/server"
	"sample-docker-gin/internal/util"
)

func main() {
	util.LoadEnv()
	util.InitLogger()
	err := server.Start()
	if err != nil {
		util.GetLogger().Errorln(err)
		os.Exit(1)
	}
}
