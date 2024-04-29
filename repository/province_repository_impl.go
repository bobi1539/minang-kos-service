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
		ctx, getSqlSaveProvince(), province.Name, province.Country.Id, province.CreatedAt, province.CreatedBy, province.CreatedByName,
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
		ctx, getSqlUpdateProvince(), province.Name, province.Country.Id,
		province.UpdatedAt, province.UpdatedBy, province.UpdatedByName, province.Id,
	)
	helper.PanicIfError(err)

	return province
}

func (repository *ProvinceRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, domainModel any) {
	panic("imp")
}

func (repository *ProvinceRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int64) (any, error) {
	rows, err := tx.QueryContext(ctx, getSqlFindByIdProvince(), id)
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
	panic("imp")
}

func (repository *ProvinceRepositoryImpl) FindTotalItem(ctx context.Context, tx *sql.Tx, searchBy map[string]any) int {
	panic("imp")
}

func getSqlSaveProvince() string {
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

func getSqlUpdateProvince() string {
	return "UPDATE m_province SET name = ?," +
		" country_id = ?," +
		" updated_at = ?," +
		" updated_by = ?," +
		" updated_by_name = ? WHERE id = ?"
}

func getSqlFindByIdProvince() string {
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
		" LEFT JOIN m_country mc ON mp.country_id = mc.id WHERE mp.id = ? AND mp.is_deleted = false"
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