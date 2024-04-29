package main

import (
	"minang-kos-service/app"
	"minang-kos-service/controller"
	"minang-kos-service/exception"
	"minang-kos-service/helper"
	"minang-kos-service/repository"
	"minang-kos-service/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	countryRepository := repository.NewCountryRepository()
	countryService := service.NewCountryService(countryRepository, db, validate)
	countryController := controller.NewCountryController(countryService)

	router := httprouter.New()
	router.POST("/api/countries", countryController.Create)
	router.PUT("/api/countries/:countryId", countryController.Update)
	router.GET("/api/countries/:countryId", countryController.FindById)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
