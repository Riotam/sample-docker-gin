package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"sample-docker-gin/internal/infra"
	"sample-docker-gin/internal/registry"
	"sample-docker-gin/internal/util"
	//"sample-docker-gin/internal/infra"
)

func loadServiceFactorySettings() *registry.ServiceFactorySettings {
	// service factoryの初期化
	serviceFactorySettings := &registry.ServiceFactorySettings{}
	return serviceFactorySettings
}

// Start starts api server
func Start() error {

	logger := util.GetLogger()

	// ログの出力設定
	logLevel, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = logrus.DebugLevel
		logger.Warnf("LOG_LEVEL is not set.")
	}

	// db engine 初期化
	engine, err := infra.SetupDBEngine(logLevel)
	if err != nil {
		logger.Fatal(err)
	}
	defer func() {
		log.Println("engine closed")
		engine.Close()
	}()

	loggerAccess := logrus.New()
	loggerAccess.Level = logLevel
	loggerAccess.Out = os.Stdout
	loggerAccess.Formatter = &logrus.JSONFormatter{}

	loggerError := logrus.New()
	loggerError.Level = logLevel
	loggerError.Out = os.Stderr
	loggerError.Formatter = &logrus.JSONFormatter{}
	gin.DefaultErrorWriter = loggerError.Writer()

	serviceFactorySettings := loadServiceFactorySettings()
	serviceFactorySettings.Engine = engine

	// Ginの初期化
	r := gin.Default()

	// route
	DefineRoutes(r)

	r.Run(":3000")

	return nil
}
