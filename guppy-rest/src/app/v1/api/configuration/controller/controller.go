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
	GenerateAdminKeys(context *gin.Context)
	AddUser(context *gin.Context)
	UserList(context *gin.Context)
	UserGet(context *gin.Context)
}

// UserDefault params
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

// GenerateAdminKeys ...
func (ctrl *ConfigurationController) GenerateAdminKeys(context *gin.Context) {
	body := &entity.ConfigurationUserRequest{}
	if err := context.ShouldBind(body); err != nil {
		rest.InvalidParameterResponse(context, err)
		return
	}
	result, err := ctrl.Service.GenerateAdminKeys(body)
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

// UserList ...
func (ctrl *ConfigurationController) UserList(context *gin.Context) {
	params := &entity.MetaListParams{}
	if err := context.ShouldBind(params); err != nil {
		rest.InvalidParameterResponse(context, err)
		return
	}
	result, err := ctrl.Service.ListUser(params)
	if err != nil {
		rest.ErrorResponse(context, err)
		return
	}
	rest.SuccessResponse(context, result, params, "")
}

// UserGet ...
func (ctrl *ConfigurationController) UserGet(context *gin.Context) {
	params := &entity.MetaUserRequest{}
	if err := context.ShouldBind(params); err != nil {
		rest.InvalidParameterResponse(context, err)
		return
	}
	result, err := ctrl.Service.GetUser(params)
	if err != nil {
		rest.ErrorResponse(context, err)
		return
	}
	rest.SuccessResponse(context, result, params, "")
}
