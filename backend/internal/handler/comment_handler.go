package handler

import (
	"net/http"
	"teacher-assistant/internal/service"

	"github.com/gin-gonic/gin"
)

// CommentHandler 评语处理器
type CommentHandler struct {
	service *service.CommentService
}

// NewCommentHandler 创建评语处理器
func NewCommentHandler(service *service.CommentService) *CommentHandler {
	return &CommentHandler{service: service}
}

// GenerateRequest 生成请求
type GenerateRequest struct {
	StudentID string `json:"student_id" binding:"required"`
	Style     string `json:"style"`
	Length    int    `json:"length"`
}

// Generate 生成评语
func (h *CommentHandler) Generate(c *gin.Context) {
	var req GenerateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment, err := h.service.Generate(service.GenerateRequest{
		StudentID: req.StudentID,
		Style:     req.Style,
		Length:    req.Length,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comment": comment})
}

// BatchGenerateRequest 批量生成请求
type BatchGenerateRequest struct {
	StudentIDs []string `json:"student_ids" binding:"required"`
	Style      string   `json:"style"`
}

// BatchGenerate 批量生成评语
func (h *CommentHandler) BatchGenerate(c *gin.Context) {
	var req BatchGenerateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results, err := h.service.BatchGenerate(req.StudentIDs, req.Style)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}

// ExportWord 导出Word
func (h *CommentHandler) ExportWord(c *gin.Context) {
	// TODO: 实现Word导出
	c.JSON(http.StatusOK, gin.H{"message": "export word endpoint"})
}

// ExportExcel 导出Excel
func (h *CommentHandler) ExportExcel(c *gin.Context) {
	// TODO: 实现Excel导出
	c.JSON(http.StatusOK, gin.H{"message": "export excel endpoint"})
}