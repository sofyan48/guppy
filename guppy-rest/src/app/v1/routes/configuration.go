package routes

import (
	"github.com/gin-gonic/gin"
)

// CONFIGROUTES ...
const CONFIGROUTES = VERSION + "/"

func (rLoader *V1RouterLoader) initConfig(router *gin.Engine) {
	group := router.Group(GUPPYROUTES)
	group.PUT("config/admin", rLoader.Config.UserDefault)
	group.POST("config/admin/key", rLoader.Config.GenerateAdminKeys)

	group.POST("config/user/key", rLoader.Config.GenerateKeys)
	group.POST("config/user/add", rLoader.Config.AddUser)
	group.GET("config/user/list", rLoader.Middleware.Auth(), rLoader.Config.UserList)
	group.GET("config/user/get", rLoader.Middleware.Auth(), rLoader.Config.UserGet)
}
