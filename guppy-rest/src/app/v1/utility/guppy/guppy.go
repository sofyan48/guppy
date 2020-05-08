package guppy

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Luzifer/go-openssl"
	"github.com/sofyan48/guppy/guppy"
	"github.com/sofyan48/guppy/guppy-rest/src/app/v1/api/guppy/entity"
	"github.com/sofyan48/guppy/guppy-rest/src/app/v1/utility/utils"
	"github.com/sofyan48/guppy/guppy/config"
	"github.com/sofyan48/guppy/guppy/libs/etcd"
)

type GuppyLibrary struct {
	Utils utils.UtilsInterface
}

func GuppyLibraryHandler() *GuppyLibrary {
	return &GuppyLibrary{
		Utils: utils.UtilsHandler(),
	}
}

type GuppyLibraryInterface interface {
	GetClients() (*etcd.EtcdLibs, error)
	EncryptValue(value string) ([]byte, error)
	DecryptValue(value string) ([]byte, error)
	PutByPath(data *entity.RequestPayload) error
}

// GetClients ...
func (libs *GuppyLibrary) GetClients() (*etcd.EtcdLibs, error) {
	config := config.NewConfig()
	config.DialTimeOut, _ = strconv.Atoi(os.Getenv("OS_DIAL_TIMEOUT"))
	config.Urls = strings.Split(os.Getenv("OS_URLS"), ",")
	return guppy.Client(config).New()
}

// EncryptValue ...
func (libs *GuppyLibrary) EncryptValue(value string) ([]byte, error) {
	ssl := openssl.New()
	return ssl.EncryptBytes(os.Getenv("PASSPHRASE"), []byte(value))
}

// DecryptValue ...
func (libs *GuppyLibrary) DecryptValue(value string) ([]byte, error) {
	ssl := openssl.New()
	return ssl.DecryptBytes(os.Getenv("PASSPHRASE"), []byte(value))
}

// PutByPath ..
func (libs *GuppyLibrary) PutByPath(data *entity.RequestPayload) error {
	client, err := libs.GetClients()
	if err != nil {
		return nil
	}
	for _, i := range data.Parameters {
		params := client.GetParameters()
		params.Path = i.Path
		if i.IsEncrypt {
			encValue, _ := libs.EncryptValue(i.Value)
			params.Value = string(encValue)
		} else {
			params.Value = i.Value
		}
		client.Put(params)
		result, _ := client.Get(params.Path)
		log.Println("Create Revision: ", result.Kvs[0].CreateRevision)
		log.Println("Mod Revision: ", result.Kvs[0].ModRevision)
	}
	return err
}
