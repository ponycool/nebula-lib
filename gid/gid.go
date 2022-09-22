package gid

import (
	"github.com/bwmarrin/snowflake"
	"sync"
)

var (
	rwLock  sync.RWMutex
	initial bool
	node    *snowflake.Node
)

// 初始化
func initNode() error {
	rwLock.Lock()
	defer rwLock.Unlock()

	var err error

	if initial {
		return nil
	}
	node, err = snowflake.NewNode(1)
	if err != nil {
		return err
	}

	initial = true
	return nil
}

// GetGid 获取一个全局ID
func GetGid() (id int64, err error) {
	err = initNode()
	if err != nil {
		return 0, err
	}

	// Generate a snowflake ID.
	id = int64(node.Generate())

	return id, nil
}
