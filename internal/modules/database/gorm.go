package database

import (
	"log"

	entity "github.com/muhammadsaefulr/escommerce/internal/domain/entity"
	"github.com/muhammadsaefulr/escommerce/internal/modules/database/seed"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB() *gorm.DB {
	dsn := "user=saepul password=epul123 dbname=escommerce_db port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&entity.User{})

	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	db.AutoMigrate(&entity.CategoryProduct{})

	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	db.AutoMigrate(&entity.ProductItems{})

	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	seed.SeedCategories(db)
	seed.SeedProducts(db)

	return db
}
