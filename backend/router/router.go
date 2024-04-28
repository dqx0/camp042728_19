package router

import (
	"github.com/dqx0/camp042728_19/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter(h handler.IBaseHandler) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // リクエストを許可するオリジン
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")             // クレデンシャルを許可する（Cookieなど）
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	})
	r.GET("/api/user/profile/:id", h.GetUserHandler().GetUser())
	r.POST("/api/user/profile", h.GetUserHandler().CreateUser())
	r.POST("/api/user/test", h.GetUserHandler().CreateTestUser())
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
