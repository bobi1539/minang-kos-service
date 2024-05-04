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

type FacilityRepositoryImpl struct {
}

func NewFacilityRepository() FacilityRepository {
	return &FacilityRepositoryImpl{}
}

func (repository *FacilityRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, domainModel any) any {
	facility := domainModel.(domain.Facility)
	result, err := tx.ExecContext(ctx, sqlSaveFacility(), argsSaveFacility(facility)...)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	facility.Id = id
	return facility
}

func (repository *FacilityRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, domainModel any) any {
	facility := domainModel.(domain.Facility)
	_, err := tx.ExecContext(ctx, sqlUpdateFacility(), argsUpdateFacility(facility)...)
	helper.PanicIfError(err)

	return facility
}

func (repository *FacilityRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, domainModel any) {
	facility := domainModel.(domain.Facility)
	_, err := tx.ExecContext(
		ctx, sqlDeleteFacility(), facility.UpdatedAt, facility.UpdatedBy,
		facility.UpdatedByName, facility.IsDeleted, facility.Id,
	)
	helper.PanicIfError(err)
}

func (repository *FacilityRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int64) (any, error) {
	sqlQuery := sqlSelectFacility() + " AND mf.id = ?"
	rows, err := tx.QueryContext(ctx, sqlQuery, id)
	helper.PanicIfError(err)
	defer rows.Close()

	facility := domain.Facility{}
	if rows.Next() {
		scanFacility(rows, &facility)
		return facility, nil
	}
	return facility, errors.New(constant.DATA_NOT_FOUND)
}

func (repository *FacilityRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, searchBy any) any {
	facilitySearch := searchBy.(search.FacilitySearch)

	sqlSearch, args := sqlSearchFacilityBy(facilitySearch.Name, facilitySearch.FacilityTypeId)
	sqlQuery := sqlSelectFacility() + sqlSearch

	if facilitySearch.Page > 0 {
		sqlQuery += " LIMIT ? OFFSET ?"
		offset := helper.GetSqlOffset(facilitySearch.Page, facilitySearch.Size)
		args = append(args, facilitySearch.Size, offset)
	}

	rows := FetchRows(ctx, tx, sqlQuery, args)
	defer rows.Close()

	return getFacilities(rows)
}

func (repository *FacilityRepositoryImpl) FindTotalItem(ctx context.Context, tx *sql.Tx, searchBy any) int {
	facilitySearch := searchBy.(search.FacilitySearch)

	sqlSearch, args := sqlSearchFacilityBy(facilitySearch.Name, facilitySearch.FacilityTypeId)
	sqlQuery := sqlFindTotalFacility() + sqlSearch

	rows := FetchRows(ctx, tx, sqlQuery, args)
	defer rows.Close()

	return ScanTotalItem(rows)
}

func sqlSaveFacility() string {
	return "INSERT INTO m_facility(name," +
		" icon," +
		" facility_type_id," +
		" created_at," +
		" created_by," +
		" created_by_name," +
		" updated_at," +
		" updated_by," +
		" updated_by_name," +
		" is_deleted) VALUES (?,?,?,?,?,?,?,?,?,?)"
}

func sqlUpdateFacility() string {
	return "UPDATE m_facility SET name = ?," +
		" icon = ?," +
		" facility_type_id = ?," +
		" updated_at = ?," +
		" updated_by = ?," +
		" updated_by_name = ? WHERE id = ?"
}

func sqlDeleteFacility() string {
	return "UPDATE m_facility SET updated_at = ?," +
		" updated_by = ?," +
		" updated_by_name = ?," +
		" is_deleted = ?" +
		" WHERE id = ?"
}

func sqlSelectFacility() string {
	return "SELECT mf.id," +
		" mf.name," +
		" mf.icon," +
		" mf.facility_type_id," +
		" mf.created_at," +
		" mf.created_by," +
		" mf.created_by_name," +
		" mf.updated_at," +
		" mf.updated_by," +
		" mf.updated_by_name," +
		" mf.is_deleted," +
		" mft.name," +
		" mft.created_at," +
		" mft.created_by," +
		" mft.created_by_name," +
		" mft.updated_at," +
		" mft.updated_by," +
		" mft.updated_by_name," +
		" mft.is_deleted FROM m_facility mf" +
		" LEFT JOIN m_facility_type mft ON mf.facility_type_id = mft.id" +
		" WHERE mf.is_deleted = false"
}

func sqlFindTotalFacility() string {
	return "SELECT COUNT(1) AS totalItem FROM m_facility mf WHERE mf.is_deleted = false"
}

func sqlSearchFacilityBy(name string, facilityTypeId int64) (string, []any) {
	var args []any
	sqlQuery := ""

	if len(name) != 0 {
		sqlQuery += " AND LOWER(mf.name) LIKE ?"
		args = append(args, helper.StringQueryLike(name))
	}

	if facilityTypeId != 0 {
		sqlQuery += " AND mf.facility_type_id = ?"
		args = append(args, facilityTypeId)
	}

	sqlQuery += " ORDER BY mf.id ASC"
	return sqlQuery, args
}

func getFacilities(rows *sql.Rows) []domain.Facility {
	var facilities []domain.Facility
	for rows.Next() {
		facility := domain.Facility{}
		scanFacility(rows, &facility)
		facilities = append(facilities, facility)
	}
	return facilities
}

func scanFacility(rows *sql.Rows, facility *domain.Facility) {
	err := rows.Scan(
		&facility.Id,
		&facility.Name,
		&facility.Icon,
		&facility.FacilityType.Id,
		&facility.CreatedAt,
		&facility.CreatedBy,
		&facility.CreatedByName,
		&facility.UpdatedAt,
		&facility.UpdatedBy,
		&facility.UpdatedByName,
		&facility.IsDeleted,
		&facility.FacilityType.Name,
		&facility.FacilityType.CreatedAt,
		&facility.FacilityType.CreatedBy,
		&facility.FacilityType.CreatedByName,
		&facility.FacilityType.UpdatedAt,
		&facility.FacilityType.UpdatedBy,
		&facility.FacilityType.UpdatedByName,
		&facility.FacilityType.IsDeleted,
	)
	helper.PanicIfError(err)
}

func argsSaveFacility(facility domain.Facility) []any {
	var args []any
	args = append(args, facility.Name)
	args = append(args, facility.Icon)
	args = append(args, facility.FacilityType.Id)
	args = append(args, facility.CreatedAt)
	args = append(args, facility.CreatedBy)
	args = append(args, facility.CreatedByName)
	args = append(args, facility.UpdatedAt)
	args = append(args, facility.UpdatedBy)
	args = append(args, facility.UpdatedByName)
	args = append(args, facility.IsDeleted)
	return args
}

func argsUpdateFacility(facility domain.Facility) []any {
	var args []any
	args = append(args, facility.Name)
	args = append(args, facility.Icon)
	args = append(args, facility.FacilityType.Id)
	args = append(args, facility.UpdatedAt)
	args = append(args, facility.UpdatedBy)
	args = append(args, facility.UpdatedByName)
	args = append(args, facility.IsDeleted)
	return args
}
