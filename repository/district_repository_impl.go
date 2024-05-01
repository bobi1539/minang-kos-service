package repository

import (
	"context"
	"database/sql"
	"errors"
	"minang-kos-service/constant"
	"minang-kos-service/helper"
	"minang-kos-service/model/domain"
)

type DistrictRepositoryImpl struct {
}

func NewDistrictRepository() DistrictRepository {
	return &DistrictRepositoryImpl{}
}

func (repository *DistrictRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, domainModel any) any {
	district := domainModel.(domain.District)
	result, err := tx.ExecContext(
		ctx, sqlSaveDistrict(), district.Name, district.City.Id, district.CreatedAt, district.CreatedBy, district.CreatedByName,
		district.UpdatedAt, district.UpdatedBy, district.UpdatedByName, district.IsDeleted,
	)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	district.Id = id
	return district
}

func (repository *DistrictRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, domainModel any) any {
	district := domainModel.(domain.District)
	_, err := tx.ExecContext(
		ctx, sqlUpdateDistrict(), district.Name, district.City.Id,
		district.UpdatedAt, district.UpdatedBy, district.UpdatedByName, district.Id,
	)
	helper.PanicIfError(err)

	return district
}

func (repository *DistrictRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, domainModel any) {
	district := domainModel.(domain.District)
	_, err := tx.ExecContext(
		ctx, sqlDeleteDistrict(), district.UpdatedAt, district.UpdatedBy,
		district.UpdatedByName, district.IsDeleted, district.Id,
	)
	helper.PanicIfError(err)
}

func (repository *DistrictRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int64) (any, error) {
	sqlQuery := sqlSelectDistrict() + " AND md.id = ?"
	rows, err := tx.QueryContext(ctx, sqlQuery, id)
	helper.PanicIfError(err)
	defer rows.Close()

	district := domain.District{}
	if rows.Next() {
		scanDistrict(rows, &district)
		return district, nil
	}
	return district, errors.New(constant.DATA_NOT_FOUND)
}

func (repository *DistrictRepositoryImpl) FindAllWithPagination(ctx context.Context, tx *sql.Tx, searchBy map[string]any) any {
	name := searchBy["name"].(string)
	cityId := searchBy["cityId"].(int64)
	page := searchBy["page"].(int)
	size := searchBy["size"].(int)

	page = helper.GetSqlOffset(page, size)

	sqlSearch, args := sqlSearchByDistrict(name, cityId)
	sqlQuery := sqlSelectDistrict() + sqlSearch + " ORDER BY md.id ASC LIMIT ? OFFSET ?"
	args = append(args, size, page)

	rows := FetchRows(ctx, tx, sqlQuery, args)
	defer rows.Close()

	return getDistricts(rows)
}

func (repository *DistrictRepositoryImpl) FindAllWithoutPagination(ctx context.Context, tx *sql.Tx, searchBy map[string]any) any {
	name := searchBy["name"].(string)
	cityId := searchBy["cityId"].(int64)

	sqlSearch, args := sqlSearchByDistrict(name, cityId)
	sqlQuery := sqlSelectDistrict() + sqlSearch + " ORDER BY md.id ASC"

	rows := FetchRows(ctx, tx, sqlQuery, args)
	defer rows.Close()

	return getDistricts(rows)
}

func (repository *DistrictRepositoryImpl) FindTotalItem(ctx context.Context, tx *sql.Tx, searchBy map[string]any) int {
	name := searchBy["name"].(string)
	cityId := searchBy["cityId"].(int64)

	sqlSearch, args := sqlSearchByDistrict(name, cityId)
	sqlQuery := sqlFindTotalDistrict() + sqlSearch

	rows := FetchRows(ctx, tx, sqlQuery, args)
	defer rows.Close()

	return ScanTotalItem(rows)
}

func sqlSaveDistrict() string {
	return "INSERT INTO m_district(name," +
		" city_id," +
		" created_at," +
		" created_by," +
		" created_by_name," +
		" updated_at," +
		" updated_by," +
		" updated_by_name," +
		" is_deleted) VALUES (?,?,?,?,?,?,?,?,?)"
}

