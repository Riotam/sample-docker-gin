package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sample-docker-gin/internal/handler"
)

// DefineRoutes ルートの設定
func DefineRoutes(r gin.IRouter) {

	// Health Check
	r.GET("/health", func(c *gin.Context) {
		type res struct {
			Status string `json:"status"`
		}
		r := &res{Status: "OK!"}
		c.JSON(http.StatusOK, r)
	})

	// sample
	r.GET("/sample", RequirePathParamStr("sample_id"), handler.SamplesSample)
}
