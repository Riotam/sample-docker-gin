package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sample-docker-gin/internal/util"
)

// Sample はスケルトンコード
func Sample(c *gin.Context) {
	res := util.Sample(1, 2)

	c.JSON(http.StatusOK, res)
}