func sqlUpdateDistrict() string {
	return "UPDATE m_district SET name = ?," +
		" city_id = ?," +
		" updated_at = ?," +
		" updated_by = ?," +
		" updated_by_name = ? WHERE id = ?"
}

func sqlDeleteDistrict() string {
	return "UPDATE m_district SET updated_at = ?," +
		" updated_by = ?," +
		" updated_by_name = ?," +
		" is_deleted = ?" +
		" WHERE id = ?"
}

func sqlSelectDistrict() string {
	return "SELECT md.id," +
		" md.name," +
		" md.city_id," +
		" md.created_at," +
		" md.created_by," +
		" md.created_by_name," +
		" md.updated_at," +
		" md.updated_by," +
		" md.updated_by_name," +
		" md.is_deleted," +
		" mct.name," +
		" mct.province_id," +
		" mct.created_at," +
		" mct.created_by," +
		" mct.created_by_name," +
		" mct.updated_at," +
		" mct.updated_by," +
		" mct.updated_by_name," +
		" mct.is_deleted," +
		" mp.name," +
		" mp.country_id," +
		" mp.created_at," +
		" mp.created_by," +
		" mp.created_by_name," +
		" mp.updated_at," +
		" mp.updated_by," +
		" mp.updated_by_name," +
		" mp.is_deleted," +
		" mc.name," +
		" mc.created_at," +
		" mc.created_by," +
		" mc.created_by_name," +
		" mc.updated_at," +
		" mc.updated_by," +
		" mc.updated_by_name," +
		" mc.is_deleted FROM m_district md" +
		" LEFT JOIN m_city mct ON md.city_id = mct.id" +
		" LEFT JOIN m_province mp ON mct.province_id = mp.id" +
		" LEFT JOIN m_country mc ON mp.country_id = mc.id" +
		" WHERE md.is_deleted = false"
}

func sqlFindTotalDistrict() string {
	return "SELECT COUNT(1) AS totalItem FROM m_district md WHERE md.is_deleted = false"
}

func sqlSearchByDistrict(name string, cityId int64) (string, []any) {
	var args []any
	sqlQuery := ""

	if len(name) != 0 {
		sqlQuery += " AND LOWER(md.name) LIKE ?"
		args = append(args, helper.StringQueryLike(name))
	}

	if cityId != 0 {
		sqlQuery += " AND md.city_id = ?"
		args = append(args, cityId)
	}
	return sqlQuery, args
}

func scanDistrict(rows *sql.Rows, district *domain.District) {
	err := rows.Scan(
		&district.Id,
		&district.Name,
		&district.City.Id,
		&district.CreatedAt,
		&district.CreatedBy,
		&district.CreatedByName,
		&district.UpdatedAt,
		&district.UpdatedBy,
		&district.UpdatedByName,
		&district.IsDeleted,
		&district.City.Name,
		&district.City.Province.Id,
		&district.City.CreatedAt,
		&district.City.CreatedBy,
		&district.City.CreatedByName,
		&district.City.UpdatedAt,
		&district.City.UpdatedBy,
		&district.City.UpdatedByName,
		&district.City.IsDeleted,
		&district.City.Province.Name,
		&district.City.Province.Country.Id,
		&district.City.Province.CreatedAt,
		&district.City.Province.CreatedBy,
		&district.City.Province.CreatedByName,
		&district.City.Province.UpdatedAt,
		&district.City.Province.UpdatedBy,
		&district.City.Province.UpdatedByName,
		&district.City.Province.IsDeleted,
		&district.City.Province.Country.Name,
		&district.City.Province.Country.CreatedAt,
		&district.City.Province.Country.CreatedBy,
		&district.City.Province.Country.CreatedByName,
		&district.City.Province.Country.UpdatedAt,
		&district.City.Province.Country.UpdatedBy,
		&district.City.Province.Country.UpdatedByName,
		&district.City.Province.Country.IsDeleted,
	)
	helper.PanicIfError(err)
}

func getDistricts(rows *sql.Rows) []domain.District {
	var districts []domain.District
	for rows.Next() {
		district := domain.District{}
		scanDistrict(rows, &district)
		districts = append(districts, district)
	}
	return districts
}
