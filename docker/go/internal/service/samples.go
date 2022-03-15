package service

import (
	"fmt"
	"sample-docker-gin/internal/model"
	"sample-docker-gin/internal/repository"
)

type SamplesInterface interface {
	GetOne(id int) (*model.Sample, error)
}

// Samples はサンプルのサービス実装
type Samples struct {
	samplesRepo repository.SamplesInterface
}

// NewSamples はサンプルのサービスを初期化
func NewSamples(samplesRepo repository.SamplesInterface) *Samples {
	u := Samples{
		samplesRepo: samplesRepo,
	}
	return &u
}

func (s *Samples) GetOne(id int) (*model.Sample, error) {
	fmt.Print("pass service!")
	return s.samplesRepo.GetOne(id)
}
