package infra

import (
	"log"
	"os"

	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
)

// SetupDBEngine DB Engine の初期化
func SetupDBEngine(logLevel logrus.Level) (*xorm.EngineGroup, error) {
	writerDbOptions := LoadMySQLWriterConfigEnv()
	log.Printf("MySQLの書き込み専用パス: %v", writerDbOptions.FormatDSN())
	writerEngine, err := InitMySQLEngine(writerDbOptions)
	if err != nil {
		return nil, err
	}
	readerDbOptions := LoadMySQLReaderConfigEnv()
	log.Printf("MySQLの読み取り専用パス: %v", readerDbOptions.FormatDSN())
	readerEngine, err := InitMySQLEngine(readerDbOptions)
	if err != nil {
		return nil, err
	}
	engine, err := xorm.NewEngineGroup(writerEngine, []*xorm.Engine{readerEngine})
	if err != nil {
		return nil, err
	}

	loggerSQL := logrus.New()
	loggerSQL.Level = logLevel
	loggerSQL.Out = os.Stdout
	loggerSQL.Formatter = &logrus.JSONFormatter{}

	engine.SetLogger(xorm.NewSimpleLogger(loggerSQL.Writer()))

	return engine, nil
}
