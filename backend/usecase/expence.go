package usecase

import (
	"github.com/dqx0/camp042728_19/model"
	"github.com/dqx0/camp042728_19/repository"
)

type IExpenceUsecase interface {
	GetByID(id uint) (model.Expense, error)
	GetAllByUserIdAndDate(userId uint, date string) ([]model.Expense, error)
	Create(expence model.Expense) (model.Expense, error)
	Update(expence model.Expense) (model.Expense, error)
	Delete(expence model.Expense) error
}
type expenceUsecase struct {
	repo repository.IExpenceRepository
}

func NewExpenceUsecase(repo repository.IExpenceRepository) IExpenceUsecase {
	return &expenceUsecase{repo}
}
func (u *expenceUsecase) GetByID(id uint) (model.Expense, error) {
	return u.repo.GetByID(id)
}
func (u *expenceUsecase) GetAllByUserIdAndDate(userId uint, date string) ([]model.Expense, error) {
	return u.repo.GetAllByUJserIdAndDate(userId, date)
}
func (u *expenceUsecase) Create(expence model.Expense) (model.Expense, error) {
	return u.repo.Create(expence)
}
func (u *expenceUsecase) Update(expence model.Expense) (model.Expense, error) {
	return u.repo.Update(expence)
}
func (u *expenceUsecase) Delete(expence model.Expense) error {
	return u.repo.Delete(expence)
}
