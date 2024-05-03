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
	endpoint.SetCityEndpoint(router, getCityController(db, validate))
	endpoint.SetDistrictEndpoint(router, getDistrictController(db, validate))
	endpoint.SetVillageEndpoint(router, getVillageController(db, validate))
	endpoint.SetRoleEndpoint(router, getRoleController(db, validate))
	endpoint.SetAuthEndpoint(router, getAuthController(db, validate))
	endpoint.SetKosTypeEndpoint(router, getKosTypeController(db, validate))
	endpoint.SetFacilityTypeEndpoint(router, getFacilityTypeController(db, validate))
	endpoint.SetFacilityEndpoint(router, getFacilityController(db, validate))
	endpoint.SetKosBedroomEndpoint(router, getKosBedroomController(db, validate))
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

func getCityController(db *sql.DB, validate *validator.Validate) controller.CityController {
	cityService := service.NewCityService(getCityRepository(), getProvinceRepository(), db, validate)
	return controller.NewCityController(cityService)
}

func getDistrictController(db *sql.DB, validate *validator.Validate) controller.DistrictController {
	districtService := service.NewDistrictService(getDistrictRepository(), getCityRepository(), db, validate)
	return controller.NewDistrictController(districtService)
}

func getVillageController(db *sql.DB, validate *validator.Validate) controller.VillageController {
	villageService := service.NewVillageService(getVillageRepository(), getDistrictRepository(), db, validate)
	return controller.NewVillageController(villageService)
}

func getKosTypeController(db *sql.DB, validate *validator.Validate) controller.KosTypeController {
	kosTypeService := service.NewKosTypeService(getKosTypeRepository(), db, validate)
	return controller.NewKosTypeController(kosTypeService)
}

func getFacilityTypeController(db *sql.DB, validate *validator.Validate) controller.FacilityTypeController {
	facilityTypeService := service.NewFacilityTypeService(getFacilityTypeRepository(), db, validate)
	return controller.NewFacilityTypeController(facilityTypeService)
}

func getFacilityController(db *sql.DB, validate *validator.Validate) controller.FacilityController {
	facilityService := service.NewFacilityService(getFacilityRepository(), getFacilityTypeRepository(), db, validate)
	return controller.NewFacilityController(facilityService)
}

func getKosBedroomController(db *sql.DB, validate *validator.Validate) controller.KosBedroomController {
	kosBedroomService := service.NewKosBedroomService(
		getKosBedroomRepository(),
		getKosTypeRepository(),
		getVillageRepository(),
		getUserRepository(),
		getKosFacilityRepository(),
		getFacilityRepository(),
		db,
		validate,
	)
	return controller.NewKosBedroomController(kosBedroomService)
}

func getCountryRepository() repository.CountryRepository {
	return repository.NewCountryRepository()
}

func getProvinceRepository() repository.ProvinceRepository {
	return repository.NewProvinceRepository()
}

func getCityRepository() repository.CityRepository {
	return repository.NewCityRepository()
}

func getDistrictRepository() repository.DistrictRepository {
	return repository.NewDistrictRepository()
}

func getRoleRepository() repository.RoleRepository {
	return repository.NewRoleRepository()
}

func getUserRepository() repository.UserRepository {
	return repository.NewUserRepository()
}

func getVillageRepository() repository.VillageRepository {
	return repository.NewVillageRepository()
}

func getKosTypeRepository() repository.KosTypeRepository {
	return repository.NewKosTypeRepository()
}

func getFacilityTypeRepository() repository.FacilityTypeRepository {
	return repository.NewFacilityTypeRepository()
}

func getFacilityRepository() repository.FacilityRepository {
	return repository.NewFacilityRepository()
}

func getKosBedroomRepository() repository.KosBedroomRepository {
	return repository.NewKosBedroomRepository()
}

func getKosFacilityRepository() repository.KosFacilityRepository {
	return repository.NewKosFacilityRepository()
}
