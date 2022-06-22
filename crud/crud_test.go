package crud

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerProducts(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/products/", nil)
	w := httptest.NewRecorder()
	handlerProducts(w, req)
	res := w.Result()

	assert.Equal(t, 200, res.StatusCode)
}

func TestHandlerProductShow(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/product/?id=1", nil)
	w := httptest.NewRecorder()
	handlerProduct(w, req)
	res := w.Result()

	assert.Equal(t, 200, res.StatusCode)
}

func TestHandlerProductCreate(t *testing.T) {
	body := strings.NewReader(`{"Name" : "Es Buah", "Price" : 700}`)
	req := httptest.NewRequest(http.MethodPost, "/product/", body)
	w := httptest.NewRecorder()
	handlerProduct(w, req)
	res := w.Result()

	assert.Equal(t, 200, res.StatusCode)
}

func TestHandlerProductUpdate(t *testing.T) {
	body := strings.NewReader(`{"Name" : "Es Buah Manis", "Price" : 600}`)
	req := httptest.NewRequest(http.MethodPut, "/product/", body)
	w := httptest.NewRecorder()
	handlerProduct(w, req)
	res := w.Result()

	assert.Equal(t, 200, res.StatusCode)
}

func TestHandlerProductDelete(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/product/?id=1", nil)
	w := httptest.NewRecorder()
	handlerProduct(w, req)
	res := w.Result()

	assert.Equal(t, 200, res.StatusCode)
}