package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sofyan48/guppy/guppy-rest/src/app/v1/utility/guppy"
)

// GetToken params
// @context: *gin.Context
// return gin.HandlerFunc
func (m *DefaultMiddleware) GetToken(context *gin.Context) string {
	token := context.Request.Header["Authorization"]
	if len(token) < 1 {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Keys Not Found",
		})
		context.Abort()
	}
	return token[0]
}

// Auth ...
func (m *DefaultMiddleware) Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		keys := m.GetToken(ctx)
		clientHandler := guppy.GuppyLibraryHandler()
		client, err := clientHandler.GetClients()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			ctx.Abort()
			return
		}
		keysData, err := clientHandler.DecryptValue(keys)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Keys Not Found",
			})
			ctx.Abort()
			return
		}
		keysDataExtract := strings.Split(string(keysData), ":")
		configApp, err := client.GetByPath("app/config/access/" + keysDataExtract[0] + "/key")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			ctx.Abort()
			return
		}
		if len(configApp.Kvs) <= 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Keys Not Found",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
