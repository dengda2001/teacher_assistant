package handler

import (
	"net/http"
	"teacher-assistant/internal/model"
	"teacher-assistant/internal/service"

	"github.com/gin-gonic/gin"
)

// PerformanceHandler 表现记录处理器
type PerformanceHandler struct {
	service *service.PerformanceService
}

// NewPerformanceHandler 创建表现记录处理器
func NewPerformanceHandler(service *service.PerformanceService) *PerformanceHandler {
	return &PerformanceHandler{service: service}
}

// List 获取表现记录列表
func (h *PerformanceHandler) List(c *gin.Context) {
	studentID := c.Query("student_id")
	records, err := h.service.List(studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, records)
}

// Create 创建表现记录
func (h *PerformanceHandler) Create(c *gin.Context) {
	var record model.PerformanceRecord
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&record); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, record)
}

// Update 更新表现记录
func (h *PerformanceHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var record model.PerformanceRecord
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record.ID = id
	if err := h.service.Update(&record); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, record)
}

// Delete 删除表现记录
func (h *PerformanceHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}