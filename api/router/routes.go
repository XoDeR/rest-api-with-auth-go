package router

import (
	"api-auth/api/controllers"

	"github.com/gin-gonic/gin"
)

func GetRoute(r *gin.Engine) {
	// public routes
	r.POST("/api/signup", controllers.Signup)
	r.POST("/api/login", controllers.Login)
}
