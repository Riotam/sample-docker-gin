package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sample-docker-gin/internal/registry"
)

// SamplesSample はスケルトンコード
func SamplesSample(c *gin.Context) {

	fmt.Println("L13!!", "", "")

	// TODO: ここでパニック起こす
	serviceMaker := c.MustGet(registry.ServiceKey).(registry.ServiceMaker)
	samplesService := serviceMaker.NewSamples()
	sampleId := c.MustGet("sample_id").(int)

	output, err := samplesService.GetOne(sampleId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, output)
}
