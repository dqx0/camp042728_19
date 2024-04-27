package repository

import (
	"gorm.io/gorm"
)

type IBaseRepository interface {
	Atomic(f func(IBaseRepository) error) error
	GetUserRepository() IUserRepository
	GetExpenceRepository() IExpenceRepository
	GetAllowanceRepository() IAllowanceRepository
}
type baseRepository struct {
	db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) IBaseRepository {
	return &baseRepository{db}
}
func (br *baseRepository) Atomic(fn func(IBaseRepository) error) error {
	return br.db.Transaction(func(tx *gorm.DB) error {
		return fn(NewBaseRepository(tx))
	})
}
func (br *baseRepository) GetUserRepository() IUserRepository {
	return NewUserRepository(br.db)
}
func (br *baseRepository) GetExpenceRepository() IExpenceRepository {
	return NewExpenceRepository(br.db)
}
func (br *baseRepository) GetAllowanceRepository() IAllowanceRepository {
	return NewAlliwanceRepository(br.db)
}
