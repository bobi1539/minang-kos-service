package service

import (
	"context"
	"database/sql"
	"minang-kos-service/helper"
	"minang-kos-service/model/domain"
	"minang-kos-service/model/web/request"
	"minang-kos-service/repository"
	"time"
)

type CountryServiceImpl struct {
	CountryRepository repository.CountryRepository
	DB                *sql.DB
}

func NewCountryService(countryRepository repository.CountryRepository, DB *sql.DB) CountryService {
	return &CountryServiceImpl{
		CountryRepository: countryRepository,
		DB:                DB,
	}
}

func (service *CountryServiceImpl) Create(ctx context.Context, webRequest any) any {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	countryRequest := webRequest.(request.CountryCreateRequest)
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
