package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ponycool/nebula-lib/uuid"
	"go.uber.org/zap"
	"sync"
)

var (
	rwLock      sync.RWMutex
	initialized bool
	sqlDB       *sql.DB
)

type Config struct {
	Driver            Driver `json:"driver"`
	URL               string `json:"url"`
	Enabled           bool   `json:"enabled"`
	MaxIdleConnection int    `json:"max_idle_connection"`
	MaxOpenConnection int    `json:"max_open_connection"`
	ConnMaxLifetime   int    `json:"conn_max_lifetime"`
}

type Driver int32

const (
	MariaDB Driver = 0
	MySQL   Driver = 1
	PgSQL   Driver = 2
)

func (driver Driver) String() string {
	switch driver {
	case MariaDB:
		return "mysql"
	case MySQL:
		return "mysql"
	case PgSQL:
		return "postgresql"
	default:
		return "mysql"
	}
}

// Init 初始化数据连接
func (db *DB) Init(conf Config, logger *zap.Logger) {
	rwLock.Lock()
	defer rwLock.Unlock()

	var err error

	if initialized {
		err = fmt.Errorf("[db] db already initialized")
		logger.Error(err.Error())
		return
	}

	switch conf.Driver {
	case MySQL, MariaDB:
		mysql := Mysql{}
		err = mysql.Init(conf)
	default:
		mysql := Mysql{}
		err = mysql.Init(conf)
	}

	if err != nil {
		defer logger.Error(err.Error())
		panic(err.Error())
	}

	logger.Info(fmt.Sprintf("[db] %s connection successful", conf.Driver.String()))
	initialized = true
}

func GetDB() *sql.DB {
	return sqlDB
}

type DB struct {
	*sql.DB
	LogSQL bool
	logger *zap.Logger
}

type RowsResult struct {
	*sql.Rows
	LastError error
}

type RowResult struct {
	rows      *sql.Rows
	LastError error
}

// RawDB 返回*sql.DB
func (db *DB) RawDB() *sql.DB {
	return db.DB
}

// Query 执行一个返回 RowsResult 的查询，通常是一个 SELECT
// args 用于查询中的任何占位符参数
func (db *DB) Query(query string, args ...interface{}) *RowsResult {
	if db.LogSQL {
		db.logger.Info(fmt.Sprintf("query sql:%s", SQL{}.Format(query, args...)))
	}
	rs, err := db.DB.Query(query, args...)
	return &RowsResult{rs, err}
}

// QueryContext 执行一个返回 RowsResult 的查询，通常是一个 SELECT
// args 用于查询中的任何占位符参数
func (db *DB) QueryContext(ctx context.Context, query string, args ...interface{}) *RowsResult {
	if db.LogSQL {
		db.logger.Info(fmt.Sprintf("QueryContext sql:%s", SQL{}.Format(query, args...)))
	}
	rs, err := db.DB.QueryContext(ctx, query, args...)
	return &RowsResult{rs, err}
}

// QueryRowContext 执行一个预期最多返回一行的查询
// QueryRowContext 总是返回一个非零值。 错误被推迟到调用行的 Scan 方法。
// 如果查询没有选择任何行，*Row 的 Scan 将返回 ErrNoRows
// 否则，*Row 的 Scan 扫描第一个选定的行并丢弃其余的部分
func (db *DB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *RowResult {
	if db.LogSQL {
		db.logger.Info(fmt.Sprintf("QueryRowContext sql:%s", SQL{}.Format(query, args...)))
	}
	rows, err := db.DB.QueryContext(ctx, query, args...)
	return &RowResult{rows: rows, LastError: err}
}

// QueryRow 执行一个预期最多返回一行的查询
// QueryRow 总是返回一个非零值。 错误被推迟到调用行的 Scan 方法
// 否则，*Row 的 Scan 扫描第一个选定的行并丢弃其余的部分
func (db *DB) QueryRow(query string, args ...interface{}) *RowResult {
	if db.LogSQL {
		db.logger.Info(fmt.Sprintf("QueryRow sql:%s", SQL{}.Format(query, args...)))
	}
	return db.QueryRowContext(context.Background(), query, args...)
}

// Exec 执行查询而不返回任何行
// args 用于查询中的任何占位符参数
func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	if db.LogSQL {
		db.logger.Info(fmt.Sprintf("Exec sql:%s", SQL{}.Format(query, args...)))
	}

	return db.DB.Exec(query, args...)
}

// ExecContext 执行查询而不返回任何行
// args 用于查询中的任何占位符参数
func (db *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	if db.LogSQL {
		db.logger.Info(fmt.Sprintf("ExecContext sql:%s", SQL{}.Format(query, args...)))
	}

	return db.DB.ExecContext(ctx, query, args...)
}

// Close 将连接返回到连接池
func (r *RowsResult) Close() error {
	return r.Rows.Close()
}

// Scan 扫描
func (r *RowsResult) Scan(dest interface{}) error {

	if r.LastError != nil {
		return r.LastError
	}

	if r.Err() != nil {
		return r.Err()
	}
	err := ScanResult(r.Rows, dest)
	return err
}

// Raw 返回一行数据
func (r *RowsResult) Raw() (*sql.Rows, error) {
	return r.Rows, r.LastError
}

// Err 返回结果的最后错误
func (r *RowResult) Err() error {
	return r.LastError
}

// Scan RowResult's scan
func (r *RowResult) Scan(dest interface{}) error {

	if r.LastError != nil {
		return r.LastError
	}

	if r.Err() != nil {
		return r.Err()
	}

	if r.rows.Err() != nil {
		return r.rows.Err()
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(r.rows)

	if !r.rows.Next() {
		if err := r.rows.Err(); err != nil {
			return err
		}
		return sql.ErrNoRows
	}

	err := Scan(r.rows, dest)

	if err != nil {
		return err
	}

	// Make sure the query can be processed to completion with no errors.
	return r.rows.Close()
}

// Open  init all the database clients
func Open(configs ClientConfig, log *zap.Logger) (*Client, error) {
	dialect := &Client{}
	dialect.Clients = make(map[string]*DB)
	dialect.Configs = configs
	dialect.logger = log

	for k := range configs.Clients {

		db, err := dialect.CreateClient(k)
		if err != nil {
			return nil, err
		}
		db.logger = dialect.logger
		dialect.Clients[k] = db
	}

	return dialect, nil
}

// Begin starts a transaction. The default isolation level is dependent on
// the driver.
func (db *DB) Begin() (*Trans, error) {
	return db.BeginTx(context.Background(), nil)
}

// BeginTx starts a transaction.
//
// The provided context is used until the transaction is committed or rolled back.
// If the context is canceled, the sql package will roll back
// the transaction. Tx.Commit will return an error if the context provided to
// BeginTx is canceled.
//
// The provided TxOptions is optional and may be nil if defaults should be used.
// If a non-default isolation level is used that the driver doesn't support,
// an error will be returned.
func (db *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Trans, error) {

	transactionID := ""
	if db.LogSQL {

		transactionID = uuid.GetUuid()
		if db.LogSQL {
			db.logger.Info(fmt.Sprintf("Executing (%s): START TRANSACTION;", transactionID))
		}
	}

	rawTx, err := db.DB.BeginTx(ctx, opts)
	return &Trans{Tx: rawTx, TransactionID: transactionID, DB: db}, err
}
