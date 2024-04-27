package handler

import (
	"net/http"
	"strconv"

	"github.com/dqx0/camp042728_19/model"
	"github.com/dqx0/camp042728_19/usecase"
	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	GetUser() gin.HandlerFunc
	CreateUser() gin.HandlerFunc
	UpdateUser() gin.HandlerFunc
	DeleteUser() gin.HandlerFunc
}
type userHandler struct {
	baseUsecase usecase.IBaseUsecase
}

func NewUserHandler(baseUsecase usecase.IBaseUsecase) IUserHandler {
	return &userHandler{baseUsecase}
}
func (h *userHandler) GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		user, err := h.baseUsecase.GetUserUsecase().GetByID(uint(id))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "ok", "data": user})
	}
}
func (h *userHandler) CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.User
		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		user, err := h.baseUsecase.GetUserUsecase().Create(user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "ok", "data": user})
	}
}
func (h *userHandler) UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.User
		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		user, err := h.baseUsecase.GetUserUsecase().Update(user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "ok", "data": user})
	}
}
func (h *userHandler) DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.User
		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if err := h.baseUsecase.GetUserUsecase().Delete(user); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}
