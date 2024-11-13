package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func MysqlInit() {
	// 连接字符串格式：username:password@tcp(host:port)/dbname
	dsn := "root:Magedu0!@tcp(192.168.111.10:3306)/blog"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("数据库连接失败:", err)
		panic(err)
	}

	fmt.Println("成功连接到数据库")
	DB = db
}
