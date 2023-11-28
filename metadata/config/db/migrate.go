package db

import (
	"github.com/jeffleon1/club_hub/metadata/pkg/entities"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(&entities.Metadata{}, &entities.Endpoint{})
}
