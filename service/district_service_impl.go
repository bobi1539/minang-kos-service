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

type DistrictServiceImpl struct {
	DistrictRepository repository.DistrictRepository
	CityRepository     repository.CityRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewDistrictService(
	districtRepository repository.DistrictRepository, cityRepository repository.CityRepository, DB *sql.DB, validate *validator.Validate,
) DistrictService {
	return &DistrictServiceImpl{
		DistrictRepository: districtRepository,
		CityRepository:     cityRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *DistrictServiceImpl) Create(ctx context.Context, webRequest any) any {
	districtRequest := webRequest.(request.DistrictCreateRequest)
	err := service.Validate.Struct(districtRequest)
	helper.PanicIfError(err)

	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	city := service.findCityById(ctx, tx, districtRequest.CityId)
	district := domain.District{
		Name:       districtRequest.Name,
		City:       city,
		BaseDomain: helper.BuildBaseDomain(ctx),
	}
	district = service.DistrictRepository.Save(ctx, tx, district).(domain.District)
	return helper.ToDistrictResponse(district)
}

func (service *DistrictServiceImpl) Update(ctx context.Context, webRequest any) any {
	districtRequest := webRequest.(request.DistrictUpdateRequest)
	err := service.Validate.Struct(districtRequest)
	helper.PanicIfError(err)

	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	city := service.findCityById(ctx, tx, districtRequest.CityId)
	district := service.findDistrictById(ctx, tx, districtRequest.Id)
	district.Name = districtRequest.Name
	district.City = city
	helper.SetUpdatedBy(ctx, &district.BaseDomain)

	district = service.DistrictRepository.Update(ctx, tx, district).(domain.District)
	return helper.ToDistrictResponse(district)
}

func (service *DistrictServiceImpl) Delete(ctx context.Context, id int64) {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	district := service.findDistrictById(ctx, tx, id)
	helper.SetUpdatedBy(ctx, &district.BaseDomain)
	district.IsDeleted = true
	service.DistrictRepository.Delete(ctx, tx, district)
}

func (service *DistrictServiceImpl) FindById(ctx context.Context, id int64) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	district := service.findDistrictById(ctx, tx, id)
	return helper.ToDistrictResponse(district)
}

func (service *DistrictServiceImpl) FindAllWithPagination(ctx context.Context, searchBy any) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	districtSearch := searchBy.(search.DistrictSearch)
	districts := service.DistrictRepository.FindAll(ctx, tx, districtSearch).([]domain.District)
	totalItem := service.DistrictRepository.FindTotalItem(ctx, tx, districtSearch)
	return helper.ToResponsePagination(districtSearch.Page, districtSearch.Size, totalItem, helper.ToDistrictResponses(districts))
}

func (service *DistrictServiceImpl) FindAllWithoutPagination(ctx context.Context, searchBy any) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	districts := service.DistrictRepository.FindAll(ctx, tx, searchBy).([]domain.District)
	return helper.ToDistrictResponses(districts)
}

func (service *DistrictServiceImpl) beginTransaction() *sql.Tx {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	return tx
}

func (service *DistrictServiceImpl) findDistrictById(ctx context.Context, tx *sql.Tx, districtId int64) domain.District {
	district, err := service.DistrictRepository.FindById(ctx, tx, districtId)
	exception.PanicErrorBadRequest(err)
	return district.(domain.District)
}

func (service *DistrictServiceImpl) findCityById(ctx context.Context, tx *sql.Tx, cityId int64) domain.City {
	city, err := service.CityRepository.FindById(ctx, tx, cityId)
	exception.PanicErrorBadRequest(err)
	return city.(domain.City)
}
