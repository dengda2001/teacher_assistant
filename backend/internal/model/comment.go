package model

import (
	"time"

	"gorm.io/gorm"
)

// Comment 评语模型
type Comment struct {
	ID          string         `gorm:"primaryKey;type:varchar(36)" json:"id"`
	StudentID   string         `gorm:"type:varchar(36);not null;index" json:"student_id"`
	Semester    string         `gorm:"type:varchar(20)" json:"semester"`
	Style       string         `gorm:"type:varchar(20);default:'comprehensive'" json:"style"`
	Content     string         `gorm:"type:text;not null" json:"content"`
	GeneratedBy string         `gorm:"type:varchar(20);default:'ai'" json:"generated_by"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联
	Student Student `json:"student,omitempty" gorm:"foreignKey:StudentID"`
}

// TableName 指定表名
func (Comment) TableName() string {
	return "comments"
}