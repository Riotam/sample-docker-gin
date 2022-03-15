package server

import "github.com/gin-gonic/gin"

// RequirePathParamStr parses PathParameter as uint64 by given param and then sets it to gin Context.
func RequirePathParamStr(param string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(param, c.Param(param))
		c.Next()
	}
}
