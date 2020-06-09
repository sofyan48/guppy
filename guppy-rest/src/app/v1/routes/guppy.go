package routes

import (
	"github.com/gin-gonic/gin"
)

// GUPPYROUTES ...
const GUPPYROUTES = VERSION + "/"

func (rLoader *V1RouterLoader) initGuppy(router *gin.Engine) {
	group := router.Group(GUPPYROUTES)
	group.GET("get/items", rLoader.Guppy.Get)
	group.GET("get/path", rLoader.Guppy.Path)
	group.POST("put/raw", rLoader.Guppy.PostRaw)
	group.POST("put/items", rLoader.Guppy.PostItems)
}
