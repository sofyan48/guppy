package libs

import (
	"os"

	"github.com/Luzifer/go-openssl"
	"github.com/sofyan48/guppy/guppy"
	"github.com/sofyan48/guppy/guppy-cli/utils"
	"github.com/sofyan48/guppy/guppy/config"
	"github.com/sofyan48/guppy/guppy/libs/etcd"
)

type Library struct {
	Utils utils.UtilsInterface
}

func LibraryHandler() *Library {
	return &Library{
		Utils: utils.UtilsHandler(),
	}
}

type LibraryInterface interface {
	GetClients(path string) (*etcd.EtcdLibs, error)
	EncryptValue(value string) ([]byte, error)
	DecryptValue(value string) ([]byte, error)
}

// GetClients ...
func (libs *Library) GetClients(path string) (*etcd.EtcdLibs, error) {
	envi := libs.Utils.LoadEnvirontment(path)
	config := config.NewConfig()
	config.DialTimeOut = envi.DialTimeOut
	config.Urls = envi.Urls
	return guppy.Client(config).New()
}

// EncryptValue ...
func (libs *Library) EncryptValue(value string) ([]byte, error) {
	ssl := openssl.New()
	return ssl.EncryptBytes(os.Getenv("PASSPHRASE"), []byte(value))
}

// DecryptValue ...
func (libs *Library) DecryptValue(value string) ([]byte, error) {
	ssl := openssl.New()
	return ssl.DecryptBytes(os.Getenv("PASSPHRASE"), []byte(value))
}
