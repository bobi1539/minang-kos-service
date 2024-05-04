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

type RoleServiceImpl struct {
	RoleRepository repository.RoleRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewRoleService(roleRepository repository.RoleRepository, DB *sql.DB, validate *validator.Validate) RoleService {
	return &RoleServiceImpl{
		RoleRepository: roleRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *RoleServiceImpl) Create(ctx context.Context, webRequest any) any {
	roleRequest := webRequest.(request.RoleCreateRequest)
	err := service.Validate.Struct(roleRequest)
	helper.PanicIfError(err)

	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	role := domain.Role{
		Name:       roleRequest.Name,
		BaseDomain: helper.BuildBaseDomain(ctx),
	}

	role = service.RoleRepository.Save(ctx, tx, role).(domain.Role)
	return helper.ToRoleResponse(role)
}

func (service *RoleServiceImpl) Update(ctx context.Context, webRequest any) any {
	roleRequest := webRequest.(request.RoleUpdateRequest)
	err := service.Validate.Struct(roleRequest)
	helper.PanicIfError(err)

	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	role := service.findRoleById(ctx, tx, roleRequest.Id)
	role.Name = roleRequest.Name
	helper.SetUpdatedBy(ctx, &role.BaseDomain)

	role = service.RoleRepository.Update(ctx, tx, role).(domain.Role)

	return helper.ToRoleResponse(role)
}

func (service *RoleServiceImpl) Delete(ctx context.Context, id int64) {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	role := service.findRoleById(ctx, tx, id)
	helper.SetUpdatedBy(ctx, &role.BaseDomain)
	role.IsDeleted = true

	service.RoleRepository.Delete(ctx, tx, role)
}

func (service *RoleServiceImpl) FindById(ctx context.Context, id int64) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	role := service.findRoleById(ctx, tx, id)
	return helper.ToRoleResponse(role)
}

func (service *RoleServiceImpl) FindAllWithPagination(ctx context.Context, searchBy any) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	roleSearch := searchBy.(search.RoleSearch)
	roles := service.RoleRepository.FindAll(ctx, tx, roleSearch).([]domain.Role)
	totalItem := service.RoleRepository.FindTotalItem(ctx, tx, roleSearch)
	return helper.ToResponsePagination(roleSearch.Page, roleSearch.Size, totalItem, helper.ToRoleResponses(roles))
}

func (service *RoleServiceImpl) FindAllWithoutPagination(ctx context.Context, searchBy any) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	roles := service.RoleRepository.FindAll(ctx, tx, searchBy).([]domain.Role)
	return helper.ToRoleResponses(roles)
}

func (service *RoleServiceImpl) beginTransaction() *sql.Tx {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	return tx
}

func (service *RoleServiceImpl) findRoleById(ctx context.Context, tx *sql.Tx, roleId int64) domain.Role {
	role, err := service.RoleRepository.FindById(ctx, tx, roleId)
	exception.PanicErrorBadRequest(err)
	return role.(domain.Role)
}
