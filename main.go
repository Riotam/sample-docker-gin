package main

import "github.com/gin-gonic/gin"

import "net/http"

func main() {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world! I'm PON!",
		})
	})
	engine.Run(":3000")
}
