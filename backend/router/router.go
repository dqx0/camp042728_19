package router

import (
	"github.com/dqx0/camp042728_19/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter(h handler.IBaseHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/user/:id", h.GetUserHandler().GetUser())
	r.POST("/user", h.GetUserHandler().CreateUser())
	r.PUT("/user", h.GetUserHandler().UpdateUser())
	r.DELETE("/user/:id", h.GetUserHandler().DeleteUser())
	r.GET("/expence/:id", h.GetExpenceHandler().GetExpence())
	r.POST("/expence", h.GetExpenceHandler().CreateExpence())
	r.PUT("/expence", h.GetExpenceHandler().UpdateExpence())
	r.DELETE("/expence/:id", h.GetExpenceHandler().DeleteExpence())
	r.GET("/allowance/:id", h.GetAllowanceHandler().GetAllowance())
	r.POST("/allowance", h.GetAllowanceHandler().CreateAllowance())
	r.PUT("/allowance", h.GetAllowanceHandler().UpdateAllowance())
	r.DELETE("/allowance/:id", h.GetAllowanceHandler().DeleteAllowance())
	return r
}
