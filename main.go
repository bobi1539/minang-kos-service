package main

import (
	"database/sql"
	"minang-kos-service/app"
	"minang-kos-service/controller"
	"minang-kos-service/endpoint"
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

	router := httprouter.New()
	endpoint.SetCountryEndpoint(router, getCountryController(db, validate))
	// router.PanicHandler = exception.ErrorHandler

	runServer(router)
}

func getCountryController(db *sql.DB, validate *validator.Validate) controller.CountryController {
	countryRepository := repository.NewCountryRepository()
	countryService := service.NewCountryService(countryRepository, db, validate)
	return controller.NewCountryController(countryService)
}

func runServer(router *httprouter.Router) {
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
