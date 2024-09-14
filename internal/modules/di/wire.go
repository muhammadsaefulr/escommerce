package di

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/muhammadsaefulr/escommerce/internal/infras/controller"
	CartController "github.com/muhammadsaefulr/escommerce/internal/infras/controller/cart"
	CpController "github.com/muhammadsaefulr/escommerce/internal/infras/controller/product"
	"github.com/muhammadsaefulr/escommerce/internal/infras/repository"
	CartRepository "github.com/muhammadsaefulr/escommerce/internal/infras/repository/cart"
	CpRepository "github.com/muhammadsaefulr/escommerce/internal/infras/repository/product"
	"github.com/muhammadsaefulr/escommerce/internal/infras/service"
	CartService "github.com/muhammadsaefulr/escommerce/internal/infras/service/cart"
	CpService "github.com/muhammadsaefulr/escommerce/internal/infras/service/product"

	"gorm.io/gorm"
)

func ProvideValidator() *validator.Validate {
	return validator.New()
}

func InitUserController(db *gorm.DB) *controller.UserController {
	wire.Build(
		repository.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
		ProvideValidator,
	)
	return &controller.UserController{}
}

func InitUserSellerController(db *gorm.DB) *controller.UserSellerController {
	wire.Build(
		repository.NewUserSellerRepository,
		service.NewUserSellerService,
		controller.NewUserSellerController,
		ProvideValidator,
	)
	return &controller.UserSellerController{}
}

func InitProductController(db *gorm.DB) *controller.ProductController {
	wire.Build(
		repository.NewProductRepository,
		service.NewProductService,
		controller.NewProductController,
		ProvideValidator,
	)

	return &controller.ProductController{}
}

func InitCategoryProductController(db *gorm.DB) *CpController.CategoryController {
	wire.Build(
		CpRepository.NewCategoryRepository,
		CpService.NewCategoryService,
		CpController.NewCategoryController,
		ProvideValidator,
	)
	return &CpController.CategoryController{}
}

func InitShoppingCartController(db *gorm.DB) *CartController.ShoppingCartController {
	wire.Build(
		CartRepository.NewShoppingCartRepository,
		CartService.NewShoppingCartService,
		CartController.NewShoppingCartController,
		ProvideValidator,
	)
	return &CartController.ShoppingCartController{}
}
