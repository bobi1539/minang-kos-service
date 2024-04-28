package service

import (
	"context"
	"database/sql"
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

	tx := beginTransaction(service)
	defer helper.CommitOrRollback(tx)

	country := domain.Country{
		Name: countryRequest.Name,
		BaseDomain: domain.BaseDomain{
			CreatedAt:     time.Now(),
			CreatedBy:     1,
			CreatedByName: "test",
			UpdatedAt:     time.Now(),
			UpdatedBy:     1,
			UpdatedByName: "test",
			IsDeleted:     false,
		},
	}

	country = service.CountryRepository.Save(ctx, tx, country).(domain.Country)

	return helper.ToCountryResponse(country)
}

func (service *CountryServiceImpl) Update(ctx context.Context, webRequest any) any {
	countryRequest := webRequest.(request.CountryUpdateRequest)
	err := service.Validate.Struct(countryRequest)
	helper.PanicIfError(err)

	tx := beginTransaction(service)
	defer helper.CommitOrRollback(tx)

	countryById, err := service.CountryRepository.FindById(ctx, tx, countryRequest.Id)
	helper.PanicIfError(err)

	country := countryById.(domain.Country)
	country.Name = countryRequest.Name
	country.UpdatedAt = time.Now()
	country.UpdatedBy = 10
	country.UpdatedByName = "ucup"

	country = service.CountryRepository.Update(ctx, tx, country).(domain.Country)

	return helper.ToCountryResponse(country)
}

func beginTransaction(service *CountryServiceImpl) *sql.Tx {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	return tx
}
