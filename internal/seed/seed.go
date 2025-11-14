package seed

import (
	"github.com/armandwipangestu/golang-simple-restful-api/internal/models"
	"gorm.io/gorm"
)

func MigrateAndSeed(db *gorm.DB) error {
	// Migration
	if err := db.AutoMigrate(&models.Address{}, &models.Student{}); err != nil {
		return err
	}

	// Check if data is exist or not
	var count int64
	db.Model(&models.Student{}).Count(&count)
	if count > 0 {
		return nil
	}

	// Seeder example
	addr := models.Address{
		City: "New York",
		Street: "Wall Street",
	}
	student := models.Student{
		Name: "Foo Bar",
		Age: 22,
		Address: addr,
	}

	if err := db.Create(&student).Error; err != nil {
		return err
	}

	return nil
}