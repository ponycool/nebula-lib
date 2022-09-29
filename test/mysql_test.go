package main

import (
	"fmt"
	"github.com/ponycool/nebula-lib/db"
	"github.com/ponycool/nebula-lib/log"
	"testing"
)

func TestTableCount(t *testing.T) {
	t.Helper()
	logInit()
	confInit()

	orm := new(db.Orm)
	orm.Init("", log.Get())

	table := "m_account"
	count := db.TableCount(table)
	fmt.Printf(fmt.Sprintf("表%s存在%d条数据", table, count))
}
