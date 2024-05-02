package service

import (
	"context"
	"database/sql"
	"minang-kos-service/exception"
	"minang-kos-service/helper"
	"minang-kos-service/model/domain"
	"minang-kos-service/model/web/request"
	"minang-kos-service/repository"

	"github.com/go-playground/validator/v10"
)

type FacilityServiceImpl struct {
	FacilityRepository     repository.FacilityRepository
	FacilityTypeRepository repository.FacilityTypeRepository
	DB                     *sql.DB
	Validate               *validator.Validate
}

func NewFacilityService(
	facilityRepository repository.FacilityRepository, facilityTypeRepository repository.FacilityTypeRepository, DB *sql.DB, validate *validator.Validate,
) FacilityService {
	return &FacilityServiceImpl{
		FacilityRepository:     facilityRepository,
		FacilityTypeRepository: facilityTypeRepository,
		DB:                     DB,
		Validate:               validate,
	}
}

func (service *FacilityServiceImpl) Create(ctx context.Context, webRequest any) any {
	facilityRequest := webRequest.(request.FacilityCreateRequest)
	err := service.Validate.Struct(facilityRequest)
	helper.PanicIfError(err)

	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	facilityType := service.findFacilityTypeById(ctx, tx, facilityRequest.FacilityTypeId)
	facility := domain.Facility{
		Name:         facilityRequest.Name,
		FacilityType: facilityType,
		BaseDomain:   helper.BuildBaseDomain(ctx),
	}
	facility = service.FacilityRepository.Save(ctx, tx, facility).(domain.Facility)
	return helper.ToFacilityResponse(facility)
}

func (service *FacilityServiceImpl) Update(ctx context.Context, webRequest any) any {
	facilityRequest := webRequest.(request.FacilityUpdateRequest)
	err := service.Validate.Struct(facilityRequest)
	helper.PanicIfError(err)

	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	facilityType := service.findFacilityTypeById(ctx, tx, facilityRequest.FacilityTypeId)
	facility := service.findFacilityById(ctx, tx, facilityRequest.Id)
	facility.Name = facilityRequest.Name
	facility.FacilityType = facilityType
	helper.SetUpdatedBy(ctx, &facility.BaseDomain)

	facility = service.FacilityRepository.Update(ctx, tx, facility).(domain.Facility)
	return helper.ToFacilityResponse(facility)
}

func (service *FacilityServiceImpl) Delete(ctx context.Context, id int64) {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	facility := service.findFacilityById(ctx, tx, id)
	helper.SetUpdatedBy(ctx, &facility.BaseDomain)
	facility.IsDeleted = true
	service.FacilityRepository.Delete(ctx, tx, facility)
}

func (service *FacilityServiceImpl) FindById(ctx context.Context, id int64) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	facility := service.findFacilityById(ctx, tx, id)
	return helper.ToFacilityResponse(facility)
}

func (service *FacilityServiceImpl) FindAllWithPagination(ctx context.Context, searchBy map[string]any) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	facilities := service.FacilityRepository.FindAllWithPagination(ctx, tx, searchBy).([]domain.Facility)
	totalItem := service.FacilityRepository.FindTotalItem(ctx, tx, searchBy)
	return helper.ToResponsePagination(
		searchBy["page"].(int), searchBy["size"].(int), totalItem, helper.ToFacilityResponses(facilities),
	)
}

func (service *FacilityServiceImpl) FindAllWithoutPagination(ctx context.Context, searchBy map[string]any) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	facilities := service.FacilityRepository.FindAllWithoutPagination(ctx, tx, searchBy).([]domain.Facility)
	return helper.ToFacilityResponses(facilities)
}

func (service *FacilityServiceImpl) beginTransaction() *sql.Tx {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	return tx
}

func (service *FacilityServiceImpl) findFacilityById(ctx context.Context, tx *sql.Tx, facilityId int64) domain.Facility {
	facility, err := service.FacilityRepository.FindById(ctx, tx, facilityId)
	exception.PanicErrorBadRequest(err)
	return facility.(domain.Facility)
}

func (service *FacilityServiceImpl) findFacilityTypeById(ctx context.Context, tx *sql.Tx, facilityTypeId int64) domain.FacilityType {
	facilityType, err := service.FacilityTypeRepository.FindById(ctx, tx, facilityTypeId)
	exception.PanicErrorBadRequest(err)
	return facilityType.(domain.FacilityType)
}
