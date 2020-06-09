package routes

import (
	"github.com/gin-gonic/gin"
	config "github.com/sofyan48/guppy/guppy-rest/src/app/v1/api/configuration/controller"
	guppy "github.com/sofyan48/guppy/guppy-rest/src/app/v1/api/guppy/controller"
	health "github.com/sofyan48/guppy/guppy-rest/src/app/v1/api/health/controller"
	"github.com/sofyan48/guppy/guppy-rest/src/middleware"
)

// VERSION ...
const VERSION = "v1"

// V1RouterLoader types
type V1RouterLoader struct {
	Middleware middleware.DefaultMiddleware
	Health     health.HealthControllerInterface
	Guppy      guppy.GuppyControllerInterface
	Config     config.ConfigurationControllerInterface
}

// V1RouterLoaderHandler ...
func V1RouterLoaderHandler() *V1RouterLoader {
	return &V1RouterLoader{
		Health: health.HealthControllerHandler(),
		Guppy:  guppy.GuppyControllerHandler(),
		Config: config.ConfigurationControllerHandler(),
	}
}

// V1RouterLoaderInterface ...
type V1RouterLoaderInterface interface {
	V1Routes(router *gin.Engine)
}

// V1Routes Params
// @router: gin.Engine
func (rLoader *V1RouterLoader) V1Routes(router *gin.Engine) {
	rLoader.initDocs(router)
	rLoader.initHealth(router)
	rLoader.initGuppy(router)
	rLoader.initConfig(router)
}
