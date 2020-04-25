package entity

import "go.etcd.io/etcd/clientv3"

// Config ...
type Client struct {
	Guppy *clientv3.Client
}
