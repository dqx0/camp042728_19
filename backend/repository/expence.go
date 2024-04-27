package repository

import (
	"github.com/dqx0/camp042728_19/model"
	"gorm.io/gorm"
)

type IExpenceRepository interface {
	GetByID(id uint) (model.Expense, error)
	GetAllByUJserIdAndDate(userId uint, date string) ([]model.Expense, error)
	GetLastOfUserId(id int) (model.Expense, error)
	Create(expence model.Expense) (model.Expense, error)
	Update(expence model.Expense) (model.Expense, error)
	Delete(expence model.Expense) error
}
type expenceRepository struct {
	db *gorm.DB
}

func NewExpenceRepository(db *gorm.DB) IExpenceRepository {
	return &expenceRepository{db}
}
func (r *expenceRepository) GetByID(id uint) (model.Expense, error) {
	var expence model.Expense
	err := r.db.First(&expence, id).Error
	return expence, err
}
func (r *expenceRepository) GetAllByUJserIdAndDate(userId uint, date string) ([]model.Expense, error) {
	var expences []model.Expense
	err := r.db.Where("user_id = ? AND date = ?", userId, date).Find(&expences).Error
	return expences, err
}
func (r *expenceRepository) GetLastOfUserId(id int) (model.Expense, error) {
	var expence model.Expense
	err := r.db.Last(&expence, id).Error
	return expence, err
}
func (r *expenceRepository) Create(expence model.Expense) (model.Expense, error) {
	err := r.db.Create(&expence).Error
	return expence, err
}
func (r *expenceRepository) Update(expence model.Expense) (model.Expense, error) {
	err := r.db.Save(&expence).Error
	return expence, err
}
func (r *expenceRepository) Delete(expence model.Expense) error {
	return r.db.Delete(&expence).Error
}
