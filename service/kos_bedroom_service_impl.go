package service

import (
	"context"
	"database/sql"
	"minang-kos-service/constant"
	"minang-kos-service/exception"
	"minang-kos-service/helper"
	"minang-kos-service/model/domain"
	"minang-kos-service/model/web/request"
	"minang-kos-service/repository"

	"github.com/go-playground/validator/v10"
)

type KosBedroomServiceImpl struct {
	KosBedroomRepository  repository.KosBedroomRepository
	KosTypeRepository     repository.KosTypeRepository
	VillageRepository     repository.VillageRepository
	UserRepository        repository.UserRepository
	KosFacilityRepository repository.KosFacilityRepository
	FacilityRepository    repository.FacilityRepository
	DB                    *sql.DB
	Validate              *validator.Validate
}

func NewKosBedroomService(
	kosBedroomRepository repository.KosBedroomRepository,
	kosTypeRepository repository.KosTypeRepository,
	villageRepository repository.VillageRepository,
	userRepository repository.UserRepository,
	kosFacilityRepository repository.KosFacilityRepository,
	facilityRepository repository.FacilityRepository,
	DB *sql.DB,
	validate *validator.Validate,
) KosBedroomService {
	return &KosBedroomServiceImpl{
		KosBedroomRepository:  kosBedroomRepository,
		KosTypeRepository:     kosTypeRepository,
		VillageRepository:     villageRepository,
		UserRepository:        userRepository,
		KosFacilityRepository: kosFacilityRepository,
		FacilityRepository:    facilityRepository,
		DB:                    DB,
		Validate:              validate,
	}
}

func (service *KosBedroomServiceImpl) Create(ctx context.Context, webRequest any) any {
	kosBedroomRequest := webRequest.(request.KosBedroomCreateRequest)
	err := service.Validate.Struct(kosBedroomRequest)
	helper.PanicIfError(err)

	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	kosType := service.findKosTypeById(ctx, tx, kosBedroomRequest.KosTypeId)
	village := service.findVillageById(ctx, tx, kosBedroomRequest.VillageId)
	baseDomain := helper.BuildBaseDomain(ctx)
	user := service.findUserById(ctx, tx, baseDomain.CreatedBy)

	kosBedroom := domain.KosBedroom{
		Title:                kosBedroomRequest.Title,
		RoomLength:           kosBedroomRequest.RoomLength,
		RoomWidth:            kosBedroomRequest.RoomWidth,
		UnitLength:           constant.METER,
		IsIncludeElectricity: kosBedroomRequest.IsIncludeElectricity,
		Price:                kosBedroomRequest.Price,
		Images:               "images",
		KosType:              kosType,
		Village:              village,
		User:                 user,
		BaseDomain:           baseDomain,
	}

	kosBedroom = service.KosBedroomRepository.Save(ctx, tx, kosBedroom).(domain.KosBedroom)
	service.createKosFacility(ctx, tx, kosBedroom, kosBedroomRequest.FacilityIds)
	return helper.ToKosBedroomResponse(kosBedroom)
}

func (service *KosBedroomServiceImpl) Update(ctx context.Context, webRequest any) any {
	panic("imp")
}

func (service *KosBedroomServiceImpl) Delete(ctx context.Context, id int64) {
	panic("imp")
}

func (service *KosBedroomServiceImpl) FindById(ctx context.Context, id int64) any {
	panic("imp")
}

func (service *KosBedroomServiceImpl) FindAllWithPagination(ctx context.Context, searchBy map[string]any) any {
	panic("imp")
}

func (service *KosBedroomServiceImpl) FindAllWithoutPagination(ctx context.Context, searchBy map[string]any) any {
	panic("imp")
}

func (service *KosBedroomServiceImpl) beginTransaction() *sql.Tx {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	return tx
}

func (service *KosBedroomServiceImpl) findKosTypeById(ctx context.Context, tx *sql.Tx, kosTypeId int64) domain.KosType {
	kosType, err := service.KosTypeRepository.FindById(ctx, tx, kosTypeId)
	exception.PanicErrorBadRequest(err)
	return kosType.(domain.KosType)
}

func (service *KosBedroomServiceImpl) findVillageById(ctx context.Context, tx *sql.Tx, villageId int64) domain.Village {
	village, err := service.VillageRepository.FindById(ctx, tx, villageId)
	exception.PanicErrorBadRequest(err)
	return village.(domain.Village)
}

func (service *KosBedroomServiceImpl) findUserById(ctx context.Context, tx *sql.Tx, userId int64) domain.User {
	user, err := service.UserRepository.FindById(ctx, tx, userId)
	exception.PanicErrorBadRequest(err)
	return user.(domain.User)
}

func (service *KosBedroomServiceImpl) findFacilityById(ctx context.Context, tx *sql.Tx, facilityId int64) domain.Facility {
	facility, err := service.FacilityRepository.FindById(ctx, tx, facilityId)
	exception.PanicErrorBadRequest(err)
	return facility.(domain.Facility)
}

func (service *KosBedroomServiceImpl) createKosFacility(
	ctx context.Context, tx *sql.Tx, kosBedroom domain.KosBedroom, facilityIds []int64,
) {
	for _, facilityId := range facilityIds {
		facility := service.findFacilityById(ctx, tx, facilityId)
		kosFacility := domain.KosFacility{
			KosBedroom: kosBedroom,
			Facility:   facility,
		}
		service.KosFacilityRepository.Save(ctx, tx, kosFacility)
	}
}
