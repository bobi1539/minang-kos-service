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
		ctx, sqlSaveCountry(), country.Name, country.CreatedAt, country.CreatedBy, country.CreatedByName,
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
	_, err := tx.ExecContext(ctx, sqlUpdateCountry(), country.Name, country.UpdatedAt, country.UpdatedBy, country.UpdatedByName, country.Id)
	helper.PanicIfError(err)

	return country
}

func (repository *CountryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, domainModel any) {
	country := domainModel.(domain.Country)
	_, err := tx.ExecContext(
		ctx, sqlDeleteCountry(), country.UpdatedAt, country.UpdatedBy, country.UpdatedByName, country.IsDeleted, country.Id,
	)
	helper.PanicIfError(err)
}

func (repository *CountryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int64) (any, error) {
	sqlQuery := sqlSelectCountry() + " AND id = ?"
	rows, err := tx.QueryContext(ctx, sqlQuery, id)
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

	sqlSearch, args := sqlSearchByCountry(name)
	sqlQuery := sqlSelectCountry() + sqlSearch + " ORDER BY id ASC LIMIT ? OFFSET ?"
	args = append(args, size, page)

	rows := FetchRows(ctx, tx, sqlQuery, args)
	defer rows.Close()

	return getCountries(rows)
}

func (repository *CountryRepositoryImpl) FindAllWithoutPagination(ctx context.Context, tx *sql.Tx, searchBy map[string]any) any {
	name := searchBy["name"].(string)

	sqlSearch, args := sqlSearchByCountry(name)
	sqlQuery := sqlSelectCountry() + sqlSearch + " ORDER BY id ASC"

	rows := FetchRows(ctx, tx, sqlQuery, args)
	defer rows.Close()

	return getCountries(rows)
}

func (repository *CountryRepositoryImpl) FindTotalItem(ctx context.Context, tx *sql.Tx, searchBy map[string]any) int {
	name := searchBy["name"].(string)

	sqlSearch, args := sqlSearchByCountry(name)
	sqlQuery := sqlFindTotalCountry() + sqlSearch

	rows := FetchRows(ctx, tx, sqlQuery, args)
	defer rows.Close()

	return ScanTotalItem(rows)
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

func sqlSaveCountry() string {
	return "INSERT INTO m_country(name," +
		" created_at," +
		" created_by," +
		" created_by_name," +
		" updated_at," +
		" updated_by," +
		" updated_by_name," +
		" is_deleted) VALUES (?,?,?,?,?,?,?,?)"
}

func sqlUpdateCountry() string {
	return "UPDATE m_country SET name = ?," +
		" updated_at = ?," +
		" updated_by = ?," +
		" updated_by_name = ?" +
		" WHERE id = ?"
}

func sqlDeleteCountry() string {
	return "UPDATE m_country SET updated_at = ?," +
		" updated_by = ?," +
		" updated_by_name = ?," +
		" is_deleted = ?" +
		" WHERE id = ?"
}

func sqlSelectCountry() string {
	return "SELECT id," +
		" name," +
		" created_at," +
		" created_by," +
		" created_by_name," +
		" updated_at," +
		" updated_by," +
		" updated_by_name," +
		" is_deleted FROM m_country" +
		" WHERE is_deleted = false"
}

func sqlFindTotalCountry() string {
	return "SELECT COUNT(1) AS totalItem FROM m_country WHERE is_deleted = false"
}

func sqlSearchByCountry(name string) (string, []any) {
	var args []any
	sqlQuery := ""

	if len(name) != 0 {
		sqlQuery += " AND LOWER(name) LIKE ?"
		args = append(args, helper.StringQueryLike(name))
	}
	return sqlQuery, args
}
