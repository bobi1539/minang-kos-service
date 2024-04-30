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
		BaseDomain: helper.BuildBaseDomain(),
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
	role.UpdatedAt = time.Now()
	role.UpdatedBy = 10
	role.UpdatedByName = "ucup"

	role = service.RoleRepository.Update(ctx, tx, role).(domain.Role)

	return helper.ToRoleResponse(role)
}

func (service *RoleServiceImpl) Delete(ctx context.Context, id int64) {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	role := service.findRoleById(ctx, tx, id)
	role.UpdatedAt = time.Now()
	role.UpdatedBy = 12
	role.UpdatedByName = "tono"
	role.IsDeleted = true

	service.RoleRepository.Delete(ctx, tx, role)
}

func (service *RoleServiceImpl) FindById(ctx context.Context, id int64) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	role := service.findRoleById(ctx, tx, id)
	return helper.ToRoleResponse(role)
}

func (service *RoleServiceImpl) FindAllWithPagination(ctx context.Context, searchBy map[string]any) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	roles := service.RoleRepository.FindAllWithPagination(ctx, tx, searchBy).([]domain.Role)
	totalItem := service.RoleRepository.FindTotalItem(ctx, tx, searchBy)
	return helper.ToResponsePagination(searchBy["page"].(int), searchBy["size"].(int), totalItem, helper.ToRoleResponses(roles))
}

func (service *RoleServiceImpl) FindAllWithoutPagination(ctx context.Context, searchBy map[string]any) any {
	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	roles := service.RoleRepository.FindAllWithoutPagination(ctx, tx, searchBy).([]domain.Role)
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
