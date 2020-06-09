package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sofyan48/guppy/guppy-rest/src/app/v1/api/guppy/entity"
	"github.com/sofyan48/guppy/guppy-rest/src/app/v1/api/guppy/service"
	"github.com/sofyan48/guppy/guppy-rest/src/app/v1/utility/rest"
)

// GuppyController ...
type GuppyController struct {
	Service service.GuppyServiceInterface
}

// GuppyControllerHandler ...
func GuppyControllerHandler() *GuppyController {
	return &GuppyController{
		Service: service.GuppyServiceHandler(),
	}
}

// GuppyControllerInterface ...
type GuppyControllerInterface interface {
	Get(context *gin.Context)
	Path(context *gin.Context)
}

// Get ...
func (handler *GuppyController) Get(context *gin.Context) {
	params := &entity.ParametersRequest{}
	if err := context.ShouldBind(params); err != nil {
		rest.InvalidParameterResponse(context, err)
		return
	}
	data, err := handler.Service.GetService(params)
	if err != nil {
		rest.InvalidParameterResponse(context, err)
		return
	}
	rest.SuccessResponse(context, data, nil, "")
	return
}

// Path ...
func (handler *GuppyController) Path(context *gin.Context) {
	params := &entity.ParametersRequest{}
	if err := context.ShouldBind(params); err != nil {
		rest.InvalidParameterResponse(context, err)
		return
	}
	data, err := handler.Service.GetServicePath(params)
	if err != nil {
		rest.InvalidParameterResponse(context, err)
		return
	}
	rest.SuccessResponse(context, data, nil, "")
	return
}
