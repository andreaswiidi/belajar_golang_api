package main

import (
	"belajar-golang-REST/app"
	"belajar-golang-REST/controller"
	"belajar-golang-REST/exception"
	"belajar-golang-REST/helper"
	"belajar-golang-REST/repository"
	"belajar-golang-REST/service"
	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryControler(categoryService)

	router := httprouter.New()
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
