package repository

import (
	"github.com/go-xorm/xorm"
	"sample-docker-gin/internal/model"
	"sample-docker-gin/internal/util"
)

// SamplesInterface has users data.
type SamplesInterface interface {
	GetOne(id int) (*model.Sample, error)
}

// Samples has users data.
type Samples struct {
	engine xorm.EngineInterface
}

// NewSamples initializes Users
func NewSamples(engine xorm.EngineInterface) *Samples {
	u := Samples{
		engine: engine,
	}
	return &u
}

func (s *Samples) GetOne(id int) (*model.Sample, error) {

	output := model.Sample{}

	session := s.engine.NewSession()
	session.Table("samples")
	session.Where("id = ?", id)
	_, err := session.Get(&output)
	if err != nil {
		util.GetLogger().Error(err)
		return nil, err
	}

	// 取得結果0の場合もエラーとする
	//if !found {
	//	return nil, model.NewError(model.ErrorResourceNotFound, "me not found")
	//}

	return &output, nil
}
