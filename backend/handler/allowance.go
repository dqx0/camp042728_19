package handler

import (
	"net/http"
	"strconv"

	"github.com/dqx0/camp042728_19/model"
	"github.com/dqx0/camp042728_19/usecase"
	"github.com/gin-gonic/gin"
)

type IAllowanceHandler interface {
	GetAllowance() gin.HandlerFunc
	CreateAllowance() gin.HandlerFunc
	UpdateAllowance() gin.HandlerFunc
	DeleteAllowance() gin.HandlerFunc
}
type allowanceHandler struct {
	baseUsecase usecase.IBaseUsecase
}

func NewAllowanceHandler(baseUsecase usecase.IBaseUsecase) IAllowanceHandler {
	return &allowanceHandler{baseUsecase}
}
func (h *allowanceHandler) GetAllowance() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		month, _ := strconv.Atoi(ctx.Param("month"))
		year, _ := strconv.Atoi(ctx.Param("year"))
		allowance, err := h.baseUsecase.GetAllowanceUsecase().GetCurrent(month, year)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "ok", "data": allowance})
	}
}
func (h *allowanceHandler) CreateAllowance() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var allowance model.Allowance
		if err := ctx.BindJSON(&allowance); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		allowance, err := h.baseUsecase.GetAllowanceUsecase().Create(allowance)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "ok", "data": allowance})
	}
}
func (h *allowanceHandler) UpdateAllowance() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var allowance model.Allowance
		if err := ctx.BindJSON(&allowance); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		allowance, err := h.baseUsecase.GetAllowanceUsecase().Update(allowance)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "ok", "data": allowance})
	}
}
func (h *allowanceHandler) DeleteAllowance() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var allowance model.Allowance
		if err := ctx.BindJSON(&allowance); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		err := h.baseUsecase.GetAllowanceUsecase().Delete(allowance)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}
