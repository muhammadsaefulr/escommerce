package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/muhammadsaefulr/escommerce/internal/domain/entity"
	"github.com/muhammadsaefulr/escommerce/internal/infras/service"
	"github.com/muhammadsaefulr/escommerce/internal/modules/auth"
	validator_format "github.com/muhammadsaefulr/escommerce/internal/modules/validator"
	"gorm.io/gorm"
)

type UserSellerController struct {
	service  *service.UserSellerService
	validate *validator.Validate
}

func NewUserSellerController(service *service.UserSellerService, validate *validator.Validate) *UserSellerController {
	return &UserSellerController{service: service, validate: validate}
}

func (c *UserSellerController) CreateUserSeller(ctx *gin.Context) {
	var UserSeller entity.UserSeller

	if err := ctx.ShouldBindJSON(&UserSeller); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.validate.Struct(UserSeller); err != nil {
		errors := validator_format.FormatValidator(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "errorValidation!", "error_validation": errors})
		return
	}

	_, err := c.service.GetUserSellerByEmail(UserSeller.Email)

	if err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "UserSeller With This Email Already Exist !"})
		return
	}

	hashedPass, err := auth.HashBcryptPassword(UserSeller.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	UserSeller.Password = hashedPass

	result, err := c.service.CreateUserSeller(&UserSeller)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Success Creating New UserSeller !", "data": result})
}

func (c *UserSellerController) AuthLoginUserSeller(ctx *gin.Context) {
	var login entity.AuthLoginUserSeller

	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.validate.Struct(login); err != nil {
		errors := validator_format.FormatValidator(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "errorValidation!", "error_validation": errors})
		return
	}

	getUserSeller, err := c.service.AuthLoginUserSeller(&login)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "UserSeller Data Not found in the database"})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	if !auth.CheckBcryptPassword(login.Password, getUserSeller.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid UserSeller Password !"})
		return
	}

	token, err := auth.GenerateJwtToken(getUserSeller.Email, getUserSeller.Name)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"jwtToken": token})

}

func (c *UserSellerController) GetUserSellerById(ctx *gin.Context) {
	id := ctx.Param("id")

	result, err := c.service.GetUserSellerById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Success !", "data": result})

}

func (c *UserSellerController) UpdateUserSellerData(ctx *gin.Context) {
	var UserSeller *entity.UserSeller

	id := ctx.Param("id")

	_, err := c.service.GetUserSellerById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Data Not Found In The Records !"})
		return
	}

	if err := c.validate.Struct(UserSeller); err != nil {
		errors := validator_format.FormatValidator(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "errorValidation!", "error_validation": errors})
		return
	}

	if err := ctx.ShouldBindJSON(&UserSeller); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.UpdateUserSellerData(id, UserSeller); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Succes Updating Data", "updated_data": UserSeller})

}

func (c *UserSellerController) DeleteUserSellerById(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := c.service.GetUserSellerById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Data Not Found In The Records !"})
		return
	}

	if err := c.service.DeleteUserSellerById(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Berhasil Menghapus Data !"})
}
