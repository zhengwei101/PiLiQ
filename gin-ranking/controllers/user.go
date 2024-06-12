package controllers

import (
	"gin-ranking/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) GetUserInfo(c *gin.Context) {
	idStr := c.Param("id")
	name := c.Param("name")
	id, _ := strconv.Atoi(idStr)

	user, _ := models.GetUserTest(id)
	ReturnSuccess(c, 0, name, user, 1)
}

func (u UserController) GetList(c *gin.Context) {
	// logger.Write("日志信息", "user")
	ReturnError(c, 4004, "没有相关信息list")

	// num1 := 1
	// num2 := 0
	// num3 := num1 / num2
	// ReturnError(c, 4004, num3)
}
