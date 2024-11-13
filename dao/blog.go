package dao

import (
	"blog_backend/utils"
	"log"
	"time"
)

type Blog struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Created     string `json:"created"`
	Content     string `json:"content"`
}
type AddBlogForm struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Created     string `json:"-"`
}

func (Blog) TableName() string {
	return "blog"
}

func (AddBlogForm) TableName() string {
	return "blog"
}

// 分页查询
func SelectAllBlogs(currentPage, pageSize int) *[]Blog {
	var blog []Blog
	offset := (currentPage - 1) * pageSize // 计算偏移量
	if err := utils.DB.Order("created desc").Offset(offset).Limit(pageSize).Find(&blog).Error; err != nil {
		log.Fatal(err)
	}
	return &blog
}

// 查询总记录数
func SelectAllBlogsCount() int64 {

	var total int64
	if err := utils.DB.Model(&Blog{}).Count(&total).Error; err != nil {
		log.Fatal("failed to count records:", err)
	}
	return total
}

// 增加更新一篇博客
func AddBlog(blog *AddBlogForm) error {
	now := time.Now()
	timeStr := now.Format("2006-01-02 15:04:05")
	if blog.Id != 0 {
		var blogResult AddBlogForm
		result := utils.DB.First(&blogResult, blog.Id)
		if result.RowsAffected > 0 {
			result = utils.DB.Updates(blog)
			return nil
		}
	}
	blog.Created = timeStr
	if err := utils.DB.Create(blog).Error; err != nil {
		return err
	}
	return nil

}

// 获取一篇博客
func SelectBlog(blogId int) *Blog {
	blog := Blog{ID: blogId}
	if err := utils.DB.Find(&blog).Error; err != nil {
		return nil
	}
	return &blog
}
