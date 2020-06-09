package service

import (
	"github.com/sofyan48/guppy/guppy-rest/src/app/v1/api/guppy/entity"
	"github.com/sofyan48/guppy/guppy-rest/src/app/v1/utility/guppy"
)

// GuppyService ...
type GuppyService struct {
	Guppy guppy.GuppyLibraryInterface
}

// GuppyServiceHandler ...
func GuppyServiceHandler() *GuppyService {
	return &GuppyService{
		Guppy: guppy.GuppyLibraryHandler(),
	}
}

// GuppyServiceInterface ...
type GuppyServiceInterface interface {
	GetService(params *entity.ParametersRequest) (*entity.GetResponse, error)
	GetServicePath(params *entity.ParametersRequest) ([]entity.GetResponse, error)
}

// GetService ...
func (service *GuppyService) GetService(params *entity.ParametersRequest) (*entity.GetResponse, error) {
	client, err := service.Guppy.GetClients()
	if err != nil {
		return nil, err
	}
	data, err := client.Get(params.Path)

	if err != nil {
		return nil, err
	}
	if len(data.Kvs) == 0 {
		return nil, nil
	}

	var value string
	if params.IsEncrypt {
		decValue, err := service.Guppy.DecryptValue(string(data.Kvs[0].Value))
		if err != nil {
			return nil, err
		}
		value = string(decValue)
	} else {
		value = string(data.Kvs[0].Value)
	}

	result := &entity.GetResponse{}
	result.Revision = data.Kvs[0].CreateRevision - data.Kvs[0].ModRevision
	result.Path = string(data.Kvs[0].Key)
	result.Value = value
	result.Version = data.Kvs[0].Version
	return result, nil
}

// GetServicePath ...
func (service *GuppyService) GetServicePath(params *entity.ParametersRequest) ([]entity.GetResponse, error) {
	client, err := service.Guppy.GetClients()
	if err != nil {
		return nil, err
	}
	data, err := client.GetByPath(params.Path)
	if err != nil {
		return nil, err
	}
	var result []entity.GetResponse
	for _, i := range data.Kvs {
		data := entity.GetResponse{}
		data.Revision = i.CreateRevision - i.ModRevision
		data.Path = string(i.Key)
		data.Value = string(i.Value)
		data.Version = i.Version
		result = append(result, data)
	}
	return result, nil
}
