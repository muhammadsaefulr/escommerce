package seed

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/muhammadsaefulr/escommerce/internal/domain/entity"
	"github.com/muhammadsaefulr/escommerce/internal/modules/auth"
	"gorm.io/gorm"
)

func SeedRole(db *gorm.DB) {
	roles := []entity.UserRole{
		{RoleName: "Admin"},
		{RoleName: "User"},
		{RoleName: "Seller"},
	}

	for _, role := range roles {
		var existingRole entity.UserRole
		if err := db.Where("role_name = ?", role.RoleName).First(&existingRole).Error; err != nil && err != gorm.ErrRecordNotFound {
			log.Fatalf("failed to check existing role: %v", err)
		}

		if existingRole.ID == 0 {
			result := db.Create(&role)
			if result.Error != nil {
				log.Fatalf("failed to seed roles: %v", result.Error)
			}
		}
	}
}

func SeedUser(db *gorm.DB) {
	hashedpass, err := auth.HashBcryptPassword("testcommerce123")
	if err != nil {
		log.Fatalf(err.Error())
	}

	users := []entity.User{
		{Name: "AdminOne", Email: "admin@gmail.com", Password: hashedpass, RoleId: 1},
		{Name: "UsersOne", Email: "user@gmail.com", Password: hashedpass, RoleId: 2},
		{Name: "SellerOne", Email: "seller@gmail.com", Password: hashedpass, RoleId: 3},
	}

	for _, user := range users {
		var existingUser entity.User
		if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err != nil && err != gorm.ErrRecordNotFound {
			log.Fatalf("failed to check existing user: %v", err)
		}

		if existingUser.ID == uuid.Nil {
			result := db.Create(&user)
			if result.Error != nil {
				log.Fatalf("failed to seed users: %v", result.Error)
			}
		}
	}
}

func SeedCategories(db *gorm.DB) {
	categories := []entity.CategoryProduct{
		{Name: "Electronics", Description: "Electronic devices and gadgets"},
		{Name: "Books", Description: "Books of all genres"},
	}

	for _, category := range categories {
		var existingCategory entity.CategoryProduct
		if err := db.Where("name = ?", category.Name).First(&existingCategory).Error; err != nil && err != gorm.ErrRecordNotFound {
			log.Fatalf("failed to check existing category: %v", err)
		}

		if existingCategory.ID == 0 {
			result := db.Create(&category)
			if result.Error != nil {
				log.Fatalf("failed to seed categories: %v", result.Error)
			}
		}
	}
}

func SeedProducts(db *gorm.DB) {
	// Retrieve categories to use their IDs
	var categories []entity.CategoryProduct
	if err := db.Find(&categories).Error; err != nil {
		log.Fatalf("failed to retrieve categories: %v", err)
	}

	categoryMap := make(map[string]uint)
	for _, category := range categories {
		categoryMap[category.Name] = category.ID
	}

	products := []entity.ProductItems{
		{
			ProductName:        "Smartphone",
			ProductDescription: "Latest model smartphone with high-end specs",
			ProductPrice:       699.99,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
			CategoryId:         categoryMap["Electronics"],
		},
		{
			ProductName:        "Novel Book",
			ProductDescription: "A thrilling novel by a best-selling author",
			ProductPrice:       19.99,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
			CategoryId:         categoryMap["Books"],
		},
	}

	for _, product := range products {
		var existingProduct entity.ProductItems
		if err := db.Where("product_name = ? AND category_id = ?", product.ProductName, product.CategoryId).First(&existingProduct).Error; err != nil && err != gorm.ErrRecordNotFound {
			log.Fatalf("failed to check existing product: %v", err)
		}

		if existingProduct.ID == "" {
			result := db.Create(&product)
			if result.Error != nil {
				log.Fatalf("failed to seed products: %v", result.Error)
			}
		}
	}
}
