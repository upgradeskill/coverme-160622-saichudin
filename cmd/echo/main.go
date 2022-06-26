package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/saichudin/golang-starter/model/web"
	"github.com/saichudin/golang-starter/service"
)

func main() {
	e := echo.New()

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
	productService := service.NewProductService()
	product := productService.FindAll()

	return c.JSON(http.StatusOK, product)
}

func handlerProductShow(c echo.Context) error {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	productService := service.NewProductService()
	product := productService.FindById(id)

	return c.JSON(http.StatusOK, product)
}

func handlerProductCreate(c echo.Context) (err error) {
	var productRequest web.ProductRequest
	if err = c.Bind(&productRequest); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	productService := service.NewProductService()
	product := productService.Create(productRequest)

	return c.JSON(http.StatusCreated, product)
}

func handlerProductUpdate(c echo.Context) (err error) {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	var productRequest web.ProductRequest
	if err = c.Bind(&productRequest); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	productService := service.NewProductService()
	product := productService.Update(id, productRequest)

	return c.JSON(http.StatusOK, product)
}

func handlerProductDelete(c echo.Context) error {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	productService := service.NewProductService()
	productService.Delete(id)

	return c.String(http.StatusOK, "Product Deleted")
}
