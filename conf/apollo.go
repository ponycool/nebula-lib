package conf

import (
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"sync"
)

var (
	rwLock  sync.RWMutex
	initial bool
	conf    agollo.Client
)

// Conf 配置结构体
type Conf struct {
	AppID          string `json:"app_id"`
	Cluster        string `json:"cluster"`
	IP             string `json:"ip"`
	NamespaceName  string `json:"namespace_name"`
	IsBackupConfig bool   `json:"is_backup_config"`
	Secret         string `json:"secret"`
}

// Init 初始化
func Init(c *Conf) (conf agollo.Client, err error) {
	rwLock.Lock()
	defer rwLock.Unlock()

	if initial {
		return conf, nil
	}

	Conf := &config.AppConfig{
		AppID:          c.AppID,
		Cluster:        c.Cluster,
		IP:             c.IP,
		NamespaceName:  c.NamespaceName,
		IsBackupConfig: c.IsBackupConfig,
		Secret:         c.Secret,
	}
	conf, err = agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return Conf, nil
	})
	if err != nil {
		return nil, err
	}

	initial = true
	return conf, nil
}

// Get 获取配置
func Get() agollo.Client {
	return conf
}
