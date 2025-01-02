package controller

import (
	"errors"
	"fmt"
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

// UserSeller Customer Register godoc
// @Tags UserSeller
// @Summary Create new user seller customer
// @Accept json
// @Param user body entity.UserSeller true "User seller data"
// @Produce json
// @Description Register new user seller customer
// @Router /user/seller/register [post]
// @Success 200 {object} entity.UserSeller "Successfully created new user seller"
func (c *UserSellerController) CreateUserSeller(ctx *gin.Context) {
	var userSeller entity.UserSeller

	// Bind JSON input to userSeller struct
	if err := ctx.ShouldBindJSON(&userSeller); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Print("This Are User ID was inputted: ", userSeller.UserID)

	// Validate input
	if err := c.validate.Struct(userSeller); err != nil {
		errors := validator_format.FormatValidator(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "errorValidation!", "error_validation": errors})
		return
	}

	_, err := c.service.GetUserSellerByUserId(userSeller.UserID.String())

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "User Seller with this user account already exists!"})
			return
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			_, err := c.service.CreateUserSeller(&userSeller)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Success creating new UserSeller!"})
}

// UserSeller Customer Auth Login godoc
// @Tags UserSeller
// @Summary UserSeller Customer Auth Login
// @Accept json
// @Param user body entity.AuthLoginUserSeller true "User seller data"
// @Produce json
// @Description Authenticates a user seller and returns user seller data
// @Router /user/seller/auth/login [post]
// @Success 200 {object} entity.UserSeller "Successfully authenticated user seller"
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

	getUser, err := c.service.GetUserByUserEmail(login.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "This User might be not registred As UserSeler"})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	getUserSeller, err := c.service.GetUserSellerByUserId(getUser.ID.String())

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "UserSeller not found in database"})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	if !auth.CheckBcryptPassword(login.Password, getUser.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid UserSeller Password !"})
		return
	}

	token, err := auth.GenerateJwtToken(getUser.Email, getUserSeller.NamaToko)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"jwtToken": token, "user_data": getUserSeller})

}

// UserSeller Customer Get By Id godoc
// @Tags UserSeller
// @Summary Get user seller by id
// @Param id path string true "User seller id"
// @Accept */*
// @Produce json
// @Description Get user seller by id
// @Router /user/seller/get/:id [get]
// @Success 200 {object} entity.UserSeller "Successfully get user seller"
func (c *UserSellerController) GetUserSellerById(ctx *gin.Context) {
	id := ctx.Param("id")

	result, err := c.service.GetUserSellerById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Success Request Data User Seller !", "data": result})

}

// User Seller Update godoc
// @Tags UserSeller
// @Summary Update User Seller Data
// @Param id path string true "User Seller id"
// @Accept json
// @Security BearerAuth
// @Param user body entity.UserSeller true "User Seller data"
// @Produce json
// @Description Update User Seller Data
// @Router /user/seller/update/:id [put]
// @Success 200 {object} entity.UserSeller "Successfully update User Seller"
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

// User Seller Delete godoc
// @Tags UserSeller
// @Summary Delete User Seller By Id
// @Param id path string true "User Seller id"
// @Accept */*
// @Produce json
// @Security BearerAuth
// @Description Delete User Seller By Id
// @Router /user/seller/delete/:id [delete]
// @Success 200 {object} entity.UserSeller "Successfully delete User Seller"
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
