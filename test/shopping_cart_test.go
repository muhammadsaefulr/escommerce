package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetShoppingCart(t *testing.T) {
	userIds := "01ef727f-7039-6cc6-bbb2-008cfa73f649"
	params := fmt.Sprintf("/api/cart/get/%s", userIds)
	req, err := http.NewRequest("GET", params, nil)

	assert.NoError(t, err, "Failed To run request")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200 OK")

}
