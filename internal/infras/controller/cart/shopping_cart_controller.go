package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/muhammadsaefulr/escommerce/internal/domain/entity"
	service "github.com/muhammadsaefulr/escommerce/internal/infras/service/cart"
	"gorm.io/gorm"
	// validator_format "github.com/muhammadsaefulr/escommerce/internal/modules/validator"
)

type ShoppingCartController struct {
	service   *service.ShoppingCartService
	validator *validator.Validate
}

func NewShoppingCartController(service *service.ShoppingCartService, validator *validator.Validate) *ShoppingCartController {
	return &ShoppingCartController{service: service, validator: validator}
}

func (c *ShoppingCartController) AddShoppingCartItem(ctx *gin.Context) {
	var CartItem entity.ShoppingCartItems

	if err := ctx.ShouldBindJSON(&CartItem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := c.service.GetShoppingCartItemById(CartItem.CartId, CartItem.ProductId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newItem := entity.ShoppingCartItems{
				CartId:    CartItem.CartId,
				ProductId: CartItem.ProductId,
				Quantity:  CartItem.Quantity,
			}

			_, err := c.service.AddShoppingCartItem(&newItem)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error Adding Product Item To Shopping Cart"})
				return
			}

		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error Fetching Shopping Cart Item"})
			return
		}
	} else {
		data.Quantity += CartItem.Quantity
		_, err := c.service.UpdateShoppingCartItem(CartItem.CartId, data)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error Updating Product Item In Shopping Cart"})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Berhasil menambah produk ke shopping cart!"})
}

func (c *ShoppingCartController) UpdateShoppingCart(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Unavailable"})
}

func (c *ShoppingCartController) GetShoppingCartById(ctx *gin.Context) {
	userId := ctx.Param("userId")

	data, err := c.service.GetShoppingCartById(userId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"messsage": "Cart tidak ditemukan !"})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Berhasil mengambil data cart", "data": data})

}
