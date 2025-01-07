package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetShoppingCart(t *testing.T) {
	userIds := UserId
	params := fmt.Sprintf("/api/cart/get/%s", userIds)

	fmt.Println("paramset: " + params)
	req, err := http.NewRequest("GET", params, nil)

	assert.NoError(t, err, "Failed To run request")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200 OK")

}
