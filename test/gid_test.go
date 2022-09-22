package main

import (
	"github.com/ponycool/nebula-lib/gid"
	"testing"
)

// 测试获取GID
func TestGetGid(t *testing.T) {
	if ans, _ := gid.GetGid(); ans == 0 {
		t.Errorf("git expected not 0 , but %d got", ans)
	}
}
