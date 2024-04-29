package repository

import (
	"context"
	"database/sql"
	"errors"
	"minang-kos-service/constant"
	"minang-kos-service/helper"
	"minang-kos-service/model/domain"
)

type ProvinceRepositoryImpl struct {
}

func NewProvinceRepository() ProvinceRepository {
	return &ProvinceRepositoryImpl{}
}

func (repository *ProvinceRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, domainModel any) any {
	province := domainModel.(domain.Province)
	result, err := tx.ExecContext(
		ctx, sqlSaveProvince(), province.Name, province.Country.Id, province.CreatedAt, province.CreatedBy, province.CreatedByName,
		province.UpdatedAt, province.UpdatedBy, province.UpdatedByName, province.IsDeleted,
	)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	province.Id = id
	return province
}

func (repository *ProvinceRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, domainModel any) any {
	province := domainModel.(domain.Province)
	_, err := tx.ExecContext(
		ctx, sqlUpdateProvince(), province.Name, province.Country.Id,
		province.UpdatedAt, province.UpdatedBy, province.UpdatedByName, province.Id,
	)
	helper.PanicIfError(err)

	return province
}

func (repository *ProvinceRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, domainModel any) {
	province := domainModel.(domain.Province)
	_, err := tx.ExecContext(
		ctx, sqlDeleteProvince(), province.UpdatedAt, province.UpdatedBy,
		province.UpdatedByName, province.IsDeleted, province.Id,
	)
	helper.PanicIfError(err)
}

func (repository *ProvinceRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int64) (any, error) {
	sqlQuery := sqlSelectProvince() + " AND mp.id = ?"
	rows, err := tx.QueryContext(ctx, sqlQuery, id)
	helper.PanicIfError(err)
	defer rows.Close()

	province := domain.Province{}
	if rows.Next() {
		scanProvince(rows, &province)
		return province, nil
	}
	return province, errors.New(constant.DATA_NOT_FOUND)
}

func (repository *ProvinceRepositoryImpl) FindAllWithPagination(ctx context.Context, tx *sql.Tx, searchBy map[string]any) any {
	panic("imp")
}

func (repository *ProvinceRepositoryImpl) FindAllWithoutPagination(ctx context.Context, tx *sql.Tx, searchBy map[string]any) any {
	name := searchBy["name"].(string)
	countryId := searchBy["countryId"].(int64)

	sqlQuery := sqlSelectProvince() + sqlSearchByProvince(name)
	sqlQuery += " AND mp.country_id = ? ORDER BY mp.id ASC"

	rows := fetchRowsProvince(ctx, tx, sqlQuery, name, countryId)
	defer rows.Close()

	return getProvinces(rows)
}

func (repository *ProvinceRepositoryImpl) FindTotalItem(ctx context.Context, tx *sql.Tx, searchBy map[string]any) int {
	panic("imp")
}

func sqlSaveProvince() string {
	return "INSERT INTO m_province(name," +
		" country_id," +
		" created_at," +
		" created_by," +
		" created_by_name," +
		" updated_at," +
		" updated_by," +
		" updated_by_name," +
		" is_deleted) VALUES (?,?,?,?,?,?,?,?,?)"
}

func sqlUpdateProvince() string {
	return "UPDATE m_province SET name = ?," +
		" country_id = ?," +
		" updated_at = ?," +
		" updated_by = ?," +
		" updated_by_name = ? WHERE id = ?"
}

func sqlDeleteProvince() string {
	return "UPDATE m_province SET updated_at = ?," +
		" updated_by = ?," +
		" updated_by_name = ?," +
		" is_deleted = ?" +
		" WHERE id = ?"
}

func sqlSelectProvince() string {
	return "SELECT mp.id," +
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
		" mc.is_deleted FROM m_province mp" +
		" LEFT JOIN m_country mc ON mp.country_id = mc.id" +
		" WHERE mp.is_deleted = false"
}

func sqlSearchByProvince(name string) string {
	if len(name) != 0 {
		return " AND LOWER(mp.name) LIKE ?"
	}
	return ""
}

func scanProvince(rows *sql.Rows, province *domain.Province) {
	err := rows.Scan(
		&province.Id,
		&province.Name,
		&province.Country.Id,
		&province.CreatedAt,
		&province.CreatedBy,
		&province.CreatedByName,
		&province.UpdatedAt,
		&province.UpdatedBy,
		&province.UpdatedByName,
		&province.IsDeleted,
		&province.Country.Name,
		&province.Country.CreatedAt,
		&province.Country.CreatedBy,
		&province.Country.CreatedByName,
		&province.Country.UpdatedAt,
		&province.Country.UpdatedBy,
		&province.Country.UpdatedByName,
		&province.Country.IsDeleted)
	helper.PanicIfError(err)
}

func fetchRowsProvince(ctx context.Context, tx *sql.Tx, sqlQuery string, name string, countryId int64) *sql.Rows {
	if len(name) != 0 {
		rows, err := tx.QueryContext(ctx, sqlQuery, helper.StringQueryLike(name), countryId)
		helper.PanicIfError(err)
		return rows
	}

	rows, err := tx.QueryContext(ctx, sqlQuery, countryId)
	helper.PanicIfError(err)
	return rows
}

func getProvinces(rows *sql.Rows) []domain.Province {
	var provinces []domain.Province
	for rows.Next() {
		province := domain.Province{}
		scanProvince(rows, &province)
		provinces = append(provinces, province)
	}
	return provinces
}
