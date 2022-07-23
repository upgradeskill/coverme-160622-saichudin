package service

import (
	"github.com/saichudin/golang-starter/model/domain"
	"github.com/saichudin/golang-starter/model/web"
)

type ProductServiceImpl struct {
}

var products []domain.Product
var productId int = 0

func NewProductService() ProductService {
	return &ProductServiceImpl{}
}

func (service *ProductServiceImpl) Create(request web.ProductRequest) domain.Product {
	var product domain.Product

	productId++
	product.Id = productId
	product.Name = request.Name
	product.Price = request.Price
	products = append(products, product)

	return product
}

func (service *ProductServiceImpl) Update(productId int, request web.ProductRequest) domain.Product {
	var result domain.Product
	for index, product := range products {
		if product.Id == productId {
			products[index].Name = request.Name
			products[index].Price = request.Price

			result.Id = productId
			result.Name = request.Name
			result.Price = request.Price
		}
	}

	return result
}

func (service *ProductServiceImpl) Delete(productId int) {
	for index, product := range products {
		if product.Id == productId {
			products = append(products[:index], products[index+1:]...)
		}
	}
}

func (service *ProductServiceImpl) FindById(productId int) domain.Product {
	var result domain.Product
	for _, product := range products {
		if product.Id == productId {
			result.Id = productId
			result.Name = product.Name
			result.Price = product.Price
		}
	}

	return result
}

func (service *ProductServiceImpl) FindAll() []domain.Product {
	return products
}
