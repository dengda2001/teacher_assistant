package service

import (
	"teacher-assistant/internal/model"

	"gorm.io/gorm"
)

// ScoreService 成绩服务
type ScoreService struct {
	DB *gorm.DB
}

// NewScoreService 创建成绩服务
func NewScoreService(db *gorm.DB) *ScoreService {
	return &ScoreService{DB: db}
}

// List 获取成绩列表
func (s *ScoreService) List(examName, subject string) ([]model.ExamScore, error) {
	var scores []model.ExamScore
	query := s.DB.Preload("Student")

	if examName != "" {
		query = query.Where("exam_name = ?", examName)
	}
	if subject != "" {
		query = query.Where("subject = ?", subject)
	}

	err := query.Find(&scores).Error
	return scores, err
}

// Create 创建成绩记录
func (s *ScoreService) Create(score *model.ExamScore) error {
	return s.DB.Create(score).Error
}

// Import 导入成绩
func (s *ScoreService) Import(scores []model.ExamScore) error {
	return s.DB.CreateInBatches(scores, 100).Error
}

// Analysis 成绩分析
func (s *ScoreService) Analysis(examName string) (map[string]interface{}, error) {
	// 这里可以实现各种分析逻辑
	analysis := map[string]interface{}{
		"average": 82.5,
		"max": 98.0,
		"min": 45.0,
		"pass_rate": 0.92,
	}
	return analysis, nil
}