package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"sample-docker-gin/internal/apps/app/mysql"
)

type TableSample struct {
	Id   int    `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

func main() {

	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {

		// MySQLとの接続
		db := mysql.GormConnect()
		defer db.Close()

		//データ取得
		var tableSample []TableSample
		db.Table("table_sample").Find(&tableSample)

		c.JSON(http.StatusOK, gin.H{
			"message": "success!",
			"data":    tableSample,
		})
	})
	engine.Run(":3000")
}
