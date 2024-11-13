package dao

import (
	"blog_backend/utils"
	"fmt"
	"log"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "user"
}
func SelectAllUser() {
	var users []User
	if err := utils.DB.Find(&users).Error; err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)
}

func SelectUser(username string) *User {
	// 查询符合条件的第一条记录
	var user User
	if err := utils.DB.Where("username = ?", username).First(&user).Error; err != nil {
		fmt.Println("User not found:", err)
		return nil
	}

	// 输出查询结果
	fmt.Println("User found:", user)
	return &user

}
