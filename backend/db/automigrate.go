package db

import (
	"github.com/anjiri1684/ticket-booking-project-v1/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	
	return db.AutoMigrate(&models.Event{}, &models.Ticket{}, &models.User{})
}
