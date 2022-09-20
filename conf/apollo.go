package conf

import (
	"fmt"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/agcache"
	"github.com/apolloconfig/agollo/v4/env/config"
	"sync"
)

var (
	rwLock  sync.RWMutex
	initial bool
	conf    agollo.Client
)

// Options  配置结构体
type Options struct {
	AppID          string `json:"app_id"`
	Cluster        string `json:"cluster"`
	IP             string `json:"ip"`
	NamespaceName  string `json:"namespace_name"`
	IsBackupConfig bool   `json:"is_backup_config"`
	Secret         string `json:"secret"`
}

// Init 初始化
func Init(opts *Options) (conf agollo.Client, err error) {
	rwLock.Lock()
	defer rwLock.Unlock()

	if initial {
		return conf, nil
	}

	c := &config.AppConfig{
		AppID:          opts.AppID,
		Cluster:        opts.Cluster,
		IP:             opts.IP,
		NamespaceName:  opts.NamespaceName,
		IsBackupConfig: opts.IsBackupConfig,
		Secret:         opts.Secret,
	}

	conf, err = agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	if err != nil {
		return nil, err
	}

	initial = true
	return conf, nil
}

// Cache 配置缓存
func Cache(namespace string, client agollo.Client) agcache.CacheInterface {
	cache := client.GetConfigCache(namespace)
	return cache
}

// CheckKey 检查配置键
func CheckKey(namespace string, client agollo.Client) {
	cache := client.GetConfigCache(namespace)
	count := 0
	cache.Range(func(key, value interface{}) bool {
		fmt.Println("key : ", key, ", value :", value)
		count++
		return true
	})
	if count < 1 {
		panic("config key can not be null")
	}
}
