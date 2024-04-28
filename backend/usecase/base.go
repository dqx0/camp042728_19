package usecase

import (
	"github.com/dqx0/camp042728_19/repository"
)

type IBaseUsecase interface {
	GetUserUsecase() IUserUsecase
	GetExpenceUsecase() IExpenceUsecase
	GetAllowanceUsecase() IAllowanceUsecase
}
type baseUsecase struct {
	baseRepo repository.IBaseRepository
}

func NewBaseUsecase(repo repository.IBaseRepository) IBaseUsecase {
	return &baseUsecase{repo}
}
func (u *baseUsecase) GetUserUsecase() IUserUsecase {
	return NewUserUsecase(u.baseRepo.GetUserRepository())
}
func (u *baseUsecase) GetExpenceUsecase() IExpenceUsecase {
	return NewExpenceUsecase(u.baseRepo.GetExpenceRepository(), u.baseRepo.GetAllowanceRepository())
}
func (u *baseUsecase) GetAllowanceUsecase() IAllowanceUsecase {
	return NewAllowanceUsecase(u.baseRepo.GetAllowanceRepository())
}
