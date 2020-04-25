package etcd

import (
	"fmt"

	"github.com/sofyan48/guppy/guppy/entity"
	"github.com/sofyan48/guppy/guppy/utils/connection"
	"go.etcd.io/etcd/clientv3"
)

// EtcdLibs ...
type EtcdLibs struct {
	Conn connection.ConnectInterface
}

// EtcdLibsHandler ...
func EtcdLibsHandler() *EtcdLibs {
	return &EtcdLibs{
		Conn: connection.ConnectHandler(),
	}
}

// EtcdLibsInterface ...
type EtcdLibsInterface interface {
	Put(client *clientv3.Client, params *entity.Parameters) (*clientv3.PutResponse, error)
}

// Put ...
func (lib *EtcdLibs) Put(client *clientv3.Client, params *entity.Parameters) (*clientv3.PutResponse, error) {
	value := fmt.Sprintf("%v", params.Value)
	return client.Put(client.Ctx(), params.Path, value)
}

// Del ...
func (lib *EtcdLibs) Del(client *clientv3.Client, params *entity.Parameters) (*clientv3.DeleteResponse, error) {
	return client.Delete(client.Ctx(), params.Path)
}
