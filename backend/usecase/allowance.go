package usecase

import (
	"github.com/dqx0/camp042728_19/model"
	"github.com/dqx0/camp042728_19/repository"
)

type IAllowanceUsecase interface {
	GetCurrent(month int, year int) (model.Allowance, error)
	Create(allowance model.Allowance) (model.Allowance, error)
	Update(allowance model.Allowance) (model.Allowance, error)
	Delete(allowance model.Allowance) error
}
type allowanceUsecase struct {
	repo repository.IAllowanceRepository
}

func NewAllowanceUsecase(repo repository.IAllowanceRepository) IAllowanceUsecase {
	return &allowanceUsecase{repo}
}
func (u *allowanceUsecase) GetCurrent(month int, year int) (model.Allowance, error) {
	return u.repo.GetCurrent(month, year)
}
func (u *allowanceUsecase) Create(allowance model.Allowance) (model.Allowance, error) {
	return u.repo.Create(allowance)
}
func (u *allowanceUsecase) Update(allowance model.Allowance) (model.Allowance, error) {
	return u.repo.Update(allowance)
}
func (u *allowanceUsecase) Delete(allowance model.Allowance) error {
	return u.repo.Delete(allowance)
}
