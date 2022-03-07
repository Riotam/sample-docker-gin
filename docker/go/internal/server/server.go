package server

import (
	"github.com/gin-gonic/gin"
	//"sample-docker-gin/internal/infra"
)

// Start starts api server
func Start() error {

	// TODO: ここじゃないかも？
	// db engine 初期化
	//db := infra.GormConnect()
	//defer db.Close()

	// Ginの初期化
	r := gin.Default()

	// route
	DefineRoutes(r)

	r.Run(":3000")

	return nil
}
