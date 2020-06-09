package routes

import (
	"github.com/gin-gonic/gin"
)

// GUPPYROUTES ...
const GUPPYROUTES = VERSION + "/"

func (rLoader *V1RouterLoader) initGuppy(router *gin.Engine) {
	group := router.Group(GUPPYROUTES)
	group.GET("get/item", rLoader.Middleware.Auth(), rLoader.Guppy.Get)
	group.GET("get/path", rLoader.Middleware.Auth(), rLoader.Guppy.Path)
	group.POST("put/raw", rLoader.Middleware.Auth(), rLoader.Guppy.PostRaw)
	group.POST("put/item", rLoader.Guppy.PostItems)
	group.DELETE("delete/item", rLoader.Middleware.Auth(), rLoader.Guppy.Delete)
}
