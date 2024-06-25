package routes

import (
	"appointment-booking-app/controllers"

	"github.com/gin-gonic/gin"
)

func AppointmentRoutes(router *gin.Engine) {
	appointmentGroup := router.Group("/appointments")
	{
		appointmentGroup.POST("/", controllers.CreateAppointment)
		appointmentGroup.GET("/:id", controllers.GetAppointment)
		appointmentGroup.PUT("/:id", controllers.UpdateAppointment)
		appointmentGroup.DELETE("/:id", controllers.DeleteAppointment)
	}
}
