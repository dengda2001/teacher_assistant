package main

import (
	"log"
	"teacher-assistant/internal/handler"
	"teacher-assistant/internal/model"
	"teacher-assistant/internal/repository"
	"teacher-assistant/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化 SQLite 数据库（无需安装 MySQL，测试用）
	db, err := repository.InitSQLiteDB("teacher_assistant.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 自动迁移
	if err := model.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// 初始化服务
	studentService := service.NewStudentService(db)
	scoreService := service.NewScoreService(db)
	performanceService := service.NewPerformanceService(db)
	commentService := service.NewCommentService(db)

	// 初始化处理器
	studentHandler := handler.NewStudentHandler(studentService)
	scoreHandler := handler.NewScoreHandler(scoreService)
	performanceHandler := handler.NewPerformanceHandler(performanceService)
	commentHandler := handler.NewCommentHandler(commentService)

	// 设置路由
	r := gin.Default()
	r.Use(corsMiddleware())

	// API路由组
	api := r.Group("/api/v1")
	{
		// 学生管理
		api.GET("/students", studentHandler.List)
		api.POST("/students", studentHandler.Create)
		api.PUT("/students/:id", studentHandler.Update)
		api.DELETE("/students/:id", studentHandler.Delete)

		// 成绩管理
		api.GET("/scores", scoreHandler.List)
		api.POST("/scores", scoreHandler.Create)
		api.POST("/scores/import", scoreHandler.Import)
		api.GET("/scores/export", scoreHandler.Export)
		api.GET("/scores/analysis", scoreHandler.Analysis)

		// 表现记录
		api.GET("/records", performanceHandler.List)
		api.POST("/records", performanceHandler.Create)
		api.PUT("/records/:id", performanceHandler.Update)
		api.DELETE("/records/:id", performanceHandler.Delete)

		// 评语生成
		api.POST("/comments/generate", commentHandler.Generate)
		api.POST("/comments/batch", commentHandler.BatchGenerate)
		api.GET("/comments/export/word", commentHandler.ExportWord)
		api.GET("/comments/export/excel", commentHandler.ExportExcel)
	}

	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}