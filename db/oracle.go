package db

import (
	"database/sql"
	"fmt"
	_ "github.com/godror/godror"
	"go.uber.org/zap"
	"strconv"
	"sync"
)

var (
	oracleLock        sync.RWMutex
	oracleInitialized bool
	oracle            *sql.DB
)

type OracleConf struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	SID      string `json:"sid"`
}

type Ora struct {
}

// OracleInit 初始化Oracle
func OracleInit(conf Config, logger *zap.Logger) {
	oracleLock.Lock()
	defer oracleLock.Unlock()

	var err error
	if oracleInitialized {
		err = fmt.Errorf("[db] oracle already initialized")
	}
	if len(conf.URL) == 0 {
		err = fmt.Errorf("[db] oracle url invalid")
	}

	oracle, err = sql.Open("godror", conf.URL)

	err = oracle.Ping()

	if err != nil {
		defer logger.Error(err.Error())
		panic(err)
	}

	logger.Info(fmt.Sprintf("[db] oracle connection successful"))
}

// FormatOracleConnUrl 格式化连接字符串
func FormatOracleConnUrl(conf *OracleConf) string {
	port, err := strconv.ParseInt(conf.Port, 10, 32)
	if err != nil || port == 0 {
		port = 1521
	}
	url := fmt.Sprintf("%s/%s@%s:%d/%s?connect_timeout=30",
		conf.User,
		conf.Password,
		conf.Host,
		port,
		conf.SID,
	)
	return url
}
