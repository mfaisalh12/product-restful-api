package main

import (
	"fmt"
	"mfaisalh12/product-restful-api/app"
	"mfaisalh12/product-restful-api/controller"
	"mfaisalh12/product-restful-api/exception"
	"mfaisalh12/product-restful-api/helper"
	"mfaisalh12/product-restful-api/repository"
	"mfaisalh12/product-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productController := controller.NewProductController(productService)

	router := httprouter.New()

	router.GET("/api/products", productController.FindAll)
	router.GET("/api/products/:productId", productController.FindById)
	router.POST("/api/products", productController.Create)
	router.PUT("/api/products/:productId", productController.Update)
	router.DELETE("/api/products/:productId", productController.Delete)

	router.PanicHandler = exception.ErrorHandler
	
	server := http.Server{
		Addr: "localhost:3000",
		Handler: router,
	}
	fmt.Println("Server Running")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}