package main

import (
	"github.com/muhammadsaefulr/escommerce/cmd/api/router"
	docs "github.com/muhammadsaefulr/escommerce/docs"
	"github.com/muhammadsaefulr/escommerce/internal/modules/database"

	di "github.com/muhammadsaefulr/escommerce/internal/modules/di/wire_gen"

	swaggerfiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

// @host localhost:8080
// @BasePath /api
// @schemes http
// @title Escommerce API
// @version 1.0
// @description Test Application

func main() {
	db := database.NewGormDB()
	userController := di.InitUserController(db)
	userSellerController := di.InitUserSellerController(db)
	productController := di.InitProductController(db)
	categoryProductController := di.InitCategoryProductController(db)
	cartController := di.InitShoppingCartController(db)

	docs.SwaggerInfo.Title = "Escommerce API"
	docs.SwaggerInfo.Description = "Escommerce API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"

	r := router.SetupRouter(userController, userSellerController, productController, categoryProductController, cartController)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}
