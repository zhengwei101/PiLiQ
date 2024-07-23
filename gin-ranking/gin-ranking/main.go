package main

import "gin-ranking/router"

func main() {
	r := router.Router()
	err := r.Run(":8080")
	if err != nil {
		return
	}
	// 使用gd快捷键，跳转到实现
	// 编译Linux下可执行文件，GOOS=linux GOARCH=amd64 go build

}
