package router

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadsaefulr/escommerce/cmd/api/middleware"
	"github.com/muhammadsaefulr/escommerce/internal/infras/controller"
	CartController "github.com/muhammadsaefulr/escommerce/internal/infras/controller/cart"
	CpController "github.com/muhammadsaefulr/escommerce/internal/infras/controller/product"
)

func SetupRouter(UserController *controller.UserController, ProductController *controller.ProductController, CategoryProductController *CpController.CategoryController, CartController *CartController.ShoppingCartController) *gin.Engine {
	r := gin.Default()

	RouterApiGroup := r.Group("/api")
	{
		UserGroup := RouterApiGroup.Group("/user")
		{
			UserGroup.POST("/auth/login", UserController.AuthLoginUser)
			UserGroup.POST("/register", UserController.CreateUser)
			UserGroup.GET("/get/:id", middleware.JwtAuth(), UserController.GetUserById)
			UserGroup.PUT("/update/:id", middleware.JwtAuth(), UserController.UpdateUserData)
			UserGroup.DELETE("/delete/:id", middleware.JwtAuth(), UserController.DeleteUserById)
		}
		ProductGroup := RouterApiGroup.Group("/product")
		{
			ProductGroup.POST("/add", ProductController.AddProductItems)
			ProductGroup.GET("/get/:id", ProductController.GetProductItems)
			ProductGroup.PUT("/update/:id", ProductController.UpdateProductItems)
			ProductGroup.DELETE("/delete/:id", ProductController.DeleteProductItems)

			CategoryProductGroup := ProductGroup.Group("/category")
			{
				CategoryProductGroup.POST("/add", CategoryProductController.CreateCategory)
				CategoryProductGroup.GET("/get/:id", CategoryProductController.GetCategoryById)
				CategoryProductGroup.PUT("/update/:id", CategoryProductController.UpdateCategoryData)
			}
		}
		ShoppingCartGroup := RouterApiGroup.Group("/cart/")
		{
			ShoppingCartGroup.POST("/add", CartController.AddShoppingCart)
		}
	}

	return r
}
