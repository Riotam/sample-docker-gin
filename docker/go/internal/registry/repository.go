package registry

import (
	"github.com/go-xorm/xorm"
	//"sample-docker-gin/internal/registry"
	//"github.com/jinzhu/gorm"
	"sample-docker-gin/internal/repository"
)

const (
	// RepositoryKey はリポジトリレジストリ取得キー名
	RepositoryKey = "repository_registry"
)

// RepositoryMaker はリポジトリレジストリ
type RepositoryMaker interface {
	// MASTER
	NewSamples() repository.SamplesInterface
}

// RepositorySettings はリポジトリレジストリの設定
// 全ての設定が必須
type RepositorySettings struct {
	Engine xorm.EngineInterface
}

// Repository はリポジトリレジストリの実装
// インフラ層の依存情報を初期化時に注入する
type Repository struct {
	settings *RepositorySettings
}

// NewRepository initializes factory with injected infra.
func NewRepository(settings *RepositorySettings) RepositoryMaker {
	r := &Repository{
		settings: settings,
	}
	return r
}

// NewSamples returns samples repository.
func (r *Repository) NewSamples() repository.SamplesInterface {
	samplesRepo := repository.NewSamples(r.settings.Engine)
	return samplesRepo
}
