package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabaseConnection() error {
	pgConfig := postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres password=password dbname=example_db port=12345 sslmode=disable TimeZone=Europe/Bucharest",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	})
	var err error
	if DB, err = gorm.Open(pgConfig, &gorm.Config{}); err != nil {
		return err
	}
	return nil
}
