package routes

import (
	"github.com/gin-gonic/gin"
)

// GUPPYROUTES ...
const GUPPYROUTES = VERSION + "/"

func (rLoader *V1RouterLoader) initGuppy(router *gin.Engine) {
	group := router.Group(GUPPYROUTES)
	group.GET("get", rLoader.Guppy.Get)
	group.GET("path", rLoader.Guppy.Get)
}
