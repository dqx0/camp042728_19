package handler

import (
	"net/http"
	"strconv"

	"github.com/dqx0/camp042728_19/model"
	"github.com/dqx0/camp042728_19/usecase"
	"github.com/gin-gonic/gin"
)

type IExpenceHandler interface {
	GetExpence() gin.HandlerFunc
	CreateExpence() gin.HandlerFunc
	UpdateExpence() gin.HandlerFunc
	DeleteExpence() gin.HandlerFunc
}
type expenceHandler struct {
	baseUsecase usecase.IBaseUsecase
}

func NewExpenceHandler(baseUsecase usecase.IBaseUsecase) IExpenceHandler {
	return &expenceHandler{baseUsecase}
}
func (h *expenceHandler) GetExpence() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		expence, err := h.baseUsecase.GetExpenceUsecase().GetByID(uint(id))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "ok", "data": expence})
	}
}
func (h *expenceHandler) CreateExpence() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var expence model.Expense
		if err := ctx.BindJSON(&expence); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		expence, err := h.baseUsecase.GetExpenceUsecase().Create(expence)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "ok", "data": expence})
	}
}
func (h *expenceHandler) UpdateExpence() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var expence model.Expense
		if err := ctx.BindJSON(&expence); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		expence, err := h.baseUsecase.GetExpenceUsecase().Update(expence)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "ok", "data": expence})
	}
}
func (h *expenceHandler) DeleteExpence() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var expence model.Expense
		if err := ctx.BindJSON(&expence); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		err := h.baseUsecase.GetExpenceUsecase().Delete(expence)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}
