package service

import (
	configdb "github.com/saichudin/golang-starter/config"
	"github.com/saichudin/golang-starter/model/domain"
	"github.com/saichudin/golang-starter/model/web"
)

type ProductServiceGormImpl struct {
}

var productsMysql []domain.ProductMysql

func NewProductServiceGorm() ProductServiceGorm {
	return &ProductServiceGormImpl{}
}

func (service *ProductServiceGormImpl) CreateGorm(request web.ProductRequest) domain.ProductMysql {
	// connect mysql
	var product domain.ProductMysql
	db := configdb.Conn()

	product.Name = request.Name
	product.Price = request.Price
	db.Create(&product) // pass pointer of data to Create

	return product
}

func (service *ProductServiceGormImpl) UpdateGorm(productId int, request web.ProductRequest) domain.ProductMysql {
	var product domain.ProductMysql
	db := configdb.Conn()
	db.First(&product, productId)
	product.Name = request.Name
	product.Price = request.Price
	db.Save(&product)

	return product
}

func (service *ProductServiceGormImpl) DeleteGorm(productId int) {
	var product domain.ProductMysql
	db := configdb.Conn()
	db.Delete(&product, productId)
}

func (service *ProductServiceGormImpl) FindByIdGorm(productId int) domain.ProductMysql {
	var product domain.ProductMysql
	
	db := configdb.Conn()
	db.First(&product, productId)

	return product
}

func (service *ProductServiceGormImpl) FindAllGorm() []domain.ProductMysql {
	db := configdb.Conn()
	var products []domain.ProductMysql
	// Get all records
	db.Find(&products)
	// SELECT * FROM users;

	return products
}
