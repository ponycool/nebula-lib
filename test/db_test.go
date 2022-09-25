package main

import (
	"fmt"
	"github.com/ponycool/nebula-lib/conf"
	"github.com/ponycool/nebula-lib/db"
	"github.com/ponycool/nebula-lib/log"
	"go.uber.org/zap"
	"os"
	"strconv"
	"testing"
)

func TestMySQL(t *testing.T) {
	// 初始化日志
	logger := log.Init(
		log.SetAppName("nebula-lib"),
		log.SetDevelopment(true),
		log.SetLevel(zap.DebugLevel),
		log.SetMaxSize(2),
		log.SetMaxBackups(100),
		log.SetMaxAge(30),
	)
	logger.Info("logger initial successful")

	// 初始化配置
	conf.Load()
	isBackupConfig, _ := strconv.ParseBool(os.Getenv("IS_BACKUP_CONFIG"))
	opts := &conf.Options{
		AppID:          os.Getenv("APP_ID"),
		Cluster:        os.Getenv("CLUSTER"),
		IP:             os.Getenv("IP"),
		NamespaceName:  os.Getenv("NAMESPACE_NAME"),
		IsBackupConfig: isBackupConfig,
		Secret:         os.Getenv("SECRET"),
	}

	c, err := conf.Init(opts)
	if err != nil {
		defer logger.Error("config initial failed", zap.String("error", err.Error()))
		panic("config initial failed")
	}

	// 初始化数据库
	var (
		dbConnUrl           string
		dbMaxIdleConnection int
		dbMaxOpenConnection int
		dbConnMaxLifetime   int
	)

	dbConnUrl = fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=True&loc=Local",
		c.GetStringValue("database.username", ""),
		c.GetValue("database.password"),
		c.GetValue("database.hostname"),
		c.GetValue("database.port"),
		c.GetValue("database.database"),
	)
	dbMaxIdleConnection = c.GetIntValue("db_max_idle_connection", 100)
	dbMaxOpenConnection = c.GetIntValue("db_max_open_connection", 100)
	dbConnMaxLifetime = c.GetIntValue("db_conn_max_lifetime", 100)

	dbConf := db.Config{
		Driver:            db.MySQL,
		URL:               dbConnUrl,
		Enabled:           true,
		MaxIdleConnection: dbMaxIdleConnection,
		MaxOpenConnection: dbMaxOpenConnection,
		ConnMaxLifetime:   dbConnMaxLifetime,
	}

	fmt.Println(dbConf)
	db.Init(dbConf, log.Get())
}
