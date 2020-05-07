package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
}

// Get ...
func (handler *GuppyController) Get(context *gin.Context) {
	rest.ResponseMessages(context, http.StatusOK, "OK")
}
