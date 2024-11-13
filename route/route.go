package route

import (
	"blog_backend/service"
	"github.com/gin-gonic/gin"
)

func SetRoute(r *gin.Engine) {
	r.POST("/login", service.Login())
	r.GET("/blog", service.Blog())
	r.GET("/blog/:id", service.SelectBlog())
	r.POST("/blog/edit", service.AddBlog())
}
