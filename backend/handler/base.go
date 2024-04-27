package handler

import (
	"github.com/dqx0/camp042728_19/usecase"
)

type IBaseHandler interface {
	GetUserHandler() IUserHandler
	GetExpenceHandler() IExpenceHandler
	GetAllowanceHandler() IAllowanceHandler
}
type baseHandler struct {
	baseUsecase usecase.IBaseUsecase
}

func NewBaseHandler(usecase usecase.IBaseUsecase) IBaseHandler {
	return &baseHandler{usecase}
}

func (h *baseHandler) GetUserHandler() IUserHandler {
	return NewUserHandler(h.baseUsecase)
}

func (h *baseHandler) GetExpenceHandler() IExpenceHandler {
	return NewExpenceHandler(h.baseUsecase)
}

func (h *baseHandler) GetAllowanceHandler() IAllowanceHandler {
	return NewAllowanceHandler(h.baseUsecase)
}
