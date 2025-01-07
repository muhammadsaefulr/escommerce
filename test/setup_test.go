package test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/muhammadsaefulr/escommerce/cmd/api/router"
	"github.com/muhammadsaefulr/escommerce/internal/modules/database"
	di "github.com/muhammadsaefulr/escommerce/internal/modules/di/wire_gen"
)

var (
	r         http.Handler
	jwtToken  string
	TokoId    string
	UserId    string
	productId string
)

func TestMain(m *testing.M) {
	db := database.NewGormDB()
	if db == nil {
		log.Fatal("failed to connect to database")
	}

	userController := di.InitUserController(db)
	userSellerController := di.InitUserSellerController(db)
	productController := di.InitProductController(db)
	categoryProductController := di.InitCategoryProductController(db)
	cartController := di.InitShoppingCartController(db)

	r = router.SetupRouter(userController, userSellerController, productController, categoryProductController, cartController)

	code := m.Run()
	os.Exit(code)
}
