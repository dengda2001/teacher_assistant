package model

import (
	"time"

	"gorm.io/gorm"
)

// ExamScore 考试成绩模型
type ExamScore struct {
	ID         string         `gorm:"primaryKey;type:varchar(36)" json:"id"`
	StudentID  string         `gorm:"type:varchar(36);not null;index" json:"student_id"`
	ExamName   string         `gorm:"type:varchar(100);not null" json:"exam_name"`
	ExamDate   string         `gorm:"type:date;not null" json:"exam_date"`
	Subject    string         `gorm:"type:varchar(50);not null" json:"subject"`
	Score      float64        `gorm:"type:decimal(5,2)" json:"score"`
	FullScore  float64        `gorm:"type:decimal(5,2);default:100" json:"full_score"`
	ClassRank  int            `json:"class_rank"`
	GradeRank  int            `json:"grade_rank"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联
	Student Student `json:"student,omitempty" gorm:"foreignKey:StudentID"`
}

// TableName 指定表名
func (ExamScore) TableName() string {
	return "exam_scores"
}