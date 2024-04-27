package main

import (
	"minang-kos-service/app"
	"minang-kos-service/controller"
	"minang-kos-service/helper"
	"minang-kos-service/repository"
	"minang-kos-service/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	countryRepository := repository.NewCountryRepository()
	countryService := service.NewCountryService(countryRepository, db)
	countryController := controller.NewCountryController(countryService)

	router := httprouter.New()
	router.POST("/api/countries", countryController.Create)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
