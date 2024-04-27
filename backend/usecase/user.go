package usecase

import (
	"github.com/dqx0/camp042728_19/model"
	"github.com/dqx0/camp042728_19/repository"
)

type IUserUsecase interface {
	GetByID(id uint) (model.User, error)
	Create(user model.User) (model.User, error)
	Update(user model.User) (model.User, error)
	Delete(user model.User) error
}
type userUsecase struct {
	repo repository.IUserRepository
}

func NewUserUsecase(repo repository.IUserRepository) IUserUsecase {
	return &userUsecase{repo}
}
func (u *userUsecase) GetByID(id uint) (model.User, error) {
	return u.repo.GetByID(id)
}
func (u *userUsecase) Create(user model.User) (model.User, error) {
	return u.repo.Create(user)
}
func (u *userUsecase) Update(user model.User) (model.User, error) {
	return u.repo.Update(user)
}
func (u *userUsecase) Delete(user model.User) error {
	return u.repo.Delete(user)
}
