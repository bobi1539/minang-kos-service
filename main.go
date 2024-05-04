package main

import (
	"minang-kos-service/app"
	"minang-kos-service/endpoint"
	"minang-kos-service/exception"
	"minang-kos-service/helper"
	"minang-kos-service/middleware"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	router := httprouter.New()
	endpoint.SetCountryEndpoint(router, db, validate)
	endpoint.SetProvinceEndpoint(router, db, validate)
	endpoint.SetCityEndpoint(router, db, validate)
	endpoint.SetDistrictEndpoint(router, db, validate)
	endpoint.SetVillageEndpoint(router, db, validate)
	endpoint.SetRoleEndpoint(router, db, validate)
	endpoint.SetAuthEndpoint(router, db, validate)
	endpoint.SetKosTypeEndpoint(router, db, validate)
	endpoint.SetFacilityTypeEndpoint(router, db, validate)
	endpoint.SetFacilityEndpoint(router, db, validate)
	endpoint.SetKosBedroomEndpoint(router, db, validate)
	endpoint.SetImageEndpoint(router)
	router.PanicHandler = exception.ErrorHandler

	runServer(router)
}

func runServer(router *httprouter.Router) {
	server := http.Server{
		Addr:    os.Getenv("APP_HOST"),
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
