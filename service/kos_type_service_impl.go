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

type KosTypeServiceImpl struct {
	KosTypeRepository repository.KosTypeRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewKosTypeService(
	kosTypeRepository repository.KosTypeRepository, DB *sql.DB, validate *validator.Validate,
) KosTypeService {
	return &KosTypeServiceImpl{
		KosTypeRepository: kosTypeRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *KosTypeServiceImpl) Create(ctx context.Context, webRequest any) any {
	kosTypeRequest := webRequest.(request.KosTypeCreateRequest)
	err := service.Validate.Struct(kosTypeRequest)
	helper.PanicIfError(err)

	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	kosType := domain.KosType{
		Name:       kosTypeRequest.Name,
		BaseDomain: helper.BuildBaseDomain(ctx),
	}

	kosType = service.KosTypeRepository.Save(ctx, tx, kosType).(domain.KosType)

	return helper.ToKosTypeResponse(kosType)
}

func (service *KosTypeServiceImpl) Update(ctx context.Context, webRequest any) any {
	kosTypeRequest := webRequest.(request.KosTypeUpdateRequest)
	err := service.Validate.Struct(kosTypeRequest)
	helper.PanicIfError(err)

	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	kosType := service.findKosTypeById(ctx, tx, kosTypeRequest.Id)
	kosType.Name = kosTypeRequest.Name
	helper.SetUpdatedBy(ctx, &kosType.BaseDomain)

	kosType = service.KosTypeRepository.Update(ctx, tx, kosType).(domain.KosType)

	return helper.ToKosTypeResponse(kosType)
}

func (service *KosTypeServiceImpl) Delete(ctx context.Context, id int64) {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	kosType := service.findKosTypeById(ctx, tx, id)
	helper.SetUpdatedBy(ctx, &kosType.BaseDomain)
	kosType.IsDeleted = true

	service.KosTypeRepository.Delete(ctx, tx, kosType)
}

func (service *KosTypeServiceImpl) FindById(ctx context.Context, id int64) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	kosType := service.findKosTypeById(ctx, tx, id)
	return helper.ToKosTypeResponse(kosType)
}

func (service *KosTypeServiceImpl) FindAllWithPagination(ctx context.Context, searchBy any) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	kosTypeSearch := searchBy.(search.KosTypeSearch)
	kosTypes := service.KosTypeRepository.FindAll(ctx, tx, kosTypeSearch).([]domain.KosType)
	totalItem := service.KosTypeRepository.FindTotalItem(ctx, tx, kosTypeSearch)
	return helper.ToResponsePagination(kosTypeSearch.Page, kosTypeSearch.Size, totalItem, helper.ToKosTypeResponses(kosTypes))
}

func (service *KosTypeServiceImpl) FindAllWithoutPagination(ctx context.Context, searchBy any) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	kosTypes := service.KosTypeRepository.FindAll(ctx, tx, searchBy).([]domain.KosType)
	return helper.ToKosTypeResponses(kosTypes)
}

func (service *KosTypeServiceImpl) beginTransaction() *sql.Tx {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	return tx
}

func (service *KosTypeServiceImpl) findKosTypeById(ctx context.Context, tx *sql.Tx, kosTypeId int64) domain.KosType {
	kosType, err := service.KosTypeRepository.FindById(ctx, tx, kosTypeId)
	exception.PanicErrorBadRequest(err)
	return kosType.(domain.KosType)
}
