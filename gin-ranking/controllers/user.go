package controllers

import "github.com/gin-gonic/gin"

func GetUserInfo(c *gin.Context) {
	ReturnSuccess(c, 0, "success", "user info", 1)
}
