package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/muhammadsaefulr/escommerce/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestAddProduct(t *testing.T) {
	payloadProduct := entity.AddProductItems{
		ProductName:        "Shampoo",
		ProductDescription: "shampoo 1 dus 24pcs",
		ProductPrice:       64000,
		SellerId:           "01ef727e-7791-69b2-bbb2-008cfa73f649",
		CategoryId:         1,
	}

	jsonValue, err := json.Marshal(payloadProduct)
	assert.NoError(t, err, "Failed to marshal JSON payload")

	req, err := http.NewRequest("POST", "/api/product/add", bytes.NewBuffer(jsonValue))
	assert.NoError(t, err, "Failed to run request")

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200 OK")

	var response struct {
		Data    entity.ProductItems `json:"data"`
		Message string              `json:"message"`
	}

	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Failed to unmarshal response")

	productId = response.Data.ID

}

func TestGetProductById(t *testing.T) {

	params := fmt.Sprintf("/api/product/get/%s", productId)
	req, err := http.NewRequest("GET", params, nil)
	assert.NoError(t, err, "Failed to run request")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200 OK")
}

func TestUpdateProduct(t *testing.T) {
	payloadProduct := entity.AddProductItems{
		ProductName:        "Shampoo",
		ProductDescription: "shampoo 1 dus 12pcs",
		ProductPrice:       34000,
		SellerId:           "01ef727e-7791-69b2-bbb2-008cfa73f649",
		CategoryId:         1,
	}

	params := fmt.Sprintf("/api/product/update/%s", productId)

	jsonValue, err := json.Marshal(payloadProduct)
	assert.NoError(t, err, "Failed to marshal JSON payload")

	req, err := http.NewRequest("PUT", params, bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200 OK")

	// req.Header.Set("Content-Type", "application/json")

	// var response struct {
	// 	Data    entity.ProductItems `json:"data"`
	// 	Message string              `json:"message"`
	// }

	// err = json.Unmarshal(w.Body.Bytes(), &response)
	// assert.NoError(t, err, "Failed to unmarshal response")

}

func TestDeleteProduct(t *testing.T) {
	params := fmt.Sprintf("/api/product/delete/%s", productId)

	req, err := http.NewRequest("DELETE", params, nil)
	assert.NoError(t, err, "Failed to run request")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200 OK")

}
