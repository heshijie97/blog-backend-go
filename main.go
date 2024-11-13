package main

import (
	"blog_backend/middleware"
	"blog_backend/route"
	"blog_backend/utils"
	"github.com/gin-gonic/gin"
)

var addr string = ":8081"

func main() {
	// 初始化数据库
	utils.MysqlInit()

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	route.SetRoute(r)
	r.Run(addr) // 监听并在 0.0.0.0:8080 上启动服务
}
