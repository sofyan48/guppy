package service

import (
	"errors"
	"strings"

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
	GenerateAdminKeys(params *entity.ConfigurationUserRequest) (*entity.ConfigurationResponse, error)
	AddUser(body *entity.AddUserConfigRequest) error
	ListUser(params *entity.MetaListParams) ([]entity.UserDataConfig, error)
	GetUser(params *entity.MetaUserRequest) (interface{}, error)
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
	userCheck, err := client.Get("app/config/user/" + params.User + "/username")
	if err != nil {
		return nil, err
	}
	if len(userCheck.Kvs) <= 0 {
		return nil, errors.New("User Not Registere")
	}

	if string(userCheck.Kvs[0].Value) != params.User {
		return nil, errors.New("User Not Registere")
	}

	keysParams := client.GetParameters()
	keysParams.Path = "app/config/access/" + params.User + "/key"
	encValue, err := service.Guppy.EncryptValue(params.User + ":" + keysParams.Path)
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

// GenerateAdminKeys ...
func (service *ConfigurationService) GenerateAdminKeys(params *entity.ConfigurationUserRequest) (*entity.ConfigurationResponse, error) {
	client, err := service.Guppy.GetClients()
	if err != nil {
		return nil, err
	}
	userCheck, err := client.Get("app/config/" + params.User + "/username")
	if err != nil {
		return nil, err
	}
	if len(userCheck.Kvs) <= 0 {
		return nil, errors.New("User Not Registere")
	}

	if string(userCheck.Kvs[0].Value) != params.User {
		return nil, errors.New("User Not Registere")
	}

	keysParams := client.GetParameters()
	keysParams.Path = "app/config/access/" + params.User + "/key"
	encValue, err := service.Guppy.EncryptValue(params.User + ":" + keysParams.Path)
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

// AddUser ...
func (service *ConfigurationService) AddUser(body *entity.AddUserConfigRequest) error {
	client, err := service.Guppy.GetClients()
	if err != nil {
		return err
	}
	checkUser, err := client.GetByPath("app/config/user/" + body.User)
	if err != nil {
		return err
	}
	if len(checkUser.Kvs) >= 1 {
		return errors.New("Username Found")
	}

	userParamsList := client.GetParameters()
	userParamsList.Path = "app/config/user/list/" + body.User
	userParamsList.Value = "app/config/user/" + body.User
	_, err = client.Put(userParamsList)
	if err != nil {
		return err
	}

	userParams := client.GetParameters()
	userParams.Path = "app/config/user/" + body.User + "/username"
	userParams.Value = body.User
	_, err = client.Put(userParams)
	if err != nil {
		return err
	}
	userPassParams := client.GetParameters()
	userPassParams.Path = "app/config/user/" + body.User + "/password"
	encValue, err := service.Guppy.EncryptValue(body.Password)
	if err != nil {
		return err
	}
	userPassParams.Value = string(encValue)
	_, err = client.Put(userPassParams)
	if err != nil {
		return err
	}
	userRolesParams := client.GetParameters()
	userRolesParams.Path = "app/config/user/" + body.User + "/roles"
	userRolesParams.Value = body.Roles
	_, err = client.Put(userRolesParams)
	if err != nil {
		return err
	}
	return nil
}

// ListUser ...
func (service *ConfigurationService) ListUser(params *entity.MetaListParams) ([]entity.UserDataConfig, error) {
	client, err := service.Guppy.GetClients()
	if err != nil {
		return nil, err
	}
	userList, err := client.GetByPath("app/config/user/list")
	if err != nil {
		return nil, err
	}
	response := []entity.UserDataConfig{}
	for _, i := range userList.Kvs {
		data := entity.UserDataConfig{}
		name := strings.Split(string(i.Value), "/")
		data.User = name[len(name)-1]
		roles, _ := client.Get(string(i.Value) + "/roles")
		data.Roles = string(roles.Kvs[0].Value)
		data.Path = string(i.Value)
		response = append(response, data)
	}
	return response, nil
}

// GetUser ...
func (service *ConfigurationService) GetUser(params *entity.MetaUserRequest) (interface{}, error) {
	client, err := service.Guppy.GetClients()
	if err != nil {
		return nil, err
	}
	userPath, err := client.GetByPath("app/config/user/" + params.User + "/")
	if err != nil {
		return nil, err
	}
	var data []interface{}
	for _, i := range userPath.Kvs {
		userParams, _ := client.Get(string(i.Key))
		indexName := strings.Split(string(userParams.Kvs[0].Key), "/")
		index := indexName[len(indexName)-1]
		dataParams := map[string]string{
			"key":   index,
			"path":  string(userParams.Kvs[0].Key),
			"value": string(userParams.Kvs[0].Value),
		}

		data = append(data, dataParams)

	}

	keyData, _ := client.Get("app/config/access/" + params.User + "/key")
	keyParams := map[string]string{
		"key":   string(keyData.Kvs[0].Key),
		"value": string(keyData.Kvs[0].Value),
	}
	dataUser := map[string]interface{}{
		"authorization": keyParams,
		"details":       data,
	}

	return dataUser, nil
}
