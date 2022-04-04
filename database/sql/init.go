package sql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 数据库链接单例
var DB *gorm.DB

func Database(dsn string) {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// Error
	if err != nil {
		panic(err)
	}

	DB = db
	migration()
}

func migration() {
	DB.AutoMigrate(&User{})
}
