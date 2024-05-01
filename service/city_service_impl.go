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

type CityServiceImpl struct {
	CityRepository     repository.CityRepository
	ProvinceRepository repository.ProvinceRepository
	DB                 *sql.DB
	Validate           validator.Validate
}

func NewCityService(
	cityRepository repository.CityRepository, provinceRepository repository.ProvinceRepository, DB *sql.DB, validate validator.Validate,
) CityService {
	return &CityServiceImpl{
		CityRepository:     cityRepository,
		ProvinceRepository: provinceRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CityServiceImpl) Create(ctx context.Context, webRequest any) any {
	cityRequest := webRequest.(request.CityCreateRequest)
	err := service.Validate.Struct(cityRequest)
	helper.PanicIfError(err)

	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	province := service.findProvinceById(ctx, tx, cityRequest.ProvinceId)
	city := domain.City{
		Name:       cityRequest.Name,
		Province:   province,
		BaseDomain: helper.BuildBaseDomain(ctx),
	}
	city = service.CityRepository.Save(ctx, tx, city).(domain.City)
	return helper.ToCityResponse(city)
}

func (service *CityServiceImpl) Update(ctx context.Context, webRequest any) any {
	cityRequest := webRequest.(request.CityUpdateRequest)
	err := service.Validate.Struct(cityRequest)
	helper.PanicIfError(err)

	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	province := service.findProvinceById(ctx, tx, cityRequest.ProvinceId)
	city := service.findCityById(ctx, tx, cityRequest.Id)
	city.Name = cityRequest.Name
	city.Province = province
	helper.SetUpdatedBy(ctx, &city.BaseDomain)

	city = service.CityRepository.Update(ctx, tx, city).(domain.City)
	return helper.ToCityResponse(city)
}

func (service *CityServiceImpl) Delete(ctx context.Context, id int64) {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	city := service.findCityById(ctx, tx, id)
	helper.SetUpdatedBy(ctx, &city.BaseDomain)
	city.IsDeleted = true
	service.CityRepository.Delete(ctx, tx, city)
}

func (service *CityServiceImpl) FindById(ctx context.Context, id int64) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	city := service.findCityById(ctx, tx, id)
	return helper.ToCityResponse(city)
}

func (service *CityServiceImpl) FindAllWithPagination(ctx context.Context, searchBy map[string]any) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	cities := service.CityRepository.FindAllWithPagination(ctx, tx, searchBy).([]domain.City)
	totalItem := service.CityRepository.FindTotalItem(ctx, tx, searchBy)
	return helper.ToResponsePagination(
		searchBy["page"].(int), searchBy["size"].(int), totalItem, helper.ToCityResponses(cities),
	)
}

func (service *CityServiceImpl) FindAllWithoutPagination(ctx context.Context, searchBy map[string]any) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	cities := service.CityRepository.FindAllWithoutPagination(ctx, tx, searchBy).([]domain.City)
	return helper.ToCityResponses(cities)
}

func (service *CityServiceImpl) beginTransaction() *sql.Tx {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	return tx
}

func (service *CityServiceImpl) findCityById(ctx context.Context, tx *sql.Tx, cityId int64) domain.City {
	city, err := service.CityRepository.FindById(ctx, tx, cityId)
	exception.PanicErrorBadRequest(err)
	return city.(domain.City)
}

func (service *CityServiceImpl) findProvinceById(ctx context.Context, tx *sql.Tx, provinceId int64) domain.Province {
	province, err := service.ProvinceRepository.FindById(ctx, tx, provinceId)
	exception.PanicErrorBadRequest(err)
	return province.(domain.Province)
}
