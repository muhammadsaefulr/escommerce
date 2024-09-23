package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"log"
	"os"

	"github.com/muhammadsaefulr/escommerce/cmd/api/router"
	"github.com/muhammadsaefulr/escommerce/internal/domain/entity"
	"github.com/muhammadsaefulr/escommerce/internal/modules/database"
	di "github.com/muhammadsaefulr/escommerce/internal/modules/di/wire_gen"
	"github.com/stretchr/testify/assert"
)

var r http.Handler
var jwtToken string
var userId string

func TestMain(m *testing.M) {
	db := database.NewGormDB()
	if db == nil {
		log.Fatal("failed to connect to database")
	}

	userController := di.InitUserController(db)
	userSellerController := di.InitUserSellerController(db)
	productController := di.InitProductController(db)
	categoryProductController := di.InitCategoryProductController(db)
	cartController := di.InitShoppingCartController(db)

	r = router.SetupRouter(userController, userSellerController, productController, categoryProductController, cartController)

	code := m.Run()
	os.Exit(code)
}

func TestRegisterUserCustomer(t *testing.T) {
	userRegisterPayload := entity.AuthRegisterUser{
		Name:     "Saepul",
		Email:    "epul@gmail.com",
		Password: "epul332",
	}

	jsonValue, err := json.Marshal(userRegisterPayload)
	if err != nil {
		t.Fatalf("Failed to marshal JSON payload: %v", err)
	}

	req, err := http.NewRequest("POST", "/api/user/register", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200 OK")

	var response struct {
		Data    entity.User `json:"data"`
		Message string      `json:"message"`
	}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	userRegisterTest := response.Data
	assert.NotEmpty(t, userRegisterTest.ID, "User ID should not be empty")
	assert.Equal(t, "Saepul", userRegisterTest.Name, "User name should match")
	assert.Equal(t, "epul@gmail.com", userRegisterTest.Email, "User email should match")

	userId = userRegisterTest.ID.String()
	log.Println(userId)
}

func TestLoginUserCustomer(t *testing.T) {
	userPayload := entity.AuthLoginUser{
		Email:    "epul@gmail.com",
		Password: "epul332",
	}

	jsonValue, err := json.Marshal(userPayload)
	if err != nil {
		t.Fatalf("Failed to marshal JSON payload: %v", err)
	}

	req, err := http.NewRequest("POST", "/api/user/auth/login", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200 OK")

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	jwtToken = response["jwtToken"]

	assert.NotEmpty(t, response["jwtToken"], "JWT token should be present in response")
}

func TestDeleteUserCustomer(t *testing.T) {

	params := fmt.Sprintf("/api/user/delete/%s", userId)

	req, err := http.NewRequest("DELETE", params, nil)

	if err != nil {
		t.Fatalf("Failed to run request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+jwtToken)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200 OK")
}
