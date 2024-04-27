package router

import (
	"github.com/dqx0/camp042728_19/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter(h handler.IBaseHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/api/user/profile/:id", h.GetUserHandler().GetUser())
	r.POST("/api/user/profile", h.GetUserHandler().CreateUser())
	r.PUT("/api/user/profile/:id", h.GetUserHandler().UpdateUser())
	r.DELETE("/api/user/profile/:id", h.GetUserHandler().DeleteUser())
	r.GET("/api/expence/:id", h.GetExpenceHandler().GetExpence())
	r.POST("/api/expence", h.GetExpenceHandler().CreateExpence())
	r.PUT("/api/expence", h.GetExpenceHandler().UpdateExpence())
	r.DELETE("/api/expence/:id", h.GetExpenceHandler().DeleteExpence())
	r.GET("/api/allowance/:id", h.GetAllowanceHandler().GetAllowance())
	r.POST("/api/allowance", h.GetAllowanceHandler().CreateAllowance())
	r.PUT("/api/allowance", h.GetAllowanceHandler().UpdateAllowance())
	r.DELETE("/api/allowance/:id", h.GetAllowanceHandler().DeleteAllowance())
	return r
}
