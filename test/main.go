package main

import (
	"fmt"
	"github.com/ponycool/nebula-lib/conf"
	"github.com/ponycool/nebula-lib/log"
	"go.uber.org/zap"
	"os"
	"strconv"
)

func main() {

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

	fmt.Println(c.GetStringValue("database.database", "1"))

	// 获取命令行参数
	initFs()
}
