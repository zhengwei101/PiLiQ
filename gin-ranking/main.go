package main

import "gin-ranking/router"

func main() {
	//r := gin.Default()

	// r.GET("/hello", func(ctx *gin.Context) {
	// 	ctx.String(http.StatusOK, "Hello World")
	// })

	// r.POST("/user/list", func(ctx *gin.Context) {
	// 	ctx.String(http.StatusOK, "user list")
	// })
	// r.PUT("/user/add", func(ctx *gin.Context) {
	// 	ctx.String(http.StatusOK, "user add")
	// })
	// r.DELETE("/user/delete", func(ctx *gin.Context) {
	// 	ctx.String(http.StatusOK, "user delete")
	// })

	// user := r.Group("/user")
	// {
	// 	user.POST("/list", func(ctx *gin.Context) {
	// 		ctx.String(http.StatusOK, "user list")
	// 	})
	// 	user.PUT("/add", func(ctx *gin.Context) {
	// 		ctx.String(http.StatusOK, "user add")
	// 	})
	// 	user.DELETE("/delete", func(ctx *gin.Context) {
	// 		ctx.String(http.StatusOK, "user delete")
	// 	})
	// }
	r := router.Router()
	r.Run(":9999")
}
