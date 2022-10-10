package db

import (
	"go.uber.org/zap"
	mysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	orm "moul.io/zapgorm2"
	"sync"
	"time"
)

var (
	ormLock        sync.RWMutex
	ormInitialized bool
)

type Orm struct {
}

// Init 初始化ORM
func (orm *Orm) Init(tablePrefix string, logger *zap.Logger) {
	ormLock.Lock()
	defer ormLock.Unlock()

	if ormInitialized {
		return
	}

	if tablePrefix == "" {
		tablePrefix = "m_"
	}
	db = getOrm(tablePrefix, logger)
}

// 获取数据库ORM实例
func getOrm(tablePrefix string, logger *zap.Logger) *gorm.DB {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	ormLogger := orm.New(logger)
	ormLogger.SetAsDefault()
	// 忽略未找到记录错误
	ormLogger.IgnoreRecordNotFoundError = true
	dbInstance, err := gorm.Open(mysqlDriver.New(mysqlDriver.Config{
		Conn: GetDB(),
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix,
			SingularTable: true,
		},
		Logger: ormLogger,
		NowFunc: func() time.Time {
			return time.Now().In(loc)
		},
	})

	if err != nil {
		logger.Error("[orm] get db instance error", zap.Error(err))
	}
	return dbInstance
}
