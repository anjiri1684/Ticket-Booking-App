package db

import (
	"github.com/anjiri1684/ticket-booking-project-v1/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	// // Drop existing table to avoid plan mismatches
	// if db.Migrator().HasTable(&models.Event{}) {
	// 	if err := db.Migrator().DropTable(&models.Event{}); err != nil {
	// 		return err
	// 	}
	// }

	// Fresh table creation
	return db.AutoMigrate(&models.Event{}, &models.Ticket{}, &models.User{})
}
