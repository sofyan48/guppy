package routes

import (
	"github.com/gin-gonic/gin"
)

// CONFIGROUTES ...
const CONFIGROUTES = VERSION + "/"

func (rLoader *V1RouterLoader) initConfig(router *gin.Engine) {
	group := router.Group(GUPPYROUTES)
	group.PUT("config/user", rLoader.Config.User)
	group.POST("config/key", rLoader.Config.GenerateKeys)
}
