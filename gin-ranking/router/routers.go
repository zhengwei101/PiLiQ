package router

import (
	"gin-ranking/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	user := r.Group("/user")
	{
		user.GET("/info", controllers.GetUserInfo)
		user.POST("/list", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "user list")
		})
		user.PUT("/add", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "user add")
		})
		user.DELETE("/delete", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "user delete")
		})
	}

	return r
}
