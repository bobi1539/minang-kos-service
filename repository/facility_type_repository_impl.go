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

type FacilityTypeRepositoryImpl struct {
}

func NewFacilityTypeRepository() FacilityTypeRepository {
	return &FacilityTypeRepositoryImpl{}
}

func (repository *FacilityTypeRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, domainModel any) any {
	facilityType := domainModel.(domain.FacilityType)
	result, err := tx.ExecContext(
		ctx, sqlSaveFacilityType(), facilityType.Name, facilityType.CreatedAt, facilityType.CreatedBy, facilityType.CreatedByName,
		facilityType.UpdatedAt, facilityType.UpdatedBy, facilityType.UpdatedByName, facilityType.IsDeleted,
	)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	facilityType.Id = id
	return facilityType
}

func (repository *FacilityTypeRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, domainModel any) any {
	facilityType := domainModel.(domain.FacilityType)
	_, err := tx.ExecContext(ctx, sqlUpdateFacilityType(), facilityType.Name, facilityType.UpdatedAt, facilityType.UpdatedBy, facilityType.UpdatedByName, facilityType.Id)
	helper.PanicIfError(err)

	return facilityType
}

func (repository *FacilityTypeRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, domainModel any) {
	facilityType := domainModel.(domain.FacilityType)
	_, err := tx.ExecContext(
		ctx, sqlDeleteFacilityType(), facilityType.UpdatedAt, facilityType.UpdatedBy, facilityType.UpdatedByName, facilityType.IsDeleted, facilityType.Id,
	)
	helper.PanicIfError(err)
}

func (repository *FacilityTypeRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int64) (any, error) {
	sqlQuery := sqlSelectFacilityType() + " AND id = ?"
	rows, err := tx.QueryContext(ctx, sqlQuery, id)
	helper.PanicIfError(err)
	defer rows.Close()

	facilityType := domain.FacilityType{}

	if rows.Next() {
		scanFacilityType(rows, &facilityType)
		return facilityType, nil
	}
	return facilityType, errors.New(constant.DATA_NOT_FOUND)
}

func (repository *FacilityTypeRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, searchBy any) any {
	facilityTypeSearch := searchBy.(search.FacilityTypeSearch)

	sqlSearch, args := sqlSearchFacilityTypeBy(facilityTypeSearch.Name)
	sqlQuery := sqlSelectFacilityType() + sqlSearch

	if facilityTypeSearch.Page > 0 {
		sqlQuery += " LIMIT ? OFFSET ?"
		offset := helper.GetSqlOffset(facilityTypeSearch.Page, facilityTypeSearch.Size)
		args = append(args, facilityTypeSearch.Size, offset)
	}

	rows := FetchRows(ctx, tx, sqlQuery, args)
	defer rows.Close()

	return getFacilityTypes(rows)
}

func (repository *FacilityTypeRepositoryImpl) FindTotalItem(ctx context.Context, tx *sql.Tx, searchBy any) int {
	facilityTypeSearch := searchBy.(search.FacilityTypeSearch)

	sqlSearch, args := sqlSearchFacilityTypeBy(facilityTypeSearch.Name)
	sqlQuery := sqlFindTotalFacilityType() + sqlSearch

	rows := FetchRows(ctx, tx, sqlQuery, args)
	defer rows.Close()

	return ScanTotalItem(rows)
}

func sqlSaveFacilityType() string {
	return "INSERT INTO m_facility_type(name," +
		" created_at," +
		" created_by," +
		" created_by_name," +
		" updated_at," +
		" updated_by," +
		" updated_by_name," +
		" is_deleted) VALUES (?,?,?,?,?,?,?,?)"
}

func sqlUpdateFacilityType() string {
	return "UPDATE m_facility_type SET name = ?," +
		" updated_at = ?," +
		" updated_by = ?," +
		" updated_by_name = ?" +
		" WHERE id = ?"
}

func sqlDeleteFacilityType() string {
	return "UPDATE m_facility_type SET updated_at = ?," +
		" updated_by = ?," +
		" updated_by_name = ?," +
		" is_deleted = ?" +
		" WHERE id = ?"
}

func sqlSelectFacilityType() string {
	return "SELECT id," +
		" name," +
		" created_at," +
		" created_by," +
		" created_by_name," +
		" updated_at," +
		" updated_by," +
		" updated_by_name," +
		" is_deleted FROM m_facility_type" +
		" WHERE is_deleted = false"
}

func sqlFindTotalFacilityType() string {
	return "SELECT COUNT(1) AS totalItem FROM m_facility_type WHERE is_deleted = false"
}

func sqlSearchFacilityTypeBy(name string) (string, []any) {
	var args []any
	sqlQuery := ""

	if len(name) != 0 {
		sqlQuery += " AND LOWER(name) LIKE ?"
		args = append(args, helper.StringQueryLike(name))
	}

	sqlQuery += " ORDER BY id ASC "
	return sqlQuery, args
}

func getFacilityTypes(rows *sql.Rows) []domain.FacilityType {
	var facilityTypes []domain.FacilityType
	for rows.Next() {
		facilityType := domain.FacilityType{}
		scanFacilityType(rows, &facilityType)
		facilityTypes = append(facilityTypes, facilityType)
	}
	return facilityTypes
}

func scanFacilityType(rows *sql.Rows, facilityType *domain.FacilityType) {
	err := rows.Scan(
		&facilityType.Id,
		&facilityType.Name,
		&facilityType.CreatedAt,
		&facilityType.CreatedBy,
		&facilityType.CreatedByName,
		&facilityType.UpdatedAt,
		&facilityType.UpdatedBy,
		&facilityType.UpdatedByName,
		&facilityType.IsDeleted,
	)
	helper.PanicIfError(err)
}
