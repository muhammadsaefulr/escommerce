package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/muhammadsaefulr/escommerce/internal/domain/entity"
	service "github.com/muhammadsaefulr/escommerce/internal/infras/service/product"
	validator_format "github.com/muhammadsaefulr/escommerce/internal/modules/validator"
)

type CategoryController struct {
	service  *service.CategoryService
	validate *validator.Validate
}

func NewCategoryController(service *service.CategoryService, validate *validator.Validate) *CategoryController {
	return &CategoryController{service: service, validate: validate}
}

func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	var Category entity.CategoryProduct

	if err := ctx.ShouldBindJSON(&Category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.validate.Struct(Category); err != nil {
		errors := validator_format.FormatValidator(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "errorValidation!", "error_validation": errors})
		return
	}

	_, err := c.service.GetCategoryByName(Category.Name)

	if err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Category With This Name Is Exist !"})
		return
	}

	result, err := c.service.CreateCategory(&Category)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Success Creating New Category !", "data": result})
}

func (c *CategoryController) GetCategoryById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.service.GetCategoryById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Success !", "data": result})

}

func (c *CategoryController) UpdateCategoryData(ctx *gin.Context) {
	var Category *entity.CategoryProduct

	id, err := strconv.Atoi(ctx.Param("id"))

	log.Println(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Params id !"})
		return
	}

	_, err = c.service.GetCategoryById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Data Not Found In The Records !"})
		return
	}

	if err := ctx.ShouldBindJSON(&Category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.validate.Struct(Category); err != nil {
		errors := validator_format.FormatValidator(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "errorValidation!", "error_validation": errors})
		return
	}

	if err := c.service.UpdateCategoryData(id, Category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Succes Updating Data", "updated_data": Category})

}

func (c *CategoryController) DeleteCategoryById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	CategoryDb, err := c.service.GetCategoryById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Data Not Found In The Records !"})
		return
	}

	if err := c.service.DeleteCategoryById(int(CategoryDb.ID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Berhasil Menghapus Data !"})
}
