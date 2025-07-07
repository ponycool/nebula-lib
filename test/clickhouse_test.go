package main

import (
	"github.com/ponycool/nebula-lib/db"
	"go.uber.org/zap"
	"strconv"
	"testing"
)

type demo struct {
	ID int
}

func TestClickHouse(t *testing.T) {
	logInit()
	confInit()

	ch := new(db.Clickhouse)
	port, _ := strconv.Atoi(config.GetValue("clickhouse.port"))
	err := ch.Init(db.ClickhouseConfig{
		Host:     config.GetValue("clickhouse.host"),
		Port:     port,
		Username: config.GetValue("clickhouse.username"),
		Database: config.GetValue("clickhouse.database"),
	})
	if err != nil {
		return
	}

	orm := new(db.ClickHouseOrm)
	err = orm.Init(ch.DB)
	if err != nil {
		return
	}
	var demos []demo
	orm.DB.Find(&demos)

	logger.Info("clickhouse test", zap.Any("demos", demos))
}
