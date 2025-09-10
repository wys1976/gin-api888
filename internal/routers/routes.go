package routes

import (
	"database/sql"
	"gin-api888/internal/handler"
	"gin-api888/internal/repository"
	"gin-api888/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(db *sql.DB, r *gin.Engine) {
	// 初始化各层依赖
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// 添加Swagger文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API路由组
	api := r.Group("/api")
	{
		// 用户相关路由
		users := api.Group("/users")
		{
			// @Summary      获取所有用户
			// @Description  获取用户列表
			// @Tags         users
			// @Produce      json
			// @Success      200  {array}   models.User
			// @Router       /users [get]
			users.GET("", userHandler.GetAllUsers)

			// @Summary      获取单个用户
			// @Description  通过ID获取用户详情
			// @Tags         users
			// @Produce      json
			// @Param        id   path      int  true  "用户ID"
			// @Success      200  {object}  models.User
			// @Failure      404  {object}  map[string]interface{}
			// @Router       /users/{id} [get]
			users.GET("/:id", userHandler.GetUserByID)

			// @Summary      创建用户
			// @Description  创建新用户
			// @Tags         users
			// @Accept       json
			// @Produce      json
			// @Param        user  body      models.User  true  "用户信息"
			// @Success      201  {object}  models.User
			// @Router       /users [post]
			users.POST("", userHandler.CreateUser)

			// @Summary      更新用户
			// @Tags         users
			// @Accept       json
			// @Produce      json
			// @Param        id    path      int          true  "用户ID"
			// @Param        user  body      models.User  true  "更新信息"
			// @Success      200  {object}  models.User
			// @Router       /users/{id} [put]
			users.PUT("/:id", userHandler.UpdateUser)

			// @Summary      删除用户
			// @Tags         users
			// @Produce      json
			// @Param        id   path      int  true  "用户ID"
			// @Success      204
			// @Router       /users/{id} [delete]
			users.DELETE("/:id", userHandler.DeleteUser)
		}
	}

	// 健康检查路由
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "✅ 服务运行正常")
	})
}
