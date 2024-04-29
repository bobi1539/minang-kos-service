package repository

import (
	"context"
	"database/sql"
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
	panic("imp")
}

func (repository *ProvinceRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, domainModel any) {
	panic("imp")
}

func (repository *ProvinceRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int64) (any, error) {
	panic("imp")
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
	return "INSERT INTO m_province(name, country_id, created_at, created_by, created_by_name, updated_at, updated_by, updated_by_name, is_deleted) VALUES (?,?,?,?,?,?,?,?,?)"
}
