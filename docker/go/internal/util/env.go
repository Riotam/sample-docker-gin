package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type logFormat struct {
	Level string `json:"level"`
	Msg   string `json:"msg"`
	Time  string `json:"time"`
}

const (
	info           = "info"
	warning        = "warning"
	noEnvSpecified = "no env file specified. try to load default .env."
	noEnvFile      = "no env file loaded ."
	envFileLoaded  = "env file loaded: "
)

// LoadEnv apply .env to ENVIRONMENT VARIABLE.
func LoadEnv() {
	if _, found := os.LookupEnv("ENV_FILE"); !found {
		logMessage := makeJSONLogString(info, noEnvSpecified)
		log.SetOutput(os.Stdout)
		log.Println(logMessage)
		os.Setenv("ENV_FILE", "config/.env")
	}

	envfile := os.Getenv("ENV_FILE")
	fmt.Println(envfile)
	err := godotenv.Load(envfile)

	if err != nil {
		logMessage := makeJSONLogString(warning, fmt.Sprintf("%v%v", noEnvFile, err))
		log.SetOutput(os.Stderr)
		log.Println(logMessage)
	} else {
		logMessage := makeJSONLogString(info, fmt.Sprintf("%v%v", envFileLoaded, envfile))
		log.SetOutput(os.Stdout)
		log.Println(logMessage)
	}
}

// makeJSONLogString はログ出力用のjsonを返します
func makeJSONLogString(level string, message string) (logMessageJSON string) {
	l := logFormat{
		Level: "",
		Msg:   "",
		Time:  time.Now().Format(time.RFC3339),
	}

	l.Level = level
	l.Msg = message
	jsonLog, err := json.Marshal(l)
	if err == nil {
		return string(jsonLog)
	}
	return ""
}
