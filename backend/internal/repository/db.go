package repository

import (
	"teacher-assistant/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB 数据库连接实例
var DB *gorm.DB

// InitDB 初始化 MySQL 数据库连接
func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	DB = db
	return db, nil
}

// InitSQLiteDB 初始化 SQLite 数据库（用于测试）
func InitSQLiteDB(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	DB = db
	return db, nil
}

// Database 数据库操作接口
type Database interface {
	Create(value interface{}) error
	Find(dest interface{}, conds ...interface{}) error
	First(dest interface{}, conds ...interface{}) error
	Updates(value interface{}) error
	Delete(value interface{}, conds ...interface{}) error
	Where(query interface{}, args ...interface{}) Database
}

// GormDB GORM实现
type GormDB struct {
	DB *gorm.DB
}

// Create 创建记录
func (g *GormDB) Create(value interface{}) error {
	return g.DB.Create(value).Error
}

// Find 查询多条记录
func (g *GormDB) Find(dest interface{}, conds ...interface{}) error {
	return g.DB.Find(dest, conds...).Error
}

// First 查询第一条记录
func (g *GormDB) First(dest interface{}, conds ...interface{}) error {
	return g.DB.First(dest, conds...).Error
}

// Updates 更新记录
func (g *GormDB) Updates(value interface{}) error {
	return g.DB.Updates(value).Error
}

// Delete 删除记录
func (g *GormDB) Delete(value interface{}, conds ...interface{}) error {
	return g.DB.Delete(value, conds...).Error
}

// Where 添加查询条件
func (g *GormDB) Where(query interface{}, args ...interface{}) Database {
	return &GormDB{DB: g.DB.Where(query, args...)}
}

// StudentRepository 学生仓库
type StudentRepository struct {
	DB *gorm.DB
}

// NewStudentRepository 创建学生仓库
func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{DB: db}
}

// Create 创建学生
func (r *StudentRepository) Create(student *model.Student) error {
	return r.DB.Create(student).Error
}

// FindAll 查询所有学生
func (r *StudentRepository) FindAll() ([]model.Student, error) {
	var students []model.Student
	err := r.DB.Find(&students).Error
	return students, err
}

// FindByID 根据ID查询学生
func (r *StudentRepository) FindByID(id string) (*model.Student, error) {
	var student model.Student
	err := r.DB.First(&student, "id = ?", id).Error
	return &student, err
}

// Update 更新学生
func (r *StudentRepository) Update(student *model.Student) error {
	return r.DB.Save(student).Error
}

// Delete 删除学生
func (r *StudentRepository) Delete(id string) error {
	return r.DB.Delete(&model.Student{}, "id = ?", id).Error
}