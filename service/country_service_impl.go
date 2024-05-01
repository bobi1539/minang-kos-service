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

type CountryServiceImpl struct {
	CountryRepository repository.CountryRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewCountryService(countryRepository repository.CountryRepository, DB *sql.DB, validate *validator.Validate) CountryService {
	return &CountryServiceImpl{
		CountryRepository: countryRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *CountryServiceImpl) Create(ctx context.Context, webRequest any) any {
	countryRequest := webRequest.(request.CountryCreateRequest)
	err := service.Validate.Struct(countryRequest)
	helper.PanicIfError(err)

	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	country := domain.Country{
		Name:       countryRequest.Name,
		BaseDomain: helper.BuildBaseDomain(ctx),
	}

	country = service.CountryRepository.Save(ctx, tx, country).(domain.Country)

	return helper.ToCountryResponse(country)
}

func (service *CountryServiceImpl) Update(ctx context.Context, webRequest any) any {
	countryRequest := webRequest.(request.CountryUpdateRequest)
	err := service.Validate.Struct(countryRequest)
	helper.PanicIfError(err)

	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	country := service.findCountryById(ctx, tx, countryRequest.Id)
	country.Name = countryRequest.Name
	helper.SetUpdatedBy(ctx, &country.BaseDomain)

	country = service.CountryRepository.Update(ctx, tx, country).(domain.Country)

	return helper.ToCountryResponse(country)
}

func (service *CountryServiceImpl) Delete(ctx context.Context, id int64) {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	country := service.findCountryById(ctx, tx, id)
	helper.SetUpdatedBy(ctx, &country.BaseDomain)
	country.IsDeleted = true

	service.CountryRepository.Delete(ctx, tx, country)
}

func (service *CountryServiceImpl) FindById(ctx context.Context, id int64) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	country := service.findCountryById(ctx, tx, id)
	return helper.ToCountryResponse(country)
}

func (service *CountryServiceImpl) FindAllWithPagination(ctx context.Context, searchBy map[string]any) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	countries := service.CountryRepository.FindAllWithPagination(ctx, tx, searchBy).([]domain.Country)
	totalItem := service.CountryRepository.FindTotalItem(ctx, tx, searchBy)
	return helper.ToResponsePagination(searchBy["page"].(int), searchBy["size"].(int), totalItem, helper.ToCountryResponses(countries))
}

func (service *CountryServiceImpl) FindAllWithoutPagination(ctx context.Context, searchBy map[string]any) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	countries := service.CountryRepository.FindAllWithoutPagination(ctx, tx, searchBy).([]domain.Country)
	return helper.ToCountryResponses(countries)
}

func (service *CountryServiceImpl) beginTransaction() *sql.Tx {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	return tx
}

func (service *CountryServiceImpl) findCountryById(ctx context.Context, tx *sql.Tx, countryId int64) domain.Country {
	country, err := service.CountryRepository.FindById(ctx, tx, countryId)
	exception.PanicErrorBadRequest(err)
	return country.(domain.Country)
}
