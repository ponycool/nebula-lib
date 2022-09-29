package main

import (
	"github.com/apolloconfig/agollo/v4"
	"github.com/ponycool/nebula-lib/conf"
	"go.uber.org/zap"
	"os"
	"strconv"
)

var config agollo.Client

// 初始化配置
func confInit() {
	if config != nil {
		return
	}
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

	var err error
	config, err = conf.Init(opts)
	if err != nil {
		defer logger.Error("config initial failed", zap.String("error", err.Error()))
		panic("config initial failed")
	}
}
