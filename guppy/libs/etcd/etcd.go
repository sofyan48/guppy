package etcd

import (
	"context"
	"fmt"
	"time"

	"github.com/sofyan48/guppy/guppy/config"
	"github.com/sofyan48/guppy/guppy/entity"
	"go.etcd.io/etcd/clientv3"
)

// EtcdLibs ...
type EtcdLibs struct {
	Client      *clientv3.Client
	DialTimeout time.Duration
}

// Init ...
func Init(config *config.Configs) (*EtcdLibs, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   config.Urls,
		DialTimeout: time.Duration(config.DialTimeOut) * time.Second,
	})
	return &EtcdLibs{
		Client:      client,
		DialTimeout: time.Duration(config.DialTimeOut) * time.Second,
	}, err

}

// EtcdLibsInterface ...
type EtcdLibsInterface interface {
	Put(client *clientv3.Client, params *entity.Parameters) (*clientv3.PutResponse, error)
	Get(client *clientv3.Client, params *entity.Parameters) (*clientv3.GetResponse, error)
	Del(client *clientv3.Client, params *entity.Parameters) (*clientv3.DeleteResponse, error)
	Init(config *config.Configs) (*clientv3.Client, error)
	GetParameters() *entity.Parameters
	GetByPath(path string) (*clientv3.GetResponse, error)
}

// GetParameters ...
func (lib *EtcdLibs) GetParameters() *entity.Parameters {
	return &entity.Parameters{}
}

// Put ...
func (lib *EtcdLibs) Put(params *entity.Parameters) (*clientv3.PutResponse, error) {
	value := fmt.Sprintf("%v", params.Value)
	ctx, cancel := context.WithTimeout(context.Background(), lib.DialTimeout)
	result, err := lib.Client.Put(ctx, params.Path, value)
	cancel()
	return result, err
}

// Del ...
func (lib *EtcdLibs) Del(params *entity.Parameters) (*clientv3.DeleteResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), lib.DialTimeout)
	result, err := lib.Client.Delete(ctx, params.Path)
	cancel()
	return result, err
}

// Get ...
func (lib *EtcdLibs) Get(path string) (*clientv3.GetResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), lib.DialTimeout)
	result, err := lib.Client.Get(ctx, path)
	cancel()
	return result, err
}

// GetByPath ...
func (lib *EtcdLibs) GetByPath(path string) (*clientv3.GetResponse, error) {
	opts := []clientv3.OpOption{
		clientv3.WithPrefix(),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend),
		clientv3.WithLimit(3),
	}

	ctx, cancel := context.WithTimeout(context.Background(), lib.DialTimeout)
	result, err := lib.Client.Get(ctx, path, opts...)
	cancel()
	return result, err
}
