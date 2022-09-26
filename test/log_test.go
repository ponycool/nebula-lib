package main

import (
	"github.com/ponycool/nebula-lib/log"
	"go.uber.org/zap"
)

var logger *zap.Logger

// 初始化日志
func logInit() {
	logger = log.Init(
		log.SetAppName("nebula-lib"),
		log.SetDevelopment(true),
		log.SetLevel(zap.DebugLevel),
		log.SetMaxSize(2),
		log.SetMaxBackups(100),
		log.SetMaxAge(30),
	)
	logger.Info("logger initial successful")
}
