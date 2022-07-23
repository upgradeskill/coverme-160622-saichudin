package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/saichudin/golang-starter/model/domain"
	"github.com/saichudin/golang-starter/model/web"
	"github.com/saichudin/golang-starter/service"
	"github.com/saichudin/golang-starter/config"
)

func main() {
	e := echo.New()

	// connect mysql
	db := configdb.Conn()

	// Migrate the schema
	db.AutoMigrate(&domain.ProductMysql{})

	e.GET("/", hello)
	e.GET("/product", handlerProducts)
	e.GET("/product/:id", handlerProductShow)
	e.POST("/product", handlerProductCreate)
	e.PUT("/product/:id", handlerProductUpdate)
	e.DELETE("/product/:id", handlerProductDelete)

	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func handlerProducts(c echo.Context) error {
	productService := service.NewProductServiceGorm()
	product := productService.FindAllGorm()

	return c.JSON(http.StatusOK, product)
}

func handlerProductShow(c echo.Context) error {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	productService := service.NewProductServiceGorm()
	product := productService.FindByIdGorm(id)

	return c.JSON(http.StatusOK, product)
}

func handlerProductCreate(c echo.Context) (err error) {
	var productRequest web.ProductRequest
	if err = c.Bind(&productRequest); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	productService := service.NewProductServiceGorm()
	product := productService.CreateGorm(productRequest)

	return c.JSON(http.StatusCreated, product)
}

func handlerProductUpdate(c echo.Context) (err error) {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	var productRequest web.ProductRequest
	if err = c.Bind(&productRequest); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	productService := service.NewProductServiceGorm()
	product := productService.UpdateGorm(id, productRequest)

	return c.JSON(http.StatusOK, product)
}

func handlerProductDelete(c echo.Context) error {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	productService := service.NewProductServiceGorm()
	productService.DeleteGorm(id)

	return c.String(http.StatusOK, "Product Deleted")
}
