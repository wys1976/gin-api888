package main

import (
	"gin-api888/internal/pkg/database"
	"gin-api888/routes"
	"log"
)

func main() {
	// 1. 初始化数据库连接
	db, err := database.InitMySQL()
	if err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}

	// 2. 初始化 Gin 路由
	r := routes.SetupRouter(db)

	// 3. 启动服务器
	log.Println("Server starting on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
