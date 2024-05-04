package service

import (
	"context"
	"database/sql"
	"minang-kos-service/exception"
	"minang-kos-service/helper"
	"minang-kos-service/model/domain"
	"minang-kos-service/model/web/request"
	"minang-kos-service/model/web/search"
	"minang-kos-service/repository"

	"github.com/go-playground/validator/v10"
)

type VillageServiceImpl struct {
	VillageRepository  repository.VillageRepository
	DistrictRepository repository.DistrictRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewVillageService(
	villageRepository repository.VillageRepository, districtRepository repository.DistrictRepository, DB *sql.DB, validate *validator.Validate,
) VillageService {
	return &VillageServiceImpl{
		VillageRepository:  villageRepository,
		DistrictRepository: districtRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *VillageServiceImpl) Create(ctx context.Context, webRequest any) any {
	villageRequest := webRequest.(request.VillageCreateRequest)
	err := service.Validate.Struct(villageRequest)
	helper.PanicIfError(err)

	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	district := service.findDistrictById(ctx, tx, villageRequest.DistrictId)
	village := domain.Village{
		Name:       villageRequest.Name,
		District:   district,
		BaseDomain: helper.BuildBaseDomain(ctx),
	}
	village = service.VillageRepository.Save(ctx, tx, village).(domain.Village)
	return helper.ToVillageResponse(village)
}

func (service *VillageServiceImpl) Update(ctx context.Context, webRequest any) any {
	villageRequest := webRequest.(request.VillageUpdateRequest)
	err := service.Validate.Struct(villageRequest)
	helper.PanicIfError(err)

	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	district := service.findDistrictById(ctx, tx, villageRequest.DistrictId)
	village := service.findVillageById(ctx, tx, villageRequest.Id)
	village.Name = villageRequest.Name
	village.District = district
	helper.SetUpdatedBy(ctx, &village.BaseDomain)

	village = service.VillageRepository.Update(ctx, tx, village).(domain.Village)
	return helper.ToVillageResponse(village)
}

func (service *VillageServiceImpl) Delete(ctx context.Context, id int64) {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	village := service.findVillageById(ctx, tx, id)
	helper.SetUpdatedBy(ctx, &village.BaseDomain)
	village.IsDeleted = true
	service.VillageRepository.Delete(ctx, tx, village)
}

func (service *VillageServiceImpl) FindById(ctx context.Context, id int64) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	village := service.findVillageById(ctx, tx, id)
	return helper.ToVillageResponse(village)
}

func (service *VillageServiceImpl) FindAllWithPagination(ctx context.Context, searchBy any) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	villageSearch := searchBy.(search.VillageSearch)
	villages := service.VillageRepository.FindAll(ctx, tx, villageSearch).([]domain.Village)
	totalItem := service.VillageRepository.FindTotalItem(ctx, tx, villageSearch)
	return helper.ToResponsePagination(villageSearch.Page, villageSearch.Size, totalItem, helper.ToVillageResponses(villages))
}

func (service *VillageServiceImpl) FindAllWithoutPagination(ctx context.Context, searchBy any) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	villages := service.VillageRepository.FindAll(ctx, tx, searchBy).([]domain.Village)
	return helper.ToVillageResponses(villages)
}

func (service *VillageServiceImpl) beginTransaction() *sql.Tx {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	return tx
}

func (service *VillageServiceImpl) findVillageById(ctx context.Context, tx *sql.Tx, villageId int64) domain.Village {
	village, err := service.VillageRepository.FindById(ctx, tx, villageId)
	exception.PanicErrorBadRequest(err)
	return village.(domain.Village)
}

func (service *VillageServiceImpl) findDistrictById(ctx context.Context, tx *sql.Tx, districtId int64) domain.District {
	district, err := service.DistrictRepository.FindById(ctx, tx, districtId)
	exception.PanicErrorBadRequest(err)
	return district.(domain.District)
}
