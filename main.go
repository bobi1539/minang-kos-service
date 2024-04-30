package main

import (
	"database/sql"
	"minang-kos-service/app"
	"minang-kos-service/controller"
	"minang-kos-service/endpoint"
	"minang-kos-service/exception"
	"minang-kos-service/helper"
	"minang-kos-service/middleware"
	"minang-kos-service/repository"
	"minang-kos-service/service"
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
	endpoint.SetCountryEndpoint(router, getCountryController(db, validate))
	endpoint.SetProvinceEndpoint(router, getProvinceController(db, validate))
	endpoint.SetRoleEndpoint(router, getRoleController(db, validate))
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

func getCountryController(db *sql.DB, validate *validator.Validate) controller.CountryController {
	countryRepository := repository.NewCountryRepository()
	countryService := service.NewCountryService(countryRepository, db, validate)
	return controller.NewCountryController(countryService)
}

func getProvinceController(db *sql.DB, validate *validator.Validate) controller.ProvinceController {
	provinceRepository := repository.NewProvinceRepository()
	countryRepository := repository.NewCountryRepository()
	provinceService := service.NewProvinceService(provinceRepository, countryRepository, db, validate)
	return controller.NewProvinceController(provinceService)
}

func getRoleController(db *sql.DB, validate *validator.Validate) controller.CountryController {
	roleRepository := repository.NewRoleRepository()
	roleService := service.NewRoleService(roleRepository, db, validate)
	return controller.NewRoleController(roleService)
}
