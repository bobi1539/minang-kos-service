package service

import (
	"context"
	"database/sql"
	"minang-kos-service/exception"
	"minang-kos-service/helper"
	"minang-kos-service/model/domain"
	"minang-kos-service/model/web/request"
	"minang-kos-service/repository"
	"time"

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
		BaseDomain: helper.BuildBaseDomain(),
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

	countryById, err := service.CountryRepository.FindById(ctx, tx, countryRequest.Id)
	exception.PanicErrorBadRequest(err)

	country := countryById.(domain.Country)
	country.Name = countryRequest.Name
	country.UpdatedAt = time.Now()
	country.UpdatedBy = 10
	country.UpdatedByName = "ucup"

	country = service.CountryRepository.Update(ctx, tx, country).(domain.Country)

	return helper.ToCountryResponse(country)
}

func (service *CountryServiceImpl) Delete(ctx context.Context, id int64) {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	countryById, err := service.CountryRepository.FindById(ctx, tx, id)
	exception.PanicErrorBadRequest(err)

	country := countryById.(domain.Country)
	country.UpdatedAt = time.Now()
	country.UpdatedBy = 12
	country.UpdatedByName = "tono"
	country.IsDeleted = true

	service.CountryRepository.Delete(ctx, tx, country)
}

func (service *CountryServiceImpl) FindById(ctx context.Context, id int64) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	countryById, err := service.CountryRepository.FindById(ctx, tx, id)
	exception.PanicErrorBadRequest(err)

	country := countryById.(domain.Country)
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
