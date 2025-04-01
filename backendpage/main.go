package main

import (
	"backendpage/config"
	"backendpage/controllers"
	"backendpage/middleware"
	"backendpage/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// 初始化数据库
	db := config.InitDB()
	
	// 自动迁移数据库表
	db.AutoMigrate(&models.User{})

	// 创建 Gin 路由
	r := gin.Default()

	// 加载HTML模板
	r.LoadHTMLGlob("templates/*")

	// 允许跨域
	r.Use(middleware.CORSMiddleware())

	// 添加测试页面
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Welcome to Storage System",
		})
	})

	// API路由
	api := r.Group("/api")
	{
		// 用户认证相关路由
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)
		
		// 用户信息相关路由
		api.GET("/user/profile", controllers.GetUserProfile)
	}

	// 启动服务器
	r.Run(":8080")
} 