package repository

import (
	"github.com/dqx0/camp042728_19/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetByID(id uint) (model.User, error)
	Create(user model.User) (model.User, error)
	Update(user model.User) (model.User, error)
	Delete(user model.User) error
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}
func (r *userRepository) GetByID(id uint) (model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	return user, err
}
func (r *userRepository) Create(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}
func (r *userRepository) Update(user model.User) (model.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}
func (r *userRepository) Delete(user model.User) error {
	return r.db.Delete(&user).Error
}
