package model

import (
	"time"

	"gorm.io/gorm"
)

// Student 学生模型
type Student struct {
	ID        string         `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Name      string         `gorm:"type:varchar(50);not null" json:"name"`
	StudentNo string         `gorm:"type:varchar(20);unique" json:"student_no"`
	Class     string         `gorm:"type:varchar(20)" json:"class"`
	Grade     int            `gorm:"type:tinyint" json:"grade"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联
	Scores      []ExamScore          `json:"scores,omitempty" gorm:"foreignKey:StudentID"`
	Records     []PerformanceRecord  `json:"records,omitempty" gorm:"foreignKey:StudentID"`
	Comments    []Comment            `json:"comments,omitempty" gorm:"foreignKey:StudentID"`
}

// TableName 指定表名
func (Student) TableName() string {
	return "students"
}