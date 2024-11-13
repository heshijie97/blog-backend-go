package service

import (
	"blog_backend/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Blog() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentPageString := c.DefaultQuery("currentPage", "1")

		currentPage, err := strconv.Atoi(currentPageString)
		if err != nil {
			c.JSON(200, gin.H{
				"code":    400,
				"message": "请输入正确的页数",
			})
		}
		pageSizeString := c.DefaultQuery("pageSize", "5")
		pageSize, err := strconv.Atoi(pageSizeString)
		if err != nil {
			c.JSON(200, gin.H{
				"code":    400,
				"message": "请输入正确的每页显示数",
			})
		}

		blogs := dao.SelectAllBlogs(currentPage, pageSize)
		fmt.Println(blogs)
		c.JSON(200, gin.H{
			"code": 200,
			"data": gin.H{
				"records":     blogs,
				"currentPage": currentPage,
				"total":       dao.SelectAllBlogsCount(),
				"pageSize":    pageSize,
			},
		})
	}
}

func AddBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		var addBlogForm dao.AddBlogForm
		// 解析 JSON 格式的请求体
		if err := c.ShouldBindJSON(&addBlogForm); err != nil {
			c.JSON(200, gin.H{
				"code":    400,
				"message": "请求体错误",
			})
			return
		}
		err := dao.AddBlog(&addBlogForm)
		if err != nil {
			c.JSON(200, gin.H{
				"code":    400,
				"message": "创建失败",
			})
			return
		}
		c.JSON(200, gin.H{
			"code":    200,
			"message": "创建成功",
		})
	}
}

func SelectBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		blogIdString := c.Param("id")
		blogId, err := strconv.Atoi(blogIdString)
		if err != nil {
			c.JSON(200, gin.H{
				"code":    400,
				"message": "id错误",
			})
			return
		}
		blog := dao.SelectBlog(blogId)
		c.JSON(200, gin.H{
			"code":    200,
			"message": "获取成功",
			"data":    blog,
		})
	}
}
