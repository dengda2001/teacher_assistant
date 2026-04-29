package service

import (
	"teacher-assistant/internal/model"

	"gorm.io/gorm"
)

// PerformanceService 表现记录服务
type PerformanceService struct {
	DB *gorm.DB
}

// NewPerformanceService 创建表现记录服务
func NewPerformanceService(db *gorm.DB) *PerformanceService {
	return &PerformanceService{DB: db}
}

// List 获取表现记录列表
func (s *PerformanceService) List(studentID string) ([]model.PerformanceRecord, error) {
	var records []model.PerformanceRecord
	query := s.DB.Preload("Student").Order("record_date DESC")

	if studentID != "" {
		query = query.Where("student_id = ?", studentID)
	}

	err := query.Find(&records).Error
	return records, err
}

// Create 创建表现记录
func (s *PerformanceService) Create(record *model.PerformanceRecord) error {
	return s.DB.Create(record).Error
}

// Update 更新表现记录
func (s *PerformanceService) Update(record *model.PerformanceRecord) error {
	return s.DB.Save(record).Error
}

// Delete 删除表现记录
func (s *PerformanceService) Delete(id string) error {
	return s.DB.Delete(&model.PerformanceRecord{}, "id = ?", id).Error
}