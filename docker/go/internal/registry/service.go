package registry

import (
	"github.com/go-xorm/xorm"
	"sample-docker-gin/internal/repository"
	//"sample-docker-gin/internal/registry"
	"sample-docker-gin/internal/service"
)

const (
	// ServiceKey is ServiceKey
	ServiceKey = "service_factory"
)

// ServiceMaker はサービスファクトリ
type ServiceMaker interface {
	NewSamples() service.SamplesInterface
}

// ServiceFactorySettings はサービスファクトリの設定
// 全ての設定が必須
type ServiceFactorySettings struct {
	Engine xorm.EngineInterface
}

// ServiceFactory はサービスファクトリの実装
// インフラ層の依存情報を初期化時に注入する
type ServiceFactory struct {
	settings *ServiceFactorySettings
}

// NewSamples returns Samples service.
func (r *ServiceFactory) NewSamples() service.SamplesInterface {
	samplesRepo := repository.NewSamples(r.settings.Engine)
	return service.NewSamples(samplesRepo)
}
