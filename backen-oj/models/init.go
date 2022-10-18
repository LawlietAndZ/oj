package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)
/**
数据库初始化文件

 */

var DB = Init()

func Init() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gin_gorm_oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm init Error",err)
	}
	return db
}
