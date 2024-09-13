package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	entity "github.com/muhammadsaefulr/escommerce/internal/domain/entity"
	"github.com/muhammadsaefulr/escommerce/internal/modules/database/seed"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", user, password, dbname, host, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&entity.UserRole{})

	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
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

	seed.SeedRole(db)
	seed.SeedUser(db)
	seed.SeedCategories(db)
	seed.SeedProducts(db)

	return db
}
