// Package database 提供数据库连接和操作的相关功能。
// 该包包含数据库初始化、连接测试和资源清理等方法。
package database

import (
	"database/sql" // 提供通用的 SQL 数据库接口
	"gin-api888/config"

	// 提供格式化 I/O 功能
	"log" // 提供简单的日志记录功能

	_ "github.com/go-sql-driver/mysql" // 匿名导入MySQL驱动，仅执行其初始化函数
)

// DB 是一个全局变量，指向已初始化的数据库连接池。
// 使用 sql.DB 类型表示数据库句柄，它是并发安全的，可供多个goroutine同时使用。
var DB *sql.DB

// InitDB 初始化数据库连接。
// 此函数使用硬编码的连接参数（根据你的要求设置）。
func InitDB(*config.Config) {
	// 构建数据库连接字符串（DSN, Data Source Name）
	// 格式: "用户名:密码@tcp(主机地址:端口)/数据库名?charset=utf8mb4&parseTime=True&loc=Local"
	// 使用你的具体参数：用户名 wys, 密码 123456, 数据库名 user_manager
	// 假设MySQL服务器在本地localhost默认端口3306
	dsn := "wys:123456@tcp(localhost:3306)/user_manager?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	// 使用 sql.Open 初始化一个数据库连接
	// 第一个参数指定驱动类型为 "mysql"
	// 注意: sql.Open() 并不立即建立连接，只是创建一个连接对象以备使用
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		// 如果初始化失败，记录错误并终止程序
		log.Fatal("Failed to initialize database connection:", err)
	}

	// 测试数据库连接是否真正可用
	// DB.Ping() 方法会实际尝试与数据库建立连接，验证连接参数的正确性
	err = DB.Ping()
	if err != nil {
		// 如果 Ping 失败，说明连接参数有误或数据库不可达，记录错误并终止程序
		log.Fatal("Failed to ping database. Please check your credentials and ensure MySQL is running:", err)
	}

	// 配置连接池参数（可选，但推荐用于生产环境）
	DB.SetMaxOpenConns(25)   // 设置与数据库的最大打开连接数
	DB.SetMaxIdleConns(10)   // 设置连接池中的最大空闲连接数
	DB.SetConnMaxLifetime(0) // 设置连接可复用的最大时间（0表示永久）

	// 连接成功，记录信息日志
	log.Println("Successfully connected to the MySQL database!")
}

// CloseDB 关闭数据库连接，释放资源。
// 通常在程序退出前调用此函数，确保数据库连接被正确关闭，避免资源泄漏。
func CloseDB() {
	if DB != nil {
		// 调用 DB.Close() 关闭连接池，释放所有打开的连接
		// 注意: 关闭后不能再使用 DB 对象执行任何数据库操作
		err := DB.Close()
		if err != nil {
			log.Printf("Warning: Error closing database connection: %v\n", err)
		} else {
			log.Println("Database connection closed gracefully.")
		}
	}
}
