package db

import (
	"database/sql"
	"rentCar/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func Init() *gorm.DB {
    sqlDB, _ := sql.Open(
        "pgx", "host=localhost port=5432 user=selvianl dbname=rentcar password=selvianl123 sslmode=disable",
    )   
    db, err := gorm.Open(postgres.New(postgres.Config{
        Conn: sqlDB,
    }), &gorm.Config{})
    if err != nil {
        panic("Failed")
    }
    db.AutoMigrate(
		&models.Customer{}, &models.Location{}, &models.Vendor{},
		&models.Office{},
	)
    return db
}