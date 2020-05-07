package guppy

import (
	"log"
	"os"

	"github.com/Luzifer/go-openssl"
	"github.com/sofyan48/guppy/guppy"
	"github.com/sofyan48/guppy/guppy-cli/entity"
	"github.com/sofyan48/guppy/guppy-cli/utils"
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
	GetClients(path string) (*etcd.EtcdLibs, error)
	EncryptValue(value string) ([]byte, error)
	DecryptValue(value string) ([]byte, error)
	PutByPath(EnvPath string, data *entity.TemplatesModels) error
}

// GetClients ...
func (libs *GuppyLibrary) GetClients(path string) (*etcd.EtcdLibs, error) {
	envi := libs.Utils.LoadEnvirontment(path)
	config := config.NewConfig()
	config.DialTimeOut = envi.DialTimeOut
	config.Urls = envi.Urls
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
func (libs *GuppyLibrary) PutByPath(EnvPath string, data *entity.TemplatesModels) error {
	client, err := libs.GetClients(EnvPath)
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
