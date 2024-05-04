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

type FacilityTypeServiceImpl struct {
	FacilityTypeRepository repository.FacilityTypeRepository
	DB                     *sql.DB
	Validate               *validator.Validate
}

func NewFacilityTypeService(facilityTypeRepository repository.FacilityTypeRepository, DB *sql.DB, validate *validator.Validate) FacilityTypeService {
	return &FacilityTypeServiceImpl{
		FacilityTypeRepository: facilityTypeRepository,
		DB:                     DB,
		Validate:               validate,
	}
}

func (service *FacilityTypeServiceImpl) Create(ctx context.Context, webRequest any) any {
	facilityTypeRequest := webRequest.(request.FacilityTypeCreateRequest)
	err := service.Validate.Struct(facilityTypeRequest)
	helper.PanicIfError(err)

	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	facilityType := domain.FacilityType{
		Name:       facilityTypeRequest.Name,
		BaseDomain: helper.BuildBaseDomain(ctx),
	}

	facilityType = service.FacilityTypeRepository.Save(ctx, tx, facilityType).(domain.FacilityType)

	return helper.ToFacilityTypeResponse(facilityType)
}

func (service *FacilityTypeServiceImpl) Update(ctx context.Context, webRequest any) any {
	facilityTypeRequest := webRequest.(request.FacilityTypeUpdateRequest)
	err := service.Validate.Struct(facilityTypeRequest)
	helper.PanicIfError(err)

	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	facilityType := service.findFacilityTypeById(ctx, tx, facilityTypeRequest.Id)
	facilityType.Name = facilityTypeRequest.Name
	helper.SetUpdatedBy(ctx, &facilityType.BaseDomain)

	facilityType = service.FacilityTypeRepository.Update(ctx, tx, facilityType).(domain.FacilityType)

	return helper.ToFacilityTypeResponse(facilityType)
}

func (service *FacilityTypeServiceImpl) Delete(ctx context.Context, id int64) {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	facilityType := service.findFacilityTypeById(ctx, tx, id)
	helper.SetUpdatedBy(ctx, &facilityType.BaseDomain)
	facilityType.IsDeleted = true

	service.FacilityTypeRepository.Delete(ctx, tx, facilityType)
}

func (service *FacilityTypeServiceImpl) FindById(ctx context.Context, id int64) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	facilityType := service.findFacilityTypeById(ctx, tx, id)
	return helper.ToFacilityTypeResponse(facilityType)
}

func (service *FacilityTypeServiceImpl) FindAllWithPagination(ctx context.Context, searchBy any) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	facilityTypeSearch := searchBy.(search.FacilityTypeSearch)
	facilityTypes := service.FacilityTypeRepository.FindAll(ctx, tx, facilityTypeSearch).([]domain.FacilityType)
	totalItem := service.FacilityTypeRepository.FindTotalItem(ctx, tx, facilityTypeSearch)
	return helper.ToResponsePagination(facilityTypeSearch.Page, facilityTypeSearch.Size, totalItem, helper.ToFacilityTypeResponses(facilityTypes))
}

func (service *FacilityTypeServiceImpl) FindAllWithoutPagination(ctx context.Context, searchBy any) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	facilityTypes := service.FacilityTypeRepository.FindAll(ctx, tx, searchBy).([]domain.FacilityType)
	return helper.ToFacilityTypeResponses(facilityTypes)
}

func (service *FacilityTypeServiceImpl) beginTransaction() *sql.Tx {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	return tx
}

func (service *FacilityTypeServiceImpl) findFacilityTypeById(ctx context.Context, tx *sql.Tx, facilityTypeId int64) domain.FacilityType {
	facilityType, err := service.FacilityTypeRepository.FindById(ctx, tx, facilityTypeId)
	exception.PanicErrorBadRequest(err)
	return facilityType.(domain.FacilityType)
}
