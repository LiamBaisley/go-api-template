package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Open() error {
	var err error
	connectionString := "host=localhost user=api password=123456 dbname=api port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return err
	}

	migrationErr := DB.AutoMigrate(&User{})
	if migrationErr != nil {
		return migrationErr
	}
	return nil
}
