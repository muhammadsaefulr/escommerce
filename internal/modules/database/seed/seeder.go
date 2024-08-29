package seed

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/muhammadsaefulr/escommerce/internal/domain/entity"
	"gorm.io/gorm"
)

func SeedCategories(db *gorm.DB) {
	categories := []entity.CategoryProduct{
		{Name: "Electronics", Description: "Electronic devices and gadgets"},
		{Name: "Books", Description: "Books of all genres"},
		{Name: "Clothing", Description: "Apparel and clothing items"},
	}

	for _, category := range categories {
		result := db.Create(&category)
		if result.Error != nil {
			log.Fatalf("failed to seed categories: %v", result.Error)
		}
	}
}

func SeedProducts(db *gorm.DB) {
	// Retrieve categories to use their IDs
	var categories []entity.CategoryProduct
	if err := db.Find(&categories).Error; err != nil {
		log.Fatalf("failed to retrieve categories: %v", err)
	}

	// Map category names to their IDs
	categoryMap := make(map[string]uint)
	for _, category := range categories {
		categoryMap[category.Name] = category.ID
	}

	products := []entity.ProductItems{
		{
			ID:                 uuid.New(),
			ProductName:        "Smartphone",
			ProductDescription: "Latest model smartphone with high-end specs",
			ProductPrice:       699.99,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
			CategoryId:         categoryMap["Electronics"],
		},
		{
			ID:                 uuid.New(),
			ProductName:        "Novel Book",
			ProductDescription: "A thrilling novel by a best-selling author",
			ProductPrice:       19.99,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
			CategoryId:         categoryMap["Books"],
		},
		{
			ID:                 uuid.New(),
			ProductName:        "T-Shirt",
			ProductDescription: "Comfortable cotton t-shirt in various sizes",
			ProductPrice:       15.99,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
			CategoryId:         categoryMap["Clothing"],
		},
	}

	for _, product := range products {
		result := db.Create(&product)
		if result.Error != nil {
			log.Fatalf("failed to seed products: %v", result.Error)
		}
	}
}
