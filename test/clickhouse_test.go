package main

import (
	"github.com/ponycool/nebula-lib/db"
	"go.uber.org/zap"
	"strconv"
	"testing"
)

type perf_hw_kh_v1 struct {
	Name1   string
	Sortl   string
	Ymonths string
}

func TestClickHouse(t *testing.T) {
	logInit()
	confInit()

	ch := &db.Clickhouse{}
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

	orm := &db.ClickHouseOrm{}
	err = orm.Init(ch.DB)
	if err != nil {
		return
	}
	var kh perf_hw_kh_v1
	orm.DB.First(&kh)

	logger.Info("clickhouse test", zap.Any("kh", kh))
}
