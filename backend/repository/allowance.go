package repository

import (
	"github.com/dqx0/camp042728_19/model"
	"gorm.io/gorm"
)

type IAllowanceRepository interface {
	GetCurrent(month int, year int) (model.Allowance, error)
	Create(allowance model.Allowance) (model.Allowance, error)
	Update(allowance model.Allowance) (model.Allowance, error)
	Delete(allowance model.Allowance) error
}
type allowanceRepository struct {
	db *gorm.DB
}

func NewAlliwanceRepository(db *gorm.DB) IAllowanceRepository {
	return &allowanceRepository{db}
}
func (r *allowanceRepository) GetCurrent(month int, year int) (model.Allowance, error) {
	var allowance model.Allowance
	err := r.db.Where("month = ? AND year = ?", month, year).First(&allowance).Error
	return allowance, err
}
func (r *allowanceRepository) Create(allowance model.Allowance) (model.Allowance, error) {
	err := r.db.Create(&allowance).Error
	return allowance, err
}
func (r *allowanceRepository) Update(allowance model.Allowance) (model.Allowance, error) {
	err := r.db.Save(&allowance).Error
	return allowance, err
}
func (r *allowanceRepository) Delete(allowance model.Allowance) error {
	return r.db.Delete(&allowance).Error
}
