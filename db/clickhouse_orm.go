package db

import (
	"database/sql"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ClickHouseOrm struct {
	DB *gorm.DB
}

func (ch *ClickHouseOrm) Init(conn *sql.DB) error {
	db, err := gorm.Open(
		clickhouse.New(clickhouse.Config{
			Conn: conn,
		}),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		},
	)
	ch.DB = db
	if err != nil {
		return err
	}
	return nil
}
