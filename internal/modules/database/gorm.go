package database

import (
	"log"

	domain "github.com/muhammadsaefulr/escommerce/internal/domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB() *gorm.DB {
	dsn := "user=saepul password=epul123 dbname=escommerce_db port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&domain.User{})
	return db
}
