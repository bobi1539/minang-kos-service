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
	endpoint.SetAuthEndpoint(router, getAuthController(db, validate))
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
	countryService := service.NewCountryService(getCountryRepository(), db, validate)
	return controller.NewCountryController(countryService)
}

func getProvinceController(db *sql.DB, validate *validator.Validate) controller.ProvinceController {
	provinceService := service.NewProvinceService(getProvinceRepository(), getCountryRepository(), db, validate)
	return controller.NewProvinceController(provinceService)
}

func getRoleController(db *sql.DB, validate *validator.Validate) controller.CountryController {
	roleService := service.NewRoleService(getRoleRepository(), db, validate)
	return controller.NewRoleController(roleService)
}

func getAuthController(db *sql.DB, validate *validator.Validate) controller.AuthController {
	authService := service.NewAuthService(getUserRepository(), db, validate)
	return controller.NewAuthController(authService)
}

func getCountryRepository() repository.CountryRepository {
	return repository.NewCountryRepository()
}

func getProvinceRepository() repository.ProvinceRepository {
	return repository.NewProvinceRepository()
}

func getRoleRepository() repository.RoleRepository {
	return repository.NewRoleRepository()
}

func getUserRepository() repository.UserRepository {
	return repository.NewUserRepository()
}
