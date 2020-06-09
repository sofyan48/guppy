package config

import (
	"github.com/gin-gonic/gin"
	"github.com/sofyan48/guppy/guppy-rest/src/app/v1/utility/guppy"
	"github.com/sofyan48/guppy/guppy-rest/src/middleware"
)

// SetupRouter server router configuration
func SetupRouter() *gin.Engine {
	// setup default config
	defaultConfig()
	defaultMiddleware := middleware.DefaultMiddleware{}
	router := gin.Default()
	router.Use(defaultMiddleware.CORSMiddleware())
	return router
}

func defaultConfig() {
	guppyConfig := guppy.GuppyLibraryHandler()
	guppyConfig.CheckUserConfig()
}
