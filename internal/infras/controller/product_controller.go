package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/muhammadsaefulr/escommerce/internal/domain/entity"
	"github.com/muhammadsaefulr/escommerce/internal/infras/service"
	validator_format "github.com/muhammadsaefulr/escommerce/internal/modules/validator"
)

type ProductController struct {
	service  *service.ProductService
	validate *validator.Validate
}

func NewProductController(service *service.ProductService, validate *validator.Validate) *ProductController {
	return &ProductController{service: service, validate: validate}
}

func (c *ProductController) AddProductItems(ctx *gin.Context) {
	var product entity.ProductItems

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if err := c.validate.Struct(product); err != nil {
		errors := validator_format.FormatValidator(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "errorValidation!", "error_validation": errors})
		return
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

func (c *ProductController) UpdateProductItems(ctx *gin.Context) {
	idProduct := ctx.Param("id")

	var product *entity.ProductItems

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	getProduct, err := c.service.GetProductItems(idProduct)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	product, err = c.service.UpdateProductItems(idProduct, getProduct, product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Berhasil Mengupdate Product", "data": product})

}

func (c *ProductController) DeleteProductItems(ctx *gin.Context) {
	IdProduct := ctx.Param("id")

	product, err := c.service.GetProductItems(IdProduct)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Product Tidak Ditemukan !"})
		return
	}

	if err := c.service.DeleteProductItems(product.ID.String()); err != nil {

		log.Println("error deleting product exceeds", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Berhasil Menghapus Data !", "removed_data": product})
}
