package service

import (
	"github.com/saichudin/golang-starter/model/domain"
	"github.com/saichudin/golang-starter/model/web"
)

type ProductServiceGorm interface {
	CreateGorm(request web.ProductRequest) domain.ProductMysql
	UpdateGorm(productId int, request web.ProductRequest) domain.ProductMysql
	DeleteGorm(productId int)
	FindByIdGorm(productId int) domain.ProductMysql
	FindAllGorm() []domain.ProductMysql
}
