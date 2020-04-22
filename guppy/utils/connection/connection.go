package connection

import (
	"time"

	"github.com/sofyan48/guppy/guppy/entity"
	"go.etcd.io/etcd/clientv3"
)

// Connect ...
type Connect struct {
}

// ConnectHandler ...
func ConnectHandler() *Connect {
	return &Connect{}
}

// ConnectInterface ...
type ConnectInterface interface {
	Init(config *entity.Config) (*clientv3.Client, error)
}

// Init ...
func (conn *Connect) Init(config *entity.Config) (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{
		Endpoints:   config.Urls,
		DialTimeout: config.DialTimeOut * time.Second,
	})
}
