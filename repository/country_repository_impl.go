package repository

import (
	"context"
	"database/sql"
	"errors"
	"minang-kos-service/constant"
	"minang-kos-service/helper"
	"minang-kos-service/model/domain"
)

type CountryRepositoryImpl struct {
}

func NewCountryRepository() CountryRepository {
	return &CountryRepositoryImpl{}
}

func (repository *CountryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, domainModel any) any {
	country := domainModel.(domain.Country)
	result, err := tx.ExecContext(
		ctx, getSqlSaveCountry(), country.Name, country.CreatedAt, country.CreatedBy, country.CreatedByName,
		country.UpdatedAt, country.UpdatedBy, country.UpdatedByName, country.IsDeleted,
	)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	country.Id = id
	return country
}

func (repository *CountryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, domainModel any) any {
	country := domainModel.(domain.Country)
	_, err := tx.ExecContext(ctx, getSqlUpdate(), country.Name, country.UpdatedAt, country.UpdatedBy, country.UpdatedByName, country.Id)
	helper.PanicIfError(err)

	return country
}

func (repository *CountryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, domainModel any) {
	country := domainModel.(domain.Country)
	_, err := tx.ExecContext(ctx, getSqlDelete(), country.UpdatedAt, country.UpdatedBy, country.UpdatedByName, true, country.Id)
	helper.PanicIfError(err)
}

func (repository *CountryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int64) (any, error) {
	rows, err := tx.QueryContext(ctx, getSqlFindById(), id)
	helper.PanicIfError(err)
	defer rows.Close()

	country := domain.Country{}

	if rows.Next() {
		scanCountry(rows, &country)
		return country, nil
	}
	return country, errors.New(constant.DATA_NOT_FOUND)
}

func (repository *CountryRepositoryImpl) FindAllWithPagination(ctx context.Context, tx *sql.Tx, searchBy map[string]any) any {
	name := searchBy["name"].(string)
	page := searchBy["page"].(int)
	size := searchBy["size"].(int)

	page = helper.GetSqlOffset(page, size)

	sqlQuery := getSqlFindAll() + addSearchByName(name)
	sqlQuery += " ORDER BY id ASC LIMIT ? OFFSET ?"

	var rows *sql.Rows
	var err error

	if len(name) != 0 {
		rows, err = tx.QueryContext(ctx, sqlQuery, helper.StringQueryLike(name), size, page)
	} else {
		rows, err = tx.QueryContext(ctx, sqlQuery, size, page)
	}
	helper.PanicIfError(err)
	defer rows.Close()

	return getCountries(rows)
}

func (repository *CountryRepositoryImpl) FindAllWithoutPagination(ctx context.Context, tx *sql.Tx, searchBy map[string]any) any {
	name := searchBy["name"].(string)

	sqlQuery := getSqlFindAll() + addSearchByName(name)
	sqlQuery += " ORDER BY id ASC"

	rows := fetchRows(ctx, tx, sqlQuery, name)
	defer rows.Close()

	return getCountries(rows)
}

func (repository *CountryRepositoryImpl) FindTotalItem(ctx context.Context, tx *sql.Tx, searchBy map[string]any) int {
	name := searchBy["name"].(string)
	sqlQuery := getSqlFindTotalItem() + addSearchByName(name)

	rows := fetchRows(ctx, tx, sqlQuery, name)
	defer rows.Close()

	return scanTotalItem(rows)
}

func getCountries(rows *sql.Rows) []domain.Country {
	var countries []domain.Country
	for rows.Next() {
		country := domain.Country{}
		scanCountry(rows, &country)
		countries = append(countries, country)
	}
	return countries
}

func scanCountry(rows *sql.Rows, country *domain.Country) {
	err := rows.Scan(&country.Id, &country.Name, &country.CreatedAt, &country.CreatedBy, &country.CreatedByName,
		&country.UpdatedAt, &country.UpdatedBy, &country.UpdatedByName, &country.IsDeleted)
	helper.PanicIfError(err)
}

func getSqlSaveCountry() string {
	return "INSERT INTO m_country(name, created_at, created_by, created_by_name, updated_at, updated_by, updated_by_name, is_deleted) VALUES (?,?,?,?,?,?,?,?)"
}

func getSqlUpdate() string {
	return "UPDATE m_country SET name = ?, updated_at = ?, updated_by = ?, updated_by_name = ? WHERE id = ?"
}

func getSqlDelete() string {
	return "UPDATE m_country SET updated_at = ?, updated_by = ?, updated_by_name = ?, is_deleted = ? WHERE id = ?"
}

func getSqlFindById() string {
	return "SELECT id, name, created_at, created_by, created_by_name, updated_at, updated_by, updated_by_name, is_deleted FROM m_country WHERE id = ? AND is_deleted = false"
}

func getSqlFindAll() string {
	return "SELECT id, name, created_at, created_by, created_by_name, updated_at, updated_by, updated_by_name, is_deleted FROM m_country WHERE is_deleted = false"
}

func getSqlFindTotalItem() string {
	return "SELECT COUNT(1) AS totalItem FROM m_country WHERE is_deleted = false"
}

func addSearchByName(name string) string {
	if len(name) != 0 {
		return " AND LOWER(name) LIKE ?"
	}
	return ""
}

func fetchRows(ctx context.Context, tx *sql.Tx, sqlQuery string, name string) *sql.Rows {
	if len(name) != 0 {
		rows, err := tx.QueryContext(ctx, sqlQuery, helper.StringQueryLike(name))
		helper.PanicIfError(err)
		return rows
	}

	rows, err := tx.QueryContext(ctx, sqlQuery)
	helper.PanicIfError(err)
	return rows
}

func scanTotalItem(rows *sql.Rows) int {
	var totalItem int
	for rows.Next() {
		err := rows.Scan(&totalItem)
		helper.PanicIfError(err)
	}
	return totalItem
}
