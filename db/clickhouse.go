package db

import (
	"database/sql"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"time"
)

type Clickhouse struct {
	DB *sql.DB
}

type ClickhouseConfig struct {
	Host             string `json:"host"`
	Port             int    `json:"port"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	Database         string `json:"database"`
	DialTimeout      int    `json:"dial_timeout"`
	MaxExecutionTime int    `json:"max_execution_time"`
	MaxIdleConns     int    `json:"max_idle_conns"`
	MaxOpenConns     int    `json:"max_open_conns"`
	// 连接最大生命周期（分钟）
	ConnMaxLifeTime int `json:"conn_max_life_time"`
}

func (ch *Clickhouse) Init(conf ClickhouseConfig) error {
	addr := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	if conf.Database == "" {
		conf.Database = "default"
	}
	if conf.Username == "" {
		conf.Username = "default"
	}
	if conf.DialTimeout == 0 {
		conf.DialTimeout = 30
	}
	if conf.MaxExecutionTime == 0 {
		conf.MaxExecutionTime = 60
	}
	conn := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{addr},
		Auth: clickhouse.Auth{
			Database: conf.Database,
			Username: conf.Username,
			Password: conf.Password,
		},
		//TLS: &tls.Config{
		//	InsecureSkipVerify: true,
		//},
		Settings: clickhouse.Settings{
			"max_execution_time": conf.MaxExecutionTime,
		},
		DialTimeout: time.Duration(conf.DialTimeout) * time.Second,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		Debug: true,
	})

	if conf.MaxIdleConns > 0 {
		conn.SetMaxIdleConns(conf.MaxIdleConns)
	}
	if conf.MaxOpenConns > 0 {
		conn.SetMaxOpenConns(conf.MaxOpenConns)
	}
	if conf.ConnMaxLifeTime > 0 {
		conn.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifeTime) * time.Second)
	}

	err := conn.Ping()
	if err != nil {
		return err
	}

	ch.DB = conn
	return nil
}

func (ch *Clickhouse) Close() {
	if ch.DB != nil {
		err := ch.DB.Close()
		if err != nil {
			return
		}
	}
}
