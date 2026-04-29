package service

import (
	"fmt"
	"teacher-assistant/internal/model"

	"gorm.io/gorm"
)

// CommentService 评语服务
type CommentService struct {
	DB *gorm.DB
}

// NewCommentService 创建评语服务
func NewCommentService(db *gorm.DB) *CommentService {
	return &CommentService{DB: db}
}

// GenerateRequest 评语生成请求
type GenerateRequest struct {
	StudentID string `json:"student_id"`
	Style     string `json:"style"`
	Length    int    `json:"length"`
}

// Generate 生成评语
func (s *CommentService) Generate(req GenerateRequest) (string, error) {
	// 获取学生信息
	var student model.Student
	if err := s.DB.First(&student, "id = ?", req.StudentID).Error; err != nil {
		return "", err
	}

	// 获取学生成绩
	var scores []model.ExamScore
	s.DB.Where("student_id = ?", req.StudentID).Find(&scores)

	// 获取表现记录
	var records []model.PerformanceRecord
	s.DB.Where("student_id = ?", req.StudentID).Find(&records)

	// TODO: 调用AI生成评语
	// 这里先用模板生成
	comment := s.generateMockComment(student, scores, records, req.Style)
	return comment, nil
}

// BatchGenerate 批量生成评语
func (s *CommentService) BatchGenerate(studentIDs []string, style string) (map[string]string, error) {
	results := make(map[string]string)
	for _, id := range studentIDs {
		comment, err := s.Generate(GenerateRequest{
			StudentID: id,
			Style:     style,
		})
		if err != nil {
			continue
		}
		results[id] = comment
	}
	return results, nil
}

// generateMockComment 生成模拟评语
func (s *CommentService) generateMockComment(student model.Student, scores []model.ExamScore, records []model.PerformanceRecord, style string) string {
	return fmt.Sprintf("%s同学本学期表现良好，学习态度端正。希望能够继续保持，争取更大进步。", student.Name)
}