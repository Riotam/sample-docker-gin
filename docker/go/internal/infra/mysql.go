package infra

import (
	"fmt"
	"os"
	"sample-docker-gin/internal/util"
	"strconv"
	"strings"
	"time"

	"github.com/go-pp/pp"
	"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

const (
	databaseConnectionMaxLifeTimeEnv        = "DB_CONNECTION_MAX_LIFETIME"
	databaseConnectionMaxLifeTimeDefaultSec = 1
)

// LoadMySQLWriterConfigEnv initializes MySQL writer config using Environment Variables.

func LoadMySQLWriterConfigEnv() *mysql.Config {
	params := map[string]string{
		"charset": "utf8mb4",
	}

	conf := &mysql.Config{
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_WRITER_HOST"),
		DBName:               os.Getenv("DB_WRITER_NAME"),
		User:                 os.Getenv("DB_WRITER_USER"),
		Passwd:               os.Getenv("DB_WRITER_PASSWORD"),
		Params:               params,
		AllowNativePasswords: true,
		// Loc:                  util.TokyoTimeLocation,
		Loc: time.UTC,
		// Aurora の Cluster エンドポイントのフェイルオーバー時にRead-Replicaに接続してしまう場合の対策
		RejectReadOnly: true,
	}

	pp.ColoringEnabled = false
	pp.Println("DATABASE HOST ADDR=", conf.Addr)

	return conf
}

// LoadMySQLReaderConfigEnv initializes MySQL reader config using Environment Variables.
func LoadMySQLReaderConfigEnv() *mysql.Config {
	params := map[string]string{
		"charset": "utf8mb4",
	}

	conf := &mysql.Config{
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_READER_HOST"),
		DBName:               os.Getenv("DB_READER_NAME"),
		User:                 os.Getenv("DB_READER_USER"),
		Passwd:               os.Getenv("DB_READER_PASSWORD"),
		Params:               params,
		AllowNativePasswords: true,
		// Loc:                  util.TokyoTimeLocation,
		Loc: time.UTC,
	}

	return conf
}

// InitMySQLEngine initialize xorm engine for mysql
func InitMySQLEngine(conf *mysql.Config) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("mysql", conf.FormatDSN())
	if err != nil {
		return nil, err
	}

	charset, ok := conf.Params["charset"]
	if !ok {
		charset = "utf8mb4"
	}

	// DBのタイムゾーンは日本時間
	engine.SetTZDatabase(time.FixedZone("Asia/Tokyo", 9*60*60))
	engine.SetTZLocation(time.FixedZone("Asia/Tokyo", 9*60*60))
	engine.Charset(charset)
	engine.SetMapper(core.GonicMapper{})
	engine.StoreEngine("InnoDb")
	showSQL := os.Getenv("SHOW_SQL")
	if showSQL == "0" || showSQL == "false" {
		engine.ShowSQL(false)
	} else {
		engine.ShowSQL(true)
	}

	var connMaxLifeTime int
	connMaxLifeTimeStr := os.Getenv(databaseConnectionMaxLifeTimeEnv)
	connMaxLifeTime, err = strconv.Atoi(connMaxLifeTimeStr)
	if err != nil {
		connMaxLifeTime = databaseConnectionMaxLifeTimeDefaultSec
		logger := util.GetLogger()
		logger.Infof("%v expects int value, but %v was given.", databaseConnectionMaxLifeTimeEnv, connMaxLifeTimeStr)
		logger.Infof("use default %v [sec] for %v", databaseConnectionMaxLifeTimeDefaultSec, databaseConnectionMaxLifeTimeEnv)
	}
	engine.SetConnMaxLifetime(time.Duration(connMaxLifeTime) * time.Second)

	logLevel, err := parseLogLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		return nil, err
	}
	engine.SetLogLevel(logLevel)

	return engine, nil
}

// parseLogLevel parses level string into xorm's LogLevel
func parseLogLevel(lvl string) (core.LogLevel, error) {
	switch strings.ToLower(lvl) {
	case "panic", "fatal", "error":
		return core.LOG_ERR, nil
	case "warn", "warning":
		return core.LOG_WARNING, nil
	case "info":
		return core.LOG_INFO, nil
	case "debug":
		return core.LOG_DEBUG, nil
	}
	return core.LOG_DEBUG, fmt.Errorf("cannot parse \"%v\" into go-xorm/core.LogLevel", lvl)
}

var escapeReplace = []struct {
	Key      string
	Replaced string
}{
	{"\\", "\\\\"},
	{`'`, `\'`},
	{"\\0", "\\\\0"},
	{"\n", "\\n"},
	{"\r", "\\r"},
	{`"`, `\"`},
	{"\x1a", "\\Z"},
}

// EscapeMySQLString prevents from SQL-injection.
func EscapeMySQLString(value string) string {
	for _, r := range escapeReplace {
		value = strings.Replace(value, r.Key, r.Replaced, -1)
	}

	return value
}
