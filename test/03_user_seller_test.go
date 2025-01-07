package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/muhammadsaefulr/escommerce/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

type LoginUserSellerResponse struct {
	JwtToken string `json:"jwtToken"`
	message  string
}

func TestRegisterUserSeller(t *testing.T) {

	parseUUID, error := uuid.Parse("01efc758-c83d-6a77-aebe-008cfa73f649")

	if error != nil {
		assert.NoError(t, error, "Failed to parse UUID")
	}

	userRegisterPayload := entity.UserSellerRegister{
		NamaToko: "TokoTest",
		UserID:   parseUUID,
	}

	jsonValue, err := json.Marshal(userRegisterPayload)
	assert.NoError(t, err, "Failed to marshal JSON payload")

	req, err := http.NewRequest("POST", "/api/user/seller/register", bytes.NewBuffer(jsonValue))
	assert.NoError(t, err, "Failed to create request")

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200 OK")

	var response struct {
		Message string `json:"message"`
		TokoId  string `json:"TokoId"`
	}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Failed to unmarshal response")

	userRegisterTest := response
	TokoId = userRegisterTest.TokoId
	assert.NotEmpty(t, userRegisterTest.TokoId, "Toko ID should not be empty")
}

func TestLoginUserSeller(t *testing.T) {
	userPayload := entity.AuthLoginUser{
		Email:    "seller@gmail.com",
		Password: "testcommerce123",
	}

	jsonValue, err := json.Marshal(userPayload)
	assert.NoError(t, err, "Failed to marshal JSON payload")

	req, err := http.NewRequest("POST", "/api/user/seller/auth/login", bytes.NewBuffer(jsonValue))
	assert.NoError(t, err, "Failed to create request")
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200 OK")

	var response LoginUserSellerResponse
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Failed to unmarshal response")
	assert.NotEmpty(t, response.JwtToken, "Jwt Token should be present in response")

	jwtToken = response.JwtToken
}

func TestDeleteUserSeller(t *testing.T) {
	params := fmt.Sprintf("/api/user/seller/delete/%s", TokoId)

	req, err := http.NewRequest("DELETE", params, nil)
	assert.NoError(t, err, "Failed to create request")
	req.Header.Set("Authorization", "Bearer "+jwtToken)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200 OK")
}
