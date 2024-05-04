package endpoint

import (
	"database/sql"
	"minang-kos-service/controller"
	"minang-kos-service/repository"
	"minang-kos-service/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

const KOS_BEDROOMS = "/api/kos-bedrooms"
const KOS_BEDROOMS_ALL = KOS_BEDROOMS + "/all"
const KOS_BEDROOMS_KOS_BEDROOM = KOS_BEDROOMS + "/kos-bedroom/:kosBedroomId"

func SetKosBedroomEndpoint(router *httprouter.Router, db *sql.DB, validate *validator.Validate) {
	kosBedroomController := getKosBedroomController(db, validate)

	router.POST(KOS_BEDROOMS, kosBedroomController.Create)
	router.PUT(KOS_BEDROOMS_KOS_BEDROOM, kosBedroomController.Update)
	router.GET(KOS_BEDROOMS_ALL, kosBedroomController.FindAllWithoutPagination)
	router.GET(KOS_BEDROOMS_KOS_BEDROOM, kosBedroomController.FindById)
	router.GET(KOS_BEDROOMS, kosBedroomController.FindAllWithPagination)
	router.DELETE(KOS_BEDROOMS_KOS_BEDROOM, kosBedroomController.Delete)
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

func getKosBedroomRepository() repository.KosBedroomRepository {
	return repository.NewKosBedroomRepository()
}

func getKosFacilityRepository() repository.KosFacilityRepository {
	return repository.NewKosFacilityRepository()
}
