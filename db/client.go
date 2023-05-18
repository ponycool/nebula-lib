package db

import (
	"database/sql"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"time"
)

type Client struct {
	Clients map[string]*DB
	Configs ClientConfig
	logger  *zap.Logger
}

type ClientConfig struct {
	Clients map[string]interface{} `json:"clients"`
	Default map[string]interface{} `json:"default"`
}

type ClientDSN interface {
	GetDSN(config map[string]interface{}) string
}

// CreateClient  create the db pool for  the database
func (client *Client) CreateClient(database string) (db *DB, err error) {

	config := client.getClientConfig(database)

	var dsn ClientDSN = nil
	var driver = ""
	var dbName = ""

	if v, ok := config["driver"]; ok {
		driver = v.(string)
	}

	if v, ok := config["database"]; ok {
		dbName = v.(string)
	}

	if len(dbName) == 0 {
		return nil, errors.New("invalid database config")
	}

	switch driver {
	case "mssql", "sqlserver":
		dsn = &Mssql{}
	case "mysql":
		dsn = &Mysql{}
	default:
		client.logger.Error(fmt.Sprintf("connect to mysql database %s with invalid %s", dbName, driver))
		return nil, err
	}

	dnsPath := dsn.GetDSN(config)

	driverDB, err := sql.Open(driver, dnsPath)
	if err != nil {
		client.logger.Error(fmt.Sprintf("connect to mysql database %s error", dbName))
		return nil, err
	}

	maxIdleCons := 10
	maxLeftTime := 7200
	maxOpenCons := 50

	pool := config["pool"].(map[string]interface{})

	if v, ok := pool["maxIdleCons"]; ok {
		maxIdleCons = int(v.(float64))
	}

	if v, ok := pool["maxLeftTime"]; ok {
		maxLeftTime = int(v.(float64))
	}

	if v, ok := pool["maxOpenCons"]; ok {
		maxOpenCons = int(v.(float64))
	}

	// SetMaxIdleCons sets the maximum number of connections in the idle connection pool.
	driverDB.SetMaxIdleConns(maxIdleCons)

	// SetMaxOpenCons sets the maximum number of open connections to the database.
	driverDB.SetMaxOpenConns(maxOpenCons)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	driverDB.SetConnMaxLifetime(time.Duration(maxLeftTime) * time.Millisecond)
	if err := driverDB.Ping(); err != nil {
		return nil, err
	}

	//set db to the clients
	db = &DB{DB: driverDB}

	if _, ok := config["logging"]; ok {
		db.LogSQL = config["logging"].(bool)
	}

	// logger.Info("create mysql db %s client success", database)

	return db, err

}

// Use get the db's conn
func (client *Client) Use(database string) (db *DB) {

	c, _ := client.Clients[database]
	return c
}

// GetClientConfig 获取客户端配置
func (client *Client) getClientConfig(clientName string) (config map[string]interface{}) {

	config = make(map[string]interface{}, 10)
	for k, v := range client.Configs.Default {
		config[k] = v
	}

	clients := client.Configs.Clients

	if _, ok := clients[clientName]; !ok {
		return config
	}

	//存在
	c := clients[clientName].(map[string]interface{})

	for k, v := range c {
		config[k] = v
	}

	return config
}

// Close the database
func (client *Client) Close() error {

	for k, v := range client.Clients {
		err := v.Close()
		if err != nil {
			client.logger.Info(fmt.Sprintf("close db %s error %+v", k, err))
			return err
		}
		client.logger.Info(fmt.Sprintf("close db %s success", k))

	}
	return nil
}
