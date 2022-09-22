package log

import (
	"go.uber.org/zap"
	"sync"
)

var (
	rwLock  sync.RWMutex
	initial bool
	logger  *zap.Logger
)

// Init 初始化日志
func Init(opts ...ModOptions) *zap.Logger {
	rwLock.Lock()
	defer rwLock.Unlock()

	if initial {
		return logger
	}

	logger = NewLogger(opts...)

	initial = true
	return logger
}

// Get 获取初始化后Logger实例
func Get() *zap.Logger {
	if logger == nil {
		panic("logger is not initialized")
	}
	return logger
}
