package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"minang-kos-service/constant"
	"minang-kos-service/exception"
	"minang-kos-service/helper"
	"minang-kos-service/model/domain"
	"minang-kos-service/model/web/request"
	"minang-kos-service/model/web/response"
	"minang-kos-service/model/web/search"
	"minang-kos-service/repository"
	"os"
	"strings"

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

	completeAddress := getCompleteAddress(kosBedroomRequest.Street, village)
	kosBedroom := domain.KosBedroom{
		Title:                kosBedroomRequest.Title,
		RoomLength:           kosBedroomRequest.RoomLength,
		RoomWidth:            kosBedroomRequest.RoomWidth,
		UnitLength:           constant.METER,
		IsIncludeElectricity: kosBedroomRequest.IsIncludeElectricity,
		Price:                kosBedroomRequest.Price,
		Street:               kosBedroomRequest.Street,
		CompleteAddress:      completeAddress,
		KosType:              kosType,
		Village:              village,
		User:                 user,
		BaseDomain:           baseDomain,
	}

	imagesFilename := service.saveImages(kosBedroomRequest.Images)
	kosBedroom.Images = imagesFilename

	kosBedroom = service.KosBedroomRepository.Save(ctx, tx, kosBedroom).(domain.KosBedroom)
	facilities, facilityTypes := service.createKosFacility(ctx, tx, kosBedroom, kosBedroomRequest.FacilityIds)
	return helper.ToKosBedroomResponse(kosBedroom, facilityTypes, facilities)
}

func (service *KosBedroomServiceImpl) Update(ctx context.Context, webRequest any) any {
	panic("imp")
}

func (service *KosBedroomServiceImpl) Delete(ctx context.Context, id int64) {
	panic("imp")
}

func (service *KosBedroomServiceImpl) FindById(ctx context.Context, id int64) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	kosBedroom := service.findKosBedroomById(ctx, tx, id)
	facilities, facilityTypes := service.findKosFacilities(ctx, tx, id, search.BuildPageSize(0, 0))
	return helper.ToKosBedroomResponse(kosBedroom, facilityTypes, facilities)
}

func (service *KosBedroomServiceImpl) FindAllWithPagination(ctx context.Context, searchBy any) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	kosBedroomSearch := searchBy.(search.KosBedroomSearch)
	kosBedrooms := service.KosBedroomRepository.FindAll(ctx, tx, kosBedroomSearch).([]domain.KosBedroom)
	totalItem := service.KosBedroomRepository.FindTotalItem(ctx, tx, kosBedroomSearch)

	var responses []response.KosBedroomResponse
	for _, kosBedroom := range kosBedrooms {
		facilities, facilityType := service.findKosFacilities(ctx, tx, kosBedroom.Id, search.BuildPageSize(1, 5))
		response := helper.ToKosBedroomResponse(kosBedroom, facilityType, facilities)
		responses = append(responses, response)
	}

	return helper.ToResponsePagination(kosBedroomSearch.Page, kosBedroomSearch.Size, totalItem, responses)
}

func (service *KosBedroomServiceImpl) FindAllWithoutPagination(ctx context.Context, searchBy any) any {
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

func (service *KosBedroomServiceImpl) findKosBedroomById(ctx context.Context, tx *sql.Tx, kosBedroomId int64) domain.KosBedroom {
	kosBedroom, err := service.KosBedroomRepository.FindById(ctx, tx, kosBedroomId)
	exception.PanicErrorBadRequest(err)
	return kosBedroom.(domain.KosBedroom)
}

func (service *KosBedroomServiceImpl) findKosFacilities(
	ctx context.Context, tx *sql.Tx, kosBedroomId int64, pageSize search.PageSize,
) ([]domain.Facility, []domain.FacilityType) {
	searchBy := search.KosFacilitySearch{
		KosBedroomId: kosBedroomId,
		FacilityId:   0,
		PageSize:     pageSize,
	}

	var facilities []domain.Facility
	var facilityTypes []domain.FacilityType

	kosFacilities := service.KosFacilityRepository.FindAll(ctx, tx, searchBy)
	for _, kosFacility := range kosFacilities {
		facilities = append(facilities, kosFacility.Facility)

		if !service.isFacilityTypeExist(kosFacility.Facility.FacilityType, facilityTypes) {
			facilityTypes = append(facilityTypes, kosFacility.Facility.FacilityType)
		}
	}

	return facilities, facilityTypes
}

func (service *KosBedroomServiceImpl) createKosFacility(
	ctx context.Context, tx *sql.Tx, kosBedroom domain.KosBedroom, facilityIds []int64,
) ([]domain.Facility, []domain.FacilityType) {
	var facilities []domain.Facility
	var facilityTypes []domain.FacilityType

	for _, facilityId := range facilityIds {
		facility := service.findFacilityById(ctx, tx, facilityId)
		kosFacility := domain.KosFacility{
			KosBedroom: kosBedroom,
			Facility:   facility,
		}
		service.KosFacilityRepository.Save(ctx, tx, kosFacility)
		facilities = append(facilities, facility)

		if !service.isFacilityTypeExist(facility.FacilityType, facilityTypes) {
			facilityTypes = append(facilityTypes, facility.FacilityType)
		}
	}
	return facilities, facilityTypes
}

func (service *KosBedroomServiceImpl) isFacilityTypeExist(facilityType domain.FacilityType, facilityTypes []domain.FacilityType) bool {
	for _, ft := range facilityTypes {
		if ft.Id == facilityType.Id {
			return true
		}
	}
	return false
}

func (service *KosBedroomServiceImpl) saveImages(images []string) string {
	var imagesFilename string
	for i, image := range images {
		arr := strings.Split(image, ";")

		filename := arr[0]
		service.validateImageExt(filename)

		imgByte := helper.DecodeBase64(arr[1])
		service.validateImageSize(len(imgByte))

		filepath := os.Getenv(constant.PATH_FILE_IMAGE_KOS_BEDROOM)
		filename = helper.GenerateImageFilename()
		fullpath := filepath + filename
		helper.WriteFile(fullpath, imgByte)

		imagesFilename += filename
		if i != (len(images) - 1) {
			imagesFilename += constant.COMMA
		}
	}
	return imagesFilename
}

func (service *KosBedroomServiceImpl) validateImageExt(filename string) {
	ext := helper.GetFileExtension(filename)
	allow := helper.AllowedExtension()
	isExtAllowed := allow[ext]
	if !isExtAllowed {
		exception.PanicErrorBadRequest(errors.New(constant.FILE_NOT_IMAGE))
	}
}

func (service *KosBedroomServiceImpl) validateImageSize(imgSize int) {
	if imgSize > constant.MAX_IMAGE_SIZE {
		exception.PanicErrorBadRequest(errors.New(constant.MAX_FILE_SIZE))
	}
}

func getCompleteAddress(street string, village domain.Village) string {
	return fmt.Sprintf(
		"%s,%s,%s,%s,%s,%s",
		street,
		village.Name,
		village.District.Name,
		village.District.City.Name,
		village.District.City.Province.Name,
		village.District.City.Province.Country.Name,
	)
}
