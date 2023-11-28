package db

import (
	"github.com/jeffleon1/club_hub/franchise/pkg/entities"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(&entities.Franchise{}, &entities.Location{}, &entities.Country{})
}
