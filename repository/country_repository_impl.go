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

func (repository *CountryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, searchBy any) any {
	countrySearch := searchBy.(search.CountrySearch)
	sqlSearch, args := sqlSearchByCountry(countrySearch.Name)
	sqlQuery := sqlSelectCountry() + sqlSearch

	if countrySearch.Page > 0 {
		sqlQuery += " LIMIT ? OFFSET ?"
		offset := helper.GetSqlOffset(countrySearch.Page, countrySearch.Size)
		args = append(args, countrySearch.Size, offset)
	}

	rows := FetchRows(ctx, tx, sqlQuery, args)
	defer rows.Close()

	return getCountries(rows)
}

func (repository *CountryRepositoryImpl) FindTotalItem(ctx context.Context, tx *sql.Tx, searchBy any) int {
	countrySearch := searchBy.(search.CountrySearch)

	sqlSearch, args := sqlSearchByCountry(countrySearch.Name)
	sqlQuery := sqlFindTotalCountry() + sqlSearch

	rows := FetchRows(ctx, tx, sqlQuery, args)
	defer rows.Close()

	return ScanTotalItem(rows)
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

	sqlQuery += " ORDER BY id ASC"
	return sqlQuery, args
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
	err := rows.Scan(
		&country.Id,
		&country.Name,
		&country.CreatedAt,
		&country.CreatedBy,
		&country.CreatedByName,
		&country.UpdatedAt,
		&country.UpdatedBy,
		&country.UpdatedByName,
		&country.IsDeleted,
	)
	helper.PanicIfError(err)
}
