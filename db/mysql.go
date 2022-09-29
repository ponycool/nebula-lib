package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
	"reflect"
	"strings"
	"time"
)

type Mysql struct {
}

func (mysql *Mysql) Init(c Config) (err error) {
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

// GetDSN 获取数据源名称
// https://github.com/go-sql-driver/mysql
// [driver[:password]@(host)][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
func (mysql *Mysql) GetDSN(config map[string]interface{}) string {
	charset := "utf8mb4"
	if _, ok := config["charset"]; ok {
		charset = config["charset"].(string)
	}

	params := []string{
		config["user"].(string),
		":",
		config["password"].(string),
		"@tcp(",
		config["host"].(string),
		":",
		fmt.Sprintf("%d", int(config["port"].(float64))),
		")/",
		config["database"].(string),
		"?charset=" + charset,
	}

	if _, ok := config["dialectOptions"]; ok {
		//存在
		options := config["dialectOptions"].(map[string]interface{})

		for k, v := range options {
			t := reflect.TypeOf(v)

			switch t.Kind() {
			case reflect.String:
				params = append(params, fmt.Sprintf("&%s=%s", k, url.QueryEscape(v.(string))))
			case reflect.Float64:
				params = append(params, fmt.Sprintf("&%s=%d", k, int(v.(float64))))
			case reflect.Bool:
				params = append(params, fmt.Sprintf("&%s=%t", k, v.(bool)))

			}

		}
	}

	dnsPath := strings.Join(params, "")
	return dnsPath
}
