package job

import "go.uber.org/zap"

type DefaultLog struct {
	logger *zap.Logger
}

func (l *DefaultLog) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Info(msg, zap.Any("data", keysAndValues))
}
func (l *DefaultLog) Error(err error, msg string, keysAndValues ...interface{}) {
	l.logger.Error(msg, zap.Any("data", keysAndValues), zap.Any("error", err))
}
