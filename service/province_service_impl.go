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

type ProvinceServiceImpl struct {
	ProvinceRepository repository.ProvinceRepository
	CountryRepository  repository.CountryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewProvinceService(
	provinceRepository repository.ProvinceRepository, countryRepository repository.CountryRepository, DB *sql.DB, validate *validator.Validate,
) ProvinceService {
	return &ProvinceServiceImpl{
		ProvinceRepository: provinceRepository,
		CountryRepository:  countryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *ProvinceServiceImpl) Create(ctx context.Context, webRequest any) any {
	provinceRequest := webRequest.(request.ProvinceCreateRequest)
	err := service.Validate.Struct(provinceRequest)
	helper.PanicIfError(err)

	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	country := service.findCountryById(ctx, tx, provinceRequest.CountryId)
	province := domain.Province{
		Name:       provinceRequest.Name,
		Country:    country,
		BaseDomain: helper.BuildBaseDomain(ctx),
	}
	province = service.ProvinceRepository.Save(ctx, tx, province).(domain.Province)
	return helper.ToProvinceResponse(province)
}

func (service *ProvinceServiceImpl) Update(ctx context.Context, webRequest any) any {
	provinceRequest := webRequest.(request.ProvinceUpdateRequest)
	err := service.Validate.Struct(provinceRequest)
	helper.PanicIfError(err)

	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	country := service.findCountryById(ctx, tx, provinceRequest.CountryId)
	province := service.findProvinceById(ctx, tx, provinceRequest.Id)
	province.Name = provinceRequest.Name
	province.Country = country
	helper.SetUpdatedBy(ctx, &province.BaseDomain)

	province = service.ProvinceRepository.Update(ctx, tx, province).(domain.Province)
	return helper.ToProvinceResponse(province)
}

func (service *ProvinceServiceImpl) Delete(ctx context.Context, id int64) {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	province := service.findProvinceById(ctx, tx, id)
	helper.SetUpdatedBy(ctx, &province.BaseDomain)
	province.IsDeleted = true
	service.ProvinceRepository.Delete(ctx, tx, province)
}

func (service *ProvinceServiceImpl) FindById(ctx context.Context, id int64) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	province := service.findProvinceById(ctx, tx, id)
	return helper.ToProvinceResponse(province)
}

func (service *ProvinceServiceImpl) FindAllWithPagination(ctx context.Context, searchBy map[string]any) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	provinces := service.ProvinceRepository.FindAllWithPagination(ctx, tx, searchBy).([]domain.Province)
	totalItem := service.ProvinceRepository.FindTotalItem(ctx, tx, searchBy)
	return helper.ToResponsePagination(
		searchBy["page"].(int), searchBy["size"].(int), totalItem, helper.ToProvinceResponses(provinces),
	)
}

func (service *ProvinceServiceImpl) FindAllWithoutPagination(ctx context.Context, searchBy map[string]any) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	provinces := service.ProvinceRepository.FindAllWithoutPagination(ctx, tx, searchBy).([]domain.Province)
	return helper.ToProvinceResponses(provinces)
}

func (service *ProvinceServiceImpl) beginTransaction() *sql.Tx {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	return tx
}

func (service *ProvinceServiceImpl) findCountryById(ctx context.Context, tx *sql.Tx, countryId int64) domain.Country {
	country, err := service.CountryRepository.FindById(ctx, tx, countryId)
	exception.PanicErrorBadRequest(err)
	return country.(domain.Country)
}

func (service *ProvinceServiceImpl) findProvinceById(ctx context.Context, tx *sql.Tx, provinceId int64) domain.Province {
	province, err := service.ProvinceRepository.FindById(ctx, tx, provinceId)
	exception.PanicErrorBadRequest(err)
	return province.(domain.Province)
}
