package model

import (
	"time"

	"gorm.io/gorm"
)

// PerformanceRecord 表现记录模型
type PerformanceRecord struct {
	ID          string         `gorm:"primaryKey;type:varchar(36)" json:"id"`
	StudentID   string         `gorm:"type:varchar(36);not null;index" json:"student_id"`
	Content     string         `gorm:"type:text;not null" json:"content"`
	Tags        string         `gorm:"type:json" json:"tags"`
	RecordDate  string         `gorm:"type:date;not null" json:"record_date"`
	Source      string         `gorm:"type:varchar(20);default:'text'" json:"source"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联
	Student Student `json:"student,omitempty" gorm:"foreignKey:StudentID"`
}

// TableName 指定表名
func (PerformanceRecord) TableName() string {
	return "performance_records"
}