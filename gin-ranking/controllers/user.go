package controllers

import "github.com/gin-gonic/gin"

func GetUserInfo(c *gin.Context) {
	ReturnSuccess(c, 0, "success", "user info", 1)
}

func GetList(c *gin.Context) {
	ReturnError(c, 4004, "没有相关信息")
}
