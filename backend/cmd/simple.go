package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(corsMiddleware())

	// 测试路由
	r.GET("/api/v1/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "backend is running"})
	})

	// 学生路由
	r.GET("/api/v1/students", func(c *gin.Context) {
		c.JSON(http.StatusOK, []gin.H{
			{"id": "1", "name": "张三", "student_no": "2024001", "class": "1班", "grade": 8},
			{"id": "2", "name": "李四", "student_no": "2024002", "class": "1班", "grade": 8},
			{"id": "3", "name": "王五", "student_no": "2024003", "class": "1班", "grade": 8},
		})
	})

	// 成绩路由
	r.GET("/api/v1/scores", func(c *gin.Context) {
		c.JSON(http.StatusOK, []gin.H{
			{"id": "1", "student_name": "张三", "exam_name": "期中考试", "subject": "数学", "score": 85, "class_rank": 5},
			{"id": "2", "student_name": "李四", "exam_name": "期中考试", "subject": "数学", "score": 92, "class_rank": 2},
		})
	})

	// 表现记录路由
	r.GET("/api/v1/records", func(c *gin.Context) {
		c.JSON(http.StatusOK, []gin.H{
			{"id": "1", "student_name": "张三", "content": "课堂积极发言", "tags": ["积极"], "record_date": "2026-04-28"},
			{"id": "2", "student_name": "李四", "content": "作业完成优秀", "tags": ["优秀"], "record_date": "2026-04-27"},
		})
	})

	r.Run(":8081")
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
