package service

import (
	"testing"

	"github.com/saichudin/golang-starter/model/web"
	"github.com/stretchr/testify/assert"
)

func TestFindAll(t *testing.T) {
	request := web.ProductRequest{
		Name:  "es jus",
		Price: 300,
	}

	// insert data
	productService := NewProductService()
	productService.Create(request)

	request.Name = "es teh"
	request.Price = 500
	productService.Create(request)

	// get all data
	resp := productService.FindAll()

	// asert all data len must be 2
	assert.Equal(t, 2, len(resp))
}

func TestCreate(t *testing.T) {
	request := web.ProductRequest{
		Name:  "es jus",
		Price: 300,
	}
	// create data
	productService := NewProductService()
	resp := productService.Create(request)

	// assert data must be equal with created data
	assert.Equal(t, "es jus", resp.Name)
	assert.Equal(t, 300, resp.Price)
}

func TestUpdate(t *testing.T) {
	request := web.ProductRequest{
		Name:  "es jus",
		Price: 300,
	}

	// create data
	productService := NewProductService()
	productService.Create(request)

	// update data
	request.Name = "es jus buah"
	request.Price = 500
	resp := productService.Update(1, request)

	// assert data must be equal with last update
	assert.Equal(t, "es jus buah", resp.Name)
	assert.Equal(t, 500, resp.Price)
}

func TestFindById(t *testing.T) {
	request := web.ProductRequest{
		Name:  "es jus",
		Price: 300,
	}

	// insert data
	productService := NewProductService()
	productService.Create(request)

	// get data by id 1 (last created)
	resp := productService.FindById(1)

	// assert product id must be 1
	assert.Equal(t, 1, resp.Id)
}

func TestDelete(t *testing.T) {
	request := web.ProductRequest{
		Name:  "es jus",
		Price: 300,
	}

	// insert data
	productService := NewProductService()
	productService.Create(request)

	// delete data
	productService.Delete(1)
	resp := productService.FindById(1)

	// assert product id must be 0 after delete
	assert.Equal(t, 0, resp.Id)
}
