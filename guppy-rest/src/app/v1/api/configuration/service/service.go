package service

import (
	"errors"

	"github.com/sofyan48/guppy/guppy-rest/src/app/v1/api/configuration/entity"
	"github.com/sofyan48/guppy/guppy-rest/src/app/v1/utility/guppy"
)

// ConfigurationService ...
type ConfigurationService struct {
	Guppy guppy.GuppyLibraryInterface
}

// ConfigurationServiceHandler ...
func ConfigurationServiceHandler() *ConfigurationService {
	return &ConfigurationService{
		Guppy: guppy.GuppyLibraryHandler(),
	}
}

// ConfigurationServiceInterface ...
type ConfigurationServiceInterface interface {
	UserConfigurationService(params *entity.ConfigurationUserRequest) error
	GenerateUserKeys(params *entity.ConfigurationUserRequest) (*entity.ConfigurationResponse, error)
}

// UserConfigurationService ...
func (service *ConfigurationService) UserConfigurationService(params *entity.ConfigurationUserRequest) error {
	client, err := service.Guppy.GetClients()
	if err != nil {
		return err
	}
	userParams := client.GetParameters()
	userParams.Path = "app/config/" + params.User + "/username"
	userParams.Value = params.User
	client.Put(userParams)

	passParams := client.GetParameters()
	passParams.Path = "app/config/" + params.User + "/password"
	encValue, err := service.Guppy.EncryptValue(params.Password)
	if err != nil {
		return err
	}
	passParams.Value = string(encValue)
	client.Put(passParams)
	return nil
}

// GenerateUserKeys ...
func (service *ConfigurationService) GenerateUserKeys(params *entity.ConfigurationUserRequest) (*entity.ConfigurationResponse, error) {
	client, err := service.Guppy.GetClients()
	if err != nil {
		return nil, err
	}
	userCheck, err := client.Get("app/config/" + params.User + "/username")
	if err != nil {
		return nil, err
	}
	if len(userCheck.Kvs) <= 0 {
		return nil, errors.New("Admin Default Config Not Found")
	}

	if string(userCheck.Kvs[0].Value) != params.User {
		return nil, errors.New("User Not Found")
	}

	keysParams := client.GetParameters()
	keysParams.Path = "app/config/access/" + params.User + "/key"
	encValue, err := service.Guppy.EncryptValue(params.User + ":" + params.Password)
	if err != nil {
		return nil, err
	}
	keysParams.Value = string(encValue)
	client.Put(keysParams)
	response := &entity.ConfigurationResponse{}
	response.Key = string(encValue)
	response.User = params.User
	return response, nil
}
