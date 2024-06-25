package models

import (
	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	UserID      uint   `json:"user_id"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Time        string `json:"time"`
}
