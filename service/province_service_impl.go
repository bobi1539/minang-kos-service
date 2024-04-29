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

	countryById, err := service.CountryRepository.FindById(ctx, tx, provinceRequest.CountryId)
	exception.PanicErrorBadRequest(err)

	country := countryById.(domain.Country)
	province := domain.Province{
		Name:       provinceRequest.Name,
		Country:    country,
		BaseDomain: helper.BuildBaseDomain(),
	}
	province = service.ProvinceRepository.Save(ctx, tx, province).(domain.Province)
	return helper.ToProvinceResponse(province)
}

func (service *ProvinceServiceImpl) Update(ctx context.Context, webRequest any) any {
	panic("imp")
}

func (service *ProvinceServiceImpl) Delete(ctx context.Context, id int64) {
	panic("imp")
}

func (service *ProvinceServiceImpl) FindById(ctx context.Context, id int64) any {
	panic("imp")
}

func (service *ProvinceServiceImpl) FindAllWithPagination(ctx context.Context, searchBy map[string]any) any {
	panic("imp")
}

func (service *ProvinceServiceImpl) FindAllWithoutPagination(ctx context.Context, searchBy map[string]any) any {
	panic("imp")
}

func (service *ProvinceServiceImpl) beginTransaction() *sql.Tx {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	return tx
}
