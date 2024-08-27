package di

import (
	"github.com/google/wire"
	"github.com/muhammadsaefulr/escommerce/internal/controller"
	"github.com/muhammadsaefulr/escommerce/internal/repository"
	"github.com/muhammadsaefulr/escommerce/internal/service"
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
