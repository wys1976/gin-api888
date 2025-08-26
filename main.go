package main

import (
	"log"

	"github.com/wys1976/gin-api888/config"
	"github.com/wys1976/gin-api888/database"
	"github.com/wys1976/gin-api888/routers"
)

func main() {
	// 加载配置
	cfg := config.LoadConfig()

	// 初始化数据库连接
	database.InitDB(cfg)
	defer database.CloseDB()

	// 设置路由
	router := routers.SetupRouter()

	// 启动服务器
	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
