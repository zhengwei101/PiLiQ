package controllers

import "github.com/gin-gonic/gin"

type OrderController struct{}

func (o OrderController) GetList(c *gin.Context) {
	ReturnError(c, 4004, "没有相关信息order")
}
