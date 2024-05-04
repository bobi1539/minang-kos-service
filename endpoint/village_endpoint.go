package endpoint

import (
	"database/sql"
	"minang-kos-service/controller"
	"minang-kos-service/repository"
	"minang-kos-service/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

const VILLAGES = "/api/villages"
const VILLAGES_ALL = VILLAGES + "/all"
const VILLAGES_VILLAGE = VILLAGES + "/village/:villageId"

func SetVillageEndpoint(router *httprouter.Router, db *sql.DB, validate *validator.Validate) {
	villageController := getVillageController(db, validate)

	router.POST(VILLAGES, villageController.Create)
	router.PUT(VILLAGES_VILLAGE, villageController.Update)
	router.GET(VILLAGES_ALL, villageController.FindAllWithoutPagination)
	router.GET(VILLAGES_VILLAGE, villageController.FindById)
	router.GET(VILLAGES, villageController.FindAllWithPagination)
	router.DELETE(VILLAGES_VILLAGE, villageController.Delete)
}

func getVillageController(db *sql.DB, validate *validator.Validate) controller.VillageController {
	villageService := service.NewVillageService(getVillageRepository(), getDistrictRepository(), db, validate)
	return controller.NewVillageController(villageService)
}

func getVillageRepository() repository.VillageRepository {
	return repository.NewVillageRepository()
}
