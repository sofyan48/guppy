package etcd

import (
	"github.com/sofyan48/guppy/guppy/entity"
	"github.com/sofyan48/guppy/guppy/utils/connection"
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
type EtcdLibsInterface interface{}

func (lib *EtcdLibs) Put(params *entity.Parameters) {

}
