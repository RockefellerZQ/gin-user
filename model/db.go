package model

import (
	"fmt"
	"gin-user/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	if db == nil {
		args := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
			config.Conf.Database.Username,
			config.Conf.Database.Password,
			config.Conf.Database.Host,
			config.Conf.Database.Dbname,
		)
		tempDb, err := gorm.Open(mysql.Open(args), &gorm.Config{})
		if err != nil {
			panic("连接数据库失败：" + err.Error())
		}
		db = tempDb
	}
	return db
}
