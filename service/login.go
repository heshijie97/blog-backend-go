package service

import (
	"blog_backend/dao"
	"blog_backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type LoginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginForm LoginForm
		c.ShouldBind(&loginForm)
		fmt.Println(loginForm.Username, loginForm.Password)
		user := dao.SelectUser(loginForm.Username)
		if user == nil {
			c.JSON(200, gin.H{
				"code":    403,
				"message": "用户不存在",
			})
			return
		}
		if loginForm.Password == user.Password {
			tokenString, err := utils.GenerateJWT(strconv.Itoa(user.ID))
			if err != nil {
				c.JSON(200, gin.H{
					"code":    403,
					"message": "token生成失败",
				})
				return
			}
			c.Header("Authorization", tokenString)
			fmt.Println(tokenString)
			c.JSON(200, gin.H{
				"code":    200,
				"message": "ok",
				"data": gin.H{
					"id":       strconv.Itoa(user.ID),
					"username": user.Username,
				},
			})
			return
		} else {
			c.JSON(200, gin.H{
				"code":    403,
				"message": "密码错误",
			})
			return
		}
	}
}
