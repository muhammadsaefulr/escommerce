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

type UserController struct {
	service  *service.UserService
	validate *validator.Validate
}

func NewUserController(service *service.UserService, validate *validator.Validate) *UserController {
	return &UserController{service: service, validate: validate}
}

// User Customer Register godoc
// @Tags UserCustomerRegister
// @Summary Create new user customer
// @Accept json
// @Param user body entity.User true "User data"
// @Produce json
// @Description Register new user customer
// @Router /user/register [post]
// @Success 200 {object} entity.User "Successfully created new user"
func (c *UserController) CreateUser(ctx *gin.Context) {
	var user entity.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.validate.Struct(user); err != nil {
		errors := validator_format.FormatValidator(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "errorValidation!", "error_validation": errors})
		return
	}

	_, err := c.service.GetUserByEmail(user.Email)

	if err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "User With This Email Already Exist !"})
		return
	}

	hashedPass, err := auth.HashBcryptPassword(user.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	user.Password = hashedPass

	result, err := c.service.CreateUser(&user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Success Creating New User !", "data": result})
}

// User Customer Auth godoc
// @Tags UserCustomerAuth
// @Summary User Customer Auth
// @Accept json
// @Param user body entity.AuthLoginUser true "User data"
// @Produce json
// @Description Authenticates a user and returns user data
// @Router /user/auth/login [post]
// @Success 200 {object} entity.UserDataReturnViews "Successfully authenticated user"
func (c *UserController) AuthLoginUser(ctx *gin.Context) {
	var login entity.AuthLoginUser

	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.validate.Struct(login); err != nil {
		errors := validator_format.FormatValidator(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "errorValidation!", "error_validation": errors})
		return
	}

	getUser, err := c.service.AuthLoginUser(&login)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User Data Not found in the database"})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	if !auth.CheckBcryptPassword(login.Password, getUser.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid User Password !"})
		return
	}

	token, err := auth.GenerateJwtToken(getUser.Email, getUser.Name)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	UserReturn := &entity.UserDataReturnViews{
		Name:   getUser.Name,
		Email:  getUser.Email,
		RoleId: getUser.RoleId,
	}

	ctx.JSON(http.StatusOK, gin.H{"jwtToken": token, "user_data": UserReturn})

}

// User Customer Get By Id godoc
// @Tags UserCustomerGetById
// @Summary Get user by id
// @Param id path string true "User id"
// @Accept */*
// @Produce json
// @Description Get user by id
// @Router /user/get/:id [get]
// @Success 200 {object} entity.User "Successfully get user"
func (c *UserController) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")

	result, err := c.service.GetUserById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Success !", "data": result})

}

// User Customer Update godoc
// @Tags UserCustomerUpdate
// @Summary Update user data
// @Param id path string true "User id"
// @Accept json
// @Param user body entity.User true "User data"
// @Produce json
// @Description Update user data
// @Router /user/update/:id [put]
// @Success 200 {object} entity.User "Successfully update user"
func (c *UserController) UpdateUserData(ctx *gin.Context) {
	var user *entity.User

	id := ctx.Param("id")

	_, err := c.service.GetUserById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Data Not Found In The Records !"})
		return
	}

	if err := c.validate.Struct(user); err != nil {
		errors := validator_format.FormatValidator(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "errorValidation!", "error_validation": errors})
		return
	}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.UpdateUserData(id, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Succes Updating Data", "updated_data": user})

}

// User Customer Delete godoc
// @Tags UserCustomerDelete
// @Summary Delete user by id
// @Param id path string true "User id"
// @Accept */*
// @Produce json
// @Description Delete user by id
// @Router /user/delete/:id [delete]
// @Success 200 {object} entity.User "Successfully delete user"
func (c *UserController) DeleteUserById(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := c.service.GetUserById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Data Not Found In The Records !"})
		return
	}

	if err := c.service.DeleteUserById(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Berhasil Menghapus Data !"})
}
