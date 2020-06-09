package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sofyan48/guppy/guppy-rest/src/app/v1/api/configuration/entity"
	"github.com/sofyan48/guppy/guppy-rest/src/app/v1/api/configuration/service"
	"github.com/sofyan48/guppy/guppy-rest/src/app/v1/utility/rest"
)

// ConfigurationController types
type ConfigurationController struct {
	Service service.ConfigurationServiceInterface
}

// ConfigurationControllerHandler ...
func ConfigurationControllerHandler() *ConfigurationController {
	return &ConfigurationController{
		Service: service.ConfigurationServiceHandler(),
	}
}

// ConfigurationControllerInterface ...
type ConfigurationControllerInterface interface {
	UserDefault(context *gin.Context)
	GenerateKeys(context *gin.Context)
	AddUser(context *gin.Context)
}

// User params
// @contex: gin Context
func (ctrl *ConfigurationController) UserDefault(context *gin.Context) {
	body := &entity.ConfigurationUserRequest{}
	if err := context.ShouldBind(body); err != nil {
		rest.InvalidParameterResponse(context, err)
		return
	}

	err := ctrl.Service.UserConfigurationService(body)
	if err != nil {
		rest.ErrorResponse(context, err)
		return
	}
	rest.SuccessResponse(context, nil, nil, "User Configuration Update")
	return
}

// GenerateKeys ...
func (ctrl *ConfigurationController) GenerateKeys(context *gin.Context) {
	body := &entity.ConfigurationUserRequest{}
	if err := context.ShouldBind(body); err != nil {
		rest.InvalidParameterResponse(context, err)
		return
	}
	result, err := ctrl.Service.GenerateUserKeys(body)
	if err != nil {
		rest.ErrorResponse(context, err)
		return
	}
	rest.SuccessResponse(context, result, nil, "Keys Generate")
	return
}

// AddUser ...
func (ctrl *ConfigurationController) AddUser(context *gin.Context) {
	body := &entity.AddUserConfigRequest{}
	if err := context.ShouldBind(body); err != nil {
		rest.InvalidParameterResponse(context, err)
		return
	}
	err := ctrl.Service.AddUser(body)
	if err != nil {
		rest.ErrorResponse(context, err)
		return
	}
	rest.SuccessResponse(context, nil, nil, "User Successfully Insert")
	return
}
