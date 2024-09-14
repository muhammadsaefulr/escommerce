// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/go-playground/validator/v10"
	"github.com/muhammadsaefulr/escommerce/internal/infras/controller"
	controller3 "github.com/muhammadsaefulr/escommerce/internal/infras/controller/cart"
	controller2 "github.com/muhammadsaefulr/escommerce/internal/infras/controller/product"
	"github.com/muhammadsaefulr/escommerce/internal/infras/repository"
	repository3 "github.com/muhammadsaefulr/escommerce/internal/infras/repository/cart"
	repository2 "github.com/muhammadsaefulr/escommerce/internal/infras/repository/product"
	"github.com/muhammadsaefulr/escommerce/internal/infras/service"
	service3 "github.com/muhammadsaefulr/escommerce/internal/infras/service/cart"
	service2 "github.com/muhammadsaefulr/escommerce/internal/infras/service/product"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitUserController(db *gorm.DB) *controller.UserController {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	validate := ProvideValidator()
	userController := controller.NewUserController(userService, validate)
	return userController
}

func InitUserSellerController(db *gorm.DB) *controller.UserSellerController {
	userSellerRepository := repository.NewUserSellerRepository(db)
	userSellerService := service.NewUserSellerService(userSellerRepository)
	validate := ProvideValidator()
	userSellerController := controller.NewUserSellerController(userSellerService, validate)
	return userSellerController
}

func InitProductController(db *gorm.DB) *controller.ProductController {
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	validate := ProvideValidator()
	productController := controller.NewProductController(productService, validate)
	return productController
}

func InitCategoryProductController(db *gorm.DB) *controller2.CategoryController {
	categoryRepository := repository2.NewCategoryRepository(db)
	categoryService := service2.NewCategoryService(categoryRepository)
	validate := ProvideValidator()
	categoryController := controller2.NewCategoryController(categoryService, validate)
	return categoryController
}

func InitShoppingCartController(db *gorm.DB) *controller3.ShoppingCartController {
	shoppingCartRepository := repository3.NewShoppingCartRepository(db)
	shoppingCartService := service3.NewShoppingCartService(shoppingCartRepository)
	validate := ProvideValidator()
	shoppingCartController := controller3.NewShoppingCartController(shoppingCartService, validate)
	return shoppingCartController
}

// wire.go:

func ProvideValidator() *validator.Validate {
	return validator.New()
}
