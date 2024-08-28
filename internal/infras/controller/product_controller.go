package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadsaefulr/escommerce/internal/domain/entity"
	"github.com/muhammadsaefulr/escommerce/internal/infras/service"
)

type ProductController struct {
	service *service.ProductService
}

func NewProductController(service *service.ProductService) *ProductController {
	return &ProductController{service: service}
}

func (c *ProductController) AddProductItems(ctx *gin.Context) {
	var product entity.ProductItems

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	result, err := c.service.AddProductItems(&product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Berhasil Membuat Product !", "data": result})
}

func (c *ProductController) GetProductItems(ctx *gin.Context) {
	IdProduct := ctx.Param("id")

	result, err := c.service.GetProductItems(IdProduct)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Berhasil Mengambil Data Product !", "data": result})

}
