package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func MysqlInit(c Config) (err error) {
	c.Enabled = true
	if !c.Enabled {
		err = fmt.Errorf("[db] mysql disabled")
		return err
	}

	if len(c.URL) == 0 {
		err = fmt.Errorf("[db] mysql url invalid")
		return err
	}

	// 创建连接
	sqlDB, err = sql.Open(MySQL.String(), c.URL)
	if err != nil {
		return err
	}

	// 设置最大连接数
	if c.MaxOpenConnection == 0 {
		c.MaxOpenConnection = 100
	}
	sqlDB.SetMaxOpenConns(c.MaxOpenConnection)

	// 设置最大闲置数
	if c.MaxIdleConnection == 0 {
		c.MaxIdleConnection = 100
	}
	sqlDB.SetMaxIdleConns(c.MaxIdleConnection)

	// 连接数据库连接超时时间
	if c.ConnMaxLifetime == 0 {
		c.ConnMaxLifetime = 100
	}
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(c.ConnMaxLifetime))

	// 激活连接
	if err = sqlDB.Ping(); err != nil {
		return err
	}

	return nil
}
