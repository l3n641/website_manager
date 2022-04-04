package conf

import (
	"github.com/joho/godotenv"
	"os"
	"website_manager/database/sql"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()
	database := os.Getenv("database")
	sql.Database(database)
}
