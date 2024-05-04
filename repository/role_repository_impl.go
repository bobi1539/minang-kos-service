package repository

import (
	"context"
	"database/sql"
	"errors"
	"minang-kos-service/constant"
	"minang-kos-service/helper"
	"minang-kos-service/model/domain"
	"minang-kos-service/model/web/search"
)

type RoleRepositoryImpl struct {
}

func NewRoleRepository() RoleRepository {
	return &RoleRepositoryImpl{}
}

func (repository *RoleRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, domainModel any) any {
	role := domainModel.(domain.Role)
	result, err := tx.ExecContext(
		ctx, sqlSaveRole(), role.Name, role.CreatedAt, role.CreatedBy, role.CreatedByName,
		role.UpdatedAt, role.UpdatedBy, role.UpdatedByName, role.IsDeleted,
	)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	role.Id = id
	return role
}

func (repository *RoleRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, domainModel any) any {
	role := domainModel.(domain.Role)
	_, err := tx.ExecContext(ctx, sqlUpdateRole(), role.Name, role.UpdatedAt, role.UpdatedBy, role.UpdatedByName, role.Id)
	helper.PanicIfError(err)

	return role
}

func (repository *RoleRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, domainModel any) {
	role := domainModel.(domain.Role)
	_, err := tx.ExecContext(
		ctx, sqlDeleteRole(), role.UpdatedAt, role.UpdatedBy, role.UpdatedByName, role.IsDeleted, role.Id,
	)
	helper.PanicIfError(err)
}

func (repository *RoleRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int64) (any, error) {
	sqlQuery := sqlSelectRole() + " AND id = ?"
	rows, err := tx.QueryContext(ctx, sqlQuery, id)
	helper.PanicIfError(err)
	defer rows.Close()

	role := domain.Role{}
	if rows.Next() {
		scanRole(rows, &role)
		return role, nil
	}
	return role, errors.New(constant.DATA_NOT_FOUND)
}

func (repository *RoleRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, searchBy any) any {
	roleSearch := searchBy.(search.RoleSearch)

	sqlSearch, args := sqlSearchRoleBy(roleSearch.Name)
	sqlQuery := sqlSelectRole() + sqlSearch + ""

	if roleSearch.Page > 0 {
		sqlQuery += " LIMIT ? OFFSET ?"
		offset := helper.GetSqlOffset(roleSearch.Page, roleSearch.Size)
		args = append(args, roleSearch.Size, offset)
	}

	rows := FetchRows(ctx, tx, sqlQuery, args)
	defer rows.Close()

	return getRoles(rows)
}

func (repository *RoleRepositoryImpl) FindTotalItem(ctx context.Context, tx *sql.Tx, searchBy any) int {
	roleSearch := searchBy.(search.RoleSearch)

	sqlSearch, args := sqlSearchRoleBy(roleSearch.Name)
	sqlQuery := sqlFindTotalRole() + sqlSearch

	rows := FetchRows(ctx, tx, sqlQuery, args)
	defer rows.Close()

	return ScanTotalItem(rows)
}

func sqlSaveRole() string {
	return "INSERT INTO m_role(name," +
		" created_at," +
		" created_by," +
		" created_by_name," +
		" updated_at," +
		" updated_by," +
		" updated_by_name," +
		" is_deleted) VALUES (?,?,?,?,?,?,?,?)"
}

func sqlUpdateRole() string {
	return "UPDATE m_role SET name = ?," +
		" updated_at = ?," +
		" updated_by = ?," +
		" updated_by_name = ?" +
		" WHERE id = ?"
}

func sqlDeleteRole() string {
	return "UPDATE m_role SET updated_at = ?," +
		" updated_by = ?," +
		" updated_by_name = ?," +
		" is_deleted = ?" +
		" WHERE id = ?"
}

func sqlSelectRole() string {
	return "SELECT id," +
		" name," +
		" created_at," +
		" created_by," +
		" created_by_name," +
		" updated_at," +
		" updated_by," +
		" updated_by_name," +
		" is_deleted FROM m_role" +
		" WHERE is_deleted = false"
}

func sqlFindTotalRole() string {
	return "SELECT COUNT(1) AS totalItem FROM m_role WHERE is_deleted = false"
}

func sqlSearchRoleBy(name string) (string, []any) {
	var args []any
	sqlQuery := ""

	if len(name) != 0 {
		sqlQuery += " AND LOWER(name) LIKE ?"
		args = append(args, helper.StringQueryLike(name))
	}

	sqlQuery += " ORDER BY id ASC"
	return sqlQuery, args
}

func getRoles(rows *sql.Rows) []domain.Role {
	var roles []domain.Role
	for rows.Next() {
		role := domain.Role{}
		scanRole(rows, &role)
		roles = append(roles, role)
	}
	return roles
}

func scanRole(rows *sql.Rows, role *domain.Role) {
	err := rows.Scan(
		&role.Id,
		&role.Name,
		&role.CreatedAt,
		&role.CreatedBy,
		&role.CreatedByName,
		&role.UpdatedAt,
		&role.UpdatedBy,
		&role.UpdatedByName,
		&role.IsDeleted,
	)
	helper.PanicIfError(err)
}
