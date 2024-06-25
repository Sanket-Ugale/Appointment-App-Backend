package services

import (
	"appointment-booking-app/config"
	"appointment-booking-app/models"
)

func CreateAppointment(appointment *models.Appointment) error {
	if err := config.DB.Create(appointment).Error; err != nil {
		return err
	}
	return nil
}

func GetAppointmentById(id string) (*models.Appointment, error) {
	var appointment models.Appointment
	if err := config.DB.First(&appointment, id).Error; err != nil {
		return nil, err
	}
	return &appointment, nil
}

func UpdateAppointment(id string, updatedAppointment *models.Appointment) error {
	var appointment models.Appointment
	if err := config.DB.First(&appointment, id).Error; err != nil {
		return err
	}
	appointment.Description = updatedAppointment.Description
	appointment.Date = updatedAppointment.Date
	appointment.Time = updatedAppointment.Time
	if err := config.DB.Save(&appointment).Error; err != nil {
		return err
	}
	return nil
}

func DeleteAppointment(id string) error {
	if err := config.DB.Delete(&models.Appointment{}, id).Error; err != nil {
		return err
	}
	return nil
}
