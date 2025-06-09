package routes

import (
    "github.com/Matheus-Armando/go-api/api/handlers"
    "github.com/gin-gonic/gin"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(router *gin.Engine) {
    // Health check
    router.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    // User routes
    userHandler := handlers.NewUserHandler()
    userRoutes := router.Group("/users")
    {
        userRoutes.GET("", userHandler.GetUsers)
        userRoutes.GET("/:id", userHandler.GetUserByID)
    }
}