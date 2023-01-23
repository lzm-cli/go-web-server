package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	router := r.Group("/api")
	registerUser(router)

	adminRouter := router.Group("/admin")
	registerAdminUser(adminRouter)
}
