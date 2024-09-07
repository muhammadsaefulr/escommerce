package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/muhammadsaefulr/escommerce/internal/domain/entity"
	service "github.com/muhammadsaefulr/escommerce/internal/infras/service/cart"
)

type ShoppingCartController struct {
	service   *service.ShoppingCartService
	validator *validator.Validate
}

func NewShoppingCartController(service *service.ShoppingCartService, validator *validator.Validate) *ShoppingCartController {
	return &ShoppingCartController{service: service, validator: validator}
}

func (c *ShoppingCartController) AddShoppingCart(ctx *gin.Context) {
	var cart entity.ShoppingCart

	if err := ctx.ShouldBindJSON(&cart); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := c.service.AddShoppingCart(&cart)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Berhasil menambah item ke shopping cart", "dataa": data})

}

func (c *ShoppingCartController) UpdateShoppingCart(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Unavailable"})
}
