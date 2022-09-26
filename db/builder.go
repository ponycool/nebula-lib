package db

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
)

var (
	ormLock        sync.RWMutex
	ormInitialized bool
	db             *gorm.DB
)

// OrmInit 初始化ORM
func OrmInit(tablePrefix string, logger *zap.Logger) {
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

// Raw 将原生SQL扫描至模型
func Raw(dest interface{}, psql string, param ...interface{}) {
	db.Raw(psql, param...).Scan(dest)
}

// Exec 执行原生SQL，并返回受影响行数
func Exec(psql string, param ...interface{}) (RowsAffected int64, err error) {
	result := db.Exec(psql, param...)
	if result.Error != nil {
		return 0, err
	}
	return result.RowsAffected, nil
}

// FindOne  获取一条数据
func FindOne(dest interface{}, cond ...interface{}) {
	db.First(dest, cond...)
}

// FindOneByCond 根据条件获取一条数据
func FindOneByCond(dest interface{}, cond interface{}) {
	db.Where(cond).First(dest)
}

// FindOneWhere 根据预处理语句获取一条数据 psql(Prepared SQL Statement)预处理语句
func FindOneWhere(dest interface{}, psql string, param ...interface{}) {
	db.Where(psql, param...).First(dest)
}

// FindWhere 根据预处理语句获取数据 psql(Prepared SQL Statement)预处理语句
func FindWhere(dest interface{}, psql string, param ...interface{}) {
	db.Where(psql, param...).Find(dest)
}

// TableCount 统计指定表的数据记录总行数
func TableCount(table string) (count int64) {
	db.Table(table).Count(&count)
	return count
}

// CountWhere 根据条件统计指定模型的数据记录总行数
func CountWhere(model interface{}, psql string, param ...interface{}) (count int64) {
	db.Model(model).Where(psql, param...).Count(&count)
	return count
}

// Save 保存一条数据
func Save(model interface{}) (RowsAffected int64, err error) {
	result := db.Create(model)
	if result.Error != nil {
		return 0, err
	}
	return result.RowsAffected, nil
}

// SaveByFields 保存一条数据，并制定字段
func SaveByFields(model interface{}, fields ...string) (RowsAffected int64, err error) {
	result := db.Select(fields).Create(model)
	if result.Error != nil {
		return 0, err
	}
	return result.RowsAffected, nil
}

// CreateInBatches 分批批量插入
func CreateInBatches(model interface{}, count int) (RowsAffected int64, err error) {
	result := db.CreateInBatches(model, count)
	if result.Error != nil {
		return 0, err
	}
	return result.RowsAffected, nil
}

// Updates 更新多列
func Updates(model interface{}, data map[string]interface{}) (RowsAffected int64, err error) {
	result := db.Model(model).Updates(data)
	if result.Error != nil {
		return 0, err
	}
	return result.RowsAffected, nil
}
