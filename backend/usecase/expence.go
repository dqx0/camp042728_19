package usecase

import (
	"time"

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
	repo   repository.IExpenceRepository
	alRepo repository.IAllowanceRepository
}

func NewExpenceUsecase(repo repository.IExpenceRepository, alRepo repository.IAllowanceRepository) IExpenceUsecase {
	return &expenceUsecase{repo, alRepo}
}
func (u *expenceUsecase) GetByID(id uint) (model.Expense, error) {
	return u.repo.GetByID(id)
}
func (u *expenceUsecase) GetLastOfUserId(id int) (model.Expense, error) {
	return u.repo.GetLastOfUserId(id)
}
func (u *expenceUsecase) GetAllByUserIdAndDate(userId uint, date string) ([]model.Expense, error) {
	return u.repo.GetAllByUJserIdAndDate(userId, date)
}
func (u *expenceUsecase) Create(expence model.Expense) (model.Expense, error) {
	expence.Date = time.Now()
	remain, err := u.GetLastOfUserId(int(expence.UserID))
	if err != nil {
		return model.Expense{}, err
	}
	if remain.Date.Day() != expence.Date.Day() {
		allowance, _ := u.alRepo.GetCurrent(int(expence.Date.Month()), int(expence.Date.Year()))
		expence.RemainingAmount = remain.RemainingAmount + allowance.Allowance/daysInMonth(int(expence.Date.Year()), int(expence.Date.Month()))
	} else {
		expence.RemainingAmount = remain.RemainingAmount
	}
	return u.repo.Create(expence)
}
func (u *expenceUsecase) Update(expence model.Expense) (model.Expense, error) {
	return u.repo.Update(expence)
}
func (u *expenceUsecase) Delete(expence model.Expense) error {
	return u.repo.Delete(expence)
}
func daysInMonth(year int, month int) int {
	t := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC)
	return t.Day()
}
