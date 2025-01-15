package controller

import (
	"fmt"
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

// Add Product Item Controller godoc
// @Tags ProductItems
// @Summary Add Product Items
// @Description Sellers can add product items if they are authorized
// @Accept json
// @Produce json
// @Param product body entity.AddProductItems true "Product Items"
// @Success 200 {object} entity.FilteredProductReturn
// @Router /product/add [post]
func (c *ProductController) AddProductItems(ctx *gin.Context) {
	var insertProduct entity.AddProductItems
	sellerIdRaw, _ := ctx.Get("user_id")
	sellerId := fmt.Sprintf("%v", sellerIdRaw)

	if err := ctx.ShouldBindJSON(&insertProduct); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if err := c.validate.Struct(insertProduct); err != nil {
		errors := validator_format.FormatValidator(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "errorValidation!", "error_validation": errors})
		return
	}

	_, err := c.service.GetSellerById(sellerId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Seller User Dengan ID Tersebut Tidak Ditemukan !"})
		return
	}

	product := entity.ProductItems{
		ProductName:        insertProduct.ProductName,
		ProductDescription: insertProduct.ProductDescription,
		ProductPrice:       insertProduct.ProductPrice,
		CategoryId:         insertProduct.CategoryId,
		SellerId:           sellerId,
	}

	result, err := c.service.AddProductItems(&product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	filteredResult := entity.FilteredProductReturn{
		ID:        result.ID,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
		Name:      result.ProductName,
		Price:     result.ProductPrice,
		SellerID:  result.SellerId,
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Berhasil Membuat Product !", "data": filteredResult})
}

// Get All Product With Seller Id godoc
// @Tags ProductItems
// @Summary Get All Product With Seller Id
// @Description Sellers can get all product items if they are authorized
// @Accept */*
// @Produce json
// @Param sellerId path string true "Seller Id"
// @Success 200 {object} entity.ProductItems
// @Router /product/get/all/{sellerId} [get]
func (c *ProductController) GetAllProductWithSellerId(ctx *gin.Context) {
	sellerId := ctx.Param("sellerId")
	result, err := c.service.GetAllProductBySellerId(sellerId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Berhasil Mengambil Semua Data !", "data": result})

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

	var productUpdate entity.UpdateProductItems
	var product *entity.ProductItems

	if err := ctx.ShouldBindJSON(&productUpdate); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	getProduct, err := c.service.GetProductItems(idProduct)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	product = &entity.ProductItems{
		ID:                 idProduct,
		ProductName:        productUpdate.ProductName,
		ProductDescription: productUpdate.ProductDescription,
		ProductPrice:       productUpdate.ProductPrice,
		CategoryId:         productUpdate.CategoryId,
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

	if err := c.service.DeleteProductItems(product.ID); err != nil {

		log.Println("error deleting product exceeds", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Berhasil Menghapus Data !", "removed_data": product})
}
