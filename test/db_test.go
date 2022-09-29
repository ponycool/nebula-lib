package main

import (
	"fmt"
	"github.com/ponycool/nebula-lib/db"
	"github.com/ponycool/nebula-lib/log"
	"testing"
)

func TestMySQL(t *testing.T) {

	logInit()
	confInit()

	// 初始化数据库
	var (
		dbConnUrl           string
		dbMaxIdleConnection int
		dbMaxOpenConnection int
		dbConnMaxLifetime   int
	)

	dbConnUrl = fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=True&loc=Local",
		config.GetStringValue("database.username", ""),
		config.GetValue("database.password"),
		config.GetValue("database.hostname"),
		config.GetValue("database.port"),
		config.GetValue("database.database"),
	)
	dbMaxIdleConnection = config.GetIntValue("db_max_idle_connection", 100)
	dbMaxOpenConnection = config.GetIntValue("db_max_open_connection", 100)
	dbConnMaxLifetime = config.GetIntValue("db_conn_max_lifetime", 100)

	dbConf := db.Config{
		Driver:            db.MySQL,
		URL:               dbConnUrl,
		Enabled:           true,
		MaxIdleConnection: dbMaxIdleConnection,
		MaxOpenConnection: dbMaxOpenConnection,
		ConnMaxLifetime:   dbConnMaxLifetime,
	}
	db := db.DB{}
	db.Init(dbConf, log.Get())
}
