package db

import (
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	"sync"
)

var (
	rwLock  sync.RWMutex
	initial bool
	sqlDB   *sql.DB
)

type Config struct {
	Driver            Driver `json:"driver"`
	URL               string `json:"url"`
	Enabled           bool   `json:"enabled"`
	MaxIdleConnection int    `json:"max_idle_connection"`
	MaxOpenConnection int    `json:"max_open_connection"`
	ConnMaxLifetime   int    `json:"conn_max_lifetime"`
}

type Driver int32

const (
	MariaDB Driver = 0
	MySQL   Driver = 1
	PgSQL   Driver = 2
)

func (driver Driver) String() string {
	switch driver {
	case MariaDB:
		return "mysql"
	case MySQL:
		return "mysql"
	case PgSQL:
		return "postgresql"
	default:
		return "mysql"
	}
}

// Init 初始化数据连接
func Init(conf Config, logger *zap.Logger) {
	rwLock.Lock()
	defer rwLock.Unlock()

	var err error

	if initial {
		err = fmt.Errorf("[db] db already initialized")
		logger.Error(err.Error())
		return
	}

	switch conf.Driver {
	case MariaDB:
		err = MysqlInit(conf)
	case MySQL:
		err = MysqlInit(conf)
	default:
		err = MysqlInit(conf)
	}

	if err != nil {
		defer logger.Error(err.Error())
		panic(err.Error())
	}

	logger.Info(fmt.Sprintf("[db] %s connection successful", conf.Driver.String()))
	initial = true
}

func GetDB() *sql.DB {
	return sqlDB
}
