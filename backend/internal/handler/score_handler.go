package handler

import (
	"net/http"
	"teacher-assistant/internal/model"
	"teacher-assistant/internal/service"

	"github.com/gin-gonic/gin"
)

// ScoreHandler 成绩处理器
type ScoreHandler struct {
	service *service.ScoreService
}

// NewScoreHandler 创建成绩处理器
func NewScoreHandler(service *service.ScoreService) *ScoreHandler {
	return &ScoreHandler{service: service}
}

// List 获取成绩列表
func (h *ScoreHandler) List(c *gin.Context) {
	examName := c.Query("exam_name")
	subject := c.Query("subject")

	scores, err := h.service.List(examName, subject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, scores)
}

// Create 创建成绩
func (h *ScoreHandler) Create(c *gin.Context) {
	var score model.ExamScore
	if err := c.ShouldBindJSON(&score); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&score); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, score)
}

// Import 导入成绩
func (h *ScoreHandler) Import(c *gin.Context) {
	// TODO: 实现Excel导入
	c.JSON(http.StatusOK, gin.H{"message": "import endpoint"})
}

// Export 导出成绩
func (h *ScoreHandler) Export(c *gin.Context) {
	// TODO: 实现Excel导出
	c.JSON(http.StatusOK, gin.H{"message": "export endpoint"})
}

// Analysis 成绩分析
func (h *ScoreHandler) Analysis(c *gin.Context) {
	examName := c.Query("exam_name")
	analysis, err := h.service.Analysis(examName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, analysis)
}