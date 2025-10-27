package routes

import (
	"github.com/gin-gonic/gin"

	"codechef/backend/controllers"
	"codechef/backend/middleware"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"ok": true}) })

	auth := r.Group("/api/auth")
	auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)

	events := r.Group("/api/events")
	events.GET("/", controllers.ListEvents)
	events.GET(":id", middleware.AuthRequired(), controllers.GetEvent)
	events.POST("/", middleware.AuthRequired(), middleware.RoleAllowed("creator"), controllers.CreateEvent)
	events.PATCH(":id", middleware.AuthRequired(), middleware.RoleAllowed("creator"), controllers.UpdateEvent)
	events.DELETE(":id", middleware.AuthRequired(), middleware.RoleAllowed("creator"), controllers.DeleteEvent)
}
