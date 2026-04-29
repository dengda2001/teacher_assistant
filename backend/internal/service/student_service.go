package service

import (
	"teacher-assistant/internal/model"
	"teacher-assistant/internal/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// StudentService 学生服务
type StudentService struct {
	repo *repository.StudentRepository
}

// NewStudentService 创建学生服务
func NewStudentService(db *gorm.DB) *StudentService {
	return &StudentService{
		repo: repository.NewStudentRepository(db),
	}
}

// Create 创建学生
func (s *StudentService) Create(student *model.Student) error {
	student.ID = uuid.New().String()
	return s.repo.Create(student)
}

// List 获取学生列表
func (s *StudentService) List() ([]model.Student, error) {
	return s.repo.FindAll()
}

// Get 获取单个学生
func (s *StudentService) Get(id string) (*model.Student, error) {
	return s.repo.FindByID(id)
}

// Update 更新学生
func (s *StudentService) Update(student *model.Student) error {
	return s.repo.Update(student)
}

// Delete 删除学生
func (s *StudentService) Delete(id string) error {
	return s.repo.Delete(id)
}