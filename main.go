package main

import (
	"belajar-golang-REST/app"
	"belajar-golang-REST/controller"
	"belajar-golang-REST/helper"
	"belajar-golang-REST/middleware"
	"belajar-golang-REST/repository"
	"belajar-golang-REST/service"
	"github.com/go-playground/validator"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryControler(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
