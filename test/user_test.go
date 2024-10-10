package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"log"

	"github.com/muhammadsaefulr/escommerce/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUserCustomer(t *testing.T) {
	userRegisterPayload := entity.AuthRegisterUser{
		Name:     "Saepul",
		Email:    "epul@gmail.com",
		Password: "epul332",
	}

	jsonValue, err := json.Marshal(userRegisterPayload)
	assert.NoError(t, err, "Failed to marshal JSON payload")

	req, err := http.NewRequest("POST", "/api/user/register", bytes.NewBuffer(jsonValue))
	assert.NoError(t, err, "Failed to create request")

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200 OK")

	var response struct {
		Data    entity.User `json:"data"`
		Message string      `json:"message"`
	}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Failed to unmarshal response")

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
	assert.NoError(t, err, "Failed to marshal JSON payload")

	req, err := http.NewRequest("POST", "/api/user/auth/login", bytes.NewBuffer(jsonValue))
	assert.NoError(t, err, "Failed to create request")
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200 OK")

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Failed to unmarshal response")

	jwtToken = response["jwtToken"]

	assert.NotEmpty(t, response["jwtToken"], "JWT token should be present in response")
}

func TestDeleteUserCustomer(t *testing.T) {
	params := fmt.Sprintf("/api/user/delete/%s", userId)

	req, err := http.NewRequest("DELETE", params, nil)
	assert.NoError(t, err, "Failed to create request")
	req.Header.Set("Authorization", "Bearer "+jwtToken)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200 OK")
}