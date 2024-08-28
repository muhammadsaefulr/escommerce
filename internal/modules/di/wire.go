package di

import (
	"github.com/google/wire"
	"github.com/muhammadsaefulr/escommerce/internal/infras/controller"
	"github.com/muhammadsaefulr/escommerce/internal/infras/repository"
	"github.com/muhammadsaefulr/escommerce/internal/infras/service"

	"gorm.io/gorm"
)

func InitUserController(db *gorm.DB) *controller.UserController {
	wire.Build(
		repository.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)
	return &controller.UserController{}
}

func InitProductController(db *gorm.DB) *controller.ProductController {
	wire.Build(
		repository.NewProductRepository,
		service.NewProductService,
		controller.NewProductController,
	)

	return &controller.ProductController{}
}
