package service

import (
	"github.com/saichudin/golang-starter/model/domain"
	"github.com/saichudin/golang-starter/model/web"
)

type ProductService interface {
	Create(request web.ProductRequest) domain.Product
	Update(productId int, request web.ProductRequest) domain.Product
	Delete(productId int)
	FindById(productId int) domain.Product
	FindAll() []domain.Product
}
