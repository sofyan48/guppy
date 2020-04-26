package utils

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/sofyan48/guppy/guppy"
	"github.com/sofyan48/guppy/guppy-cli/entity"
	"github.com/sofyan48/guppy/guppy/config"
	"github.com/sofyan48/guppy/guppy/libs/etcd"
)

// Utils ...
type Utils struct{}

// UtilsHandler ...
func UtilsHandler() *Utils {
	return &Utils{}
}

// UtilsInterface ..
type UtilsInterface interface {
	GetClients(urls []string) (*etcd.EtcdLibs, error)
	ReadEnvironment(path string) *entity.Environment
}

// GetClients ...
func (util *Utils) GetClients(envi *entity.Environment) (*etcd.EtcdLibs, error) {
	config := config.NewConfig()
	config.DialTimeOut = envi.DialTimeOut
	config.Urls = envi.Urls
	return guppy.Client(config).New()
}

// ReadEnvironment ...
func (util *Utils) ReadEnvironment(path string) *entity.Environment {
	envi := &entity.Environment{}
	if path != "" {
		log.Println("ENV BY PATH")
	}
	urls := strings.Split(os.Getenv("URLS"), ",")
	envi.DialTimeOut, _ = strconv.Atoi(os.Getenv("DIAL_TIMEOUT"))
	envi.Urls = urls
	return envi
}
