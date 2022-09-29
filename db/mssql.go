package db

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

type Mssql struct {
}

// GetDSN 获取数据源名称
// https://github.com/denisenkom/go-mssqldb
// sqlserver://username:password@host:port/instance?param1=value&param2=value
func (mssql *Mssql) GetDSN(config map[string]interface{}) string {
	params := []string{"sqlserver://", config["user"].(string),
		":",
		url.QueryEscape(config["password"].(string)),
		"@",
		config["host"].(string),
		":",
		fmt.Sprintf("%d", int(config["port"].(float64))),
		"?database=", config["database"].(string),
	}

	if _, ok := config["dialectOptions"]; ok {
		//存在
		options := config["dialectOptions"].(map[string]interface{})

		for k, v := range options {
			t := reflect.TypeOf(v)

			switch t.Kind() {
			case reflect.String:
				params = append(params, fmt.Sprintf("&%s=%s", url.QueryEscape(k), url.QueryEscape(v.(string))))
			case reflect.Float64:
				params = append(params, fmt.Sprintf("&%s=%d", url.QueryEscape(k), int(v.(float64))))
			case reflect.Bool:
				params = append(params, fmt.Sprintf("&%s=%t", url.QueryEscape(k), v.(bool)))

			}

		}
	}

	dnsPath := strings.Join(params, "")
	return dnsPath
}
