package util

import (
	"fmt"
	"io"
	"os"

	"github.com/go-pp/pp"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var globalLogger *logrus.Logger

// InitLogger Loggerについて初期化する
func InitLogger() {
	globalLogger = nil
	pp.ColoringEnabled = false
}

// GetLogger returns global logger
func GetLogger() *logrus.Logger {

	if globalLogger != nil {
		return globalLogger
	}

	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	// Output to stdout instead of the default stderr
	// log.SetOutput(os.Stdout)

	logger.Out = os.Stdout

	level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logger.Warningln(err)
		level = logrus.DebugLevel
	}
	logger.Level = level
	logger.Infof("Global Log Level is %v", level.String())

	globalLogger = logger

	return logger
}

// WriteError エラー情報を出力します。
// `LOG_STACK_TRACE` が設定されているときのみ、追加でスタックトレースを標準出力に出力します。
// それ以外の場合は、エラーメッセージのみを出力します。
func WriteError(err error) {
	logger := GetLogger()
	logger.Error(err)

	if os.Getenv("LOG_STACK_TRACE") != "" {
		logger.Error(fmt.Sprintf("%+v", errors.WithStack(err)))
	}
}

// DebugLogWriter ログ出力 io.Writer
func DebugLogWriter() *io.PipeWriter {
	return GetLogger().WriterLevel(logrus.DebugLevel)
}

// ErrorLogWriter ログ出力 io.Writer
func ErrorLogWriter() *io.PipeWriter {
	return GetLogger().WriterLevel(logrus.ErrorLevel)
}
