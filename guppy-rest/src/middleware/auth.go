package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sofyan48/guppy/guppy-rest/src/app/v1/utility/guppy"
)

// Auth ...
func (m *DefaultMiddleware) Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientHandler := guppy.GuppyLibraryHandler()
		client, err := clientHandler.GetClients()
		if err != nil {
			ctx.AbortWithError(500, err)
		}
		configApp, err := client.GetByPath("app/config/")
		if err != nil {
			ctx.AbortWithError(500, err)
		}
		if len(configApp.Kvs) <= 0 {
			ctx.AbortWithStatus(401)
		}

		ctx.Next()
	}
}
