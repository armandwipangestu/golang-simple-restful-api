package repository

import (
	"github.com/armandwipangestu/golang-simple-restful-api/internal/models"
	"gorm.io/gorm"
)

type StudentRepo struct {
	DB *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *StudentRepo { 
	return &StudentRepo{DB: db}
}

func (r *StudentRepo) GetAll() ([]models.Student, error) {
	var students []models.Student
	err := r.DB.Preload("Address").Find(&students).Error

	return students, err
}

func (r *StudentRepo) GetById(id uint) (models.Student, error) {
	var s models.Student
	err := r.DB.Preload("Address").First(&s, id).Error

	return s, err
}

func (r *StudentRepo) Create(s *models.Student) (*models.Student, error) {
	err := r.DB.Create(s).Error
	if err != nil {
		return nil, err
	}

	return s, nil
}