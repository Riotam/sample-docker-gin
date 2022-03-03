package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	//_ "github.com/go-sql-driver/mysql"
	//"github.com/jinzhu/gorm"
	"net/http"
	"github.com/Riotam/sample-docker-gin/cmd/app/mysql"
)

type TableSample struct {
	Id   int    `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

//func GormConnect() *gorm.DB {
//	DBMS := "mysql"
//	USER := "docker"
//	PASS := "docker"
//	PROTOCOL := "tcp(db:3306)"
//	DBNAME := "sample"
//
//	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
//	db, err := gorm.Open(DBMS, CONNECT)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	return db
//}

func main() {

	// MySQLとの接続
	db := mysql.GormConnect()
	defer db.Close()

	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {

		// データ取得
		var tableSample []TableSample
		db.Table("table_sample").Find(&tableSample)

		fmt.Println("tableSample: ", tableSample)

		c.JSON(http.StatusOK, gin.H{
			"message": "success!",
			"data":    tableSample,
		})
	})
	engine.Run(":3000")
}
