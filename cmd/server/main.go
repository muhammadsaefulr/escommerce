package main

import (
	"github.com/muhammadsaefulr/escommerce/cmd/api/router"
	"github.com/muhammadsaefulr/escommerce/internal/modules/database"
	di "github.com/muhammadsaefulr/escommerce/internal/modules/di/wire_gen"
)

func main() {
	db := database.NewGormDB()
	userController := di.InitUserController(db)
	productController := di.InitProductController(db)
	categoryProductController := di.InitCategoryProductController(db)
	cartController := di.InitShoppingCartController(db)
	r := router.SetupRouter(userController, productController, categoryProductController, cartController)
	r.Run()
}
