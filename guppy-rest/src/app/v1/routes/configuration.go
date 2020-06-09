package routes

import (
	"github.com/gin-gonic/gin"
)

// CONFIGROUTES ...
const CONFIGROUTES = VERSION + "/"

func (rLoader *V1RouterLoader) initConfig(router *gin.Engine) {
	group := router.Group(GUPPYROUTES)
	group.PUT("config/user", rLoader.Config.UserDefault)
	group.POST("config/key", rLoader.Config.GenerateKeys)
	group.POST("config/user/add", rLoader.Config.AddUser)
}
