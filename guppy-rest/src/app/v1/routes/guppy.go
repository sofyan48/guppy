package routes

import (
	"github.com/gin-gonic/gin"
)

// GUPPYROUTES ...
const GUPPYROUTES = VERSION + "/"

func (rLoader *V1RouterLoader) initGuppy(router *gin.Engine) {
	group := router.Group(GUPPYROUTES)
	group.GET("get/item", rLoader.Guppy.Get)
	group.GET("get/path", rLoader.Guppy.Path)
	group.POST("put/raw", rLoader.Guppy.PostRaw)
	group.POST("put/item", rLoader.Guppy.PostItems)
	group.DELETE("delete/item", rLoader.Guppy.Delete)
}
