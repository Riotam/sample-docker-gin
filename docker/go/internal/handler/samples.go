package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sample-docker-gin/internal/util"
	"strconv"
)

// Sample はスケルトンコード
func Sample(c *gin.Context) {

	aStr := c.Query("a")
	bStr := c.Query("b")

	aInt, err := strconv.Atoi(aStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	bInt, err := strconv.Atoi(bStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	res := util.Sample(aInt, bInt)

	c.JSON(http.StatusOK, res)
}
