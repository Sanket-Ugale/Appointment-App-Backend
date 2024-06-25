package main

import (
	"appointment-app/middleware"
	"appointment-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Middleware
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.AuthMiddleware())

	// Routes
	routes.UserRoutes(router)
	routes.AppointmentRoutes(router)
	routes.AdminRoutes(router)

	router.Run(":8080")
}
