package main

import "gin-ranking/router"

func main() {
	r := router.Router()
	r.Run(":8080") // 使用gd快捷键，跳转到实现
}
