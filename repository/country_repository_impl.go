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
	sqlQuery := "insert into m_country(name, created_at, created_by, created_by_name, " +
		"updated_at, updated_by, updated_by_name, is_deleted) values (?,?,?,?,?,?,?,?)"

	country := domainModel.(domain.Country)
	result, err := tx.ExecContext(ctx, sqlQuery, country.Name, country.CreatedAt, country.CreatedBy, country.CreatedByName,
		country.UpdatedAt, country.UpdatedBy, country.UpdatedByName, country.IsDeleted)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	country.Id = id

	return country
}

func (repository *CountryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, domainModel any) any {
	sqlQuery := "update m_country set name = ?, updated_at = ?, updated_by = ?, updated_by_name = ? where id = ?"

	country := domainModel.(domain.Country)
	_, err := tx.ExecContext(ctx, sqlQuery, country.Name, country.UpdatedAt, country.UpdatedBy, country.UpdatedByName, country.Id)
	helper.PanicIfError(err)

	return country
}

func (repository *CountryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, domainModel any) {
	sqlQuery := "update m_country set updated_at = ?, updated_by = ?, updated_by_name = ?, is_deleted = ? where id = ?"

	country := domainModel.(domain.Country)
	_, err := tx.ExecContext(ctx, sqlQuery, country.UpdatedAt, country.UpdatedBy, country.UpdatedByName, true, country.Id)
	helper.PanicIfError(err)
}

func (repository *CountryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int64) (any, error) {
	sqlQuery := "select id, name, created_at, created_by, created_by_name, " +
		"updated_at, updated_by, updated_by_name, is_deleted from m_country where id = ? and is_deleted = false"

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

func (repository *CountryRepositoryImpl) FindAllWithPagination(ctx context.Context, tx *sql.Tx, searchBy map[any]any) (any, int64) {
	name := searchBy["name"]
	page := searchBy["page"]
	size := searchBy["size"]

	page = page.(int) - 1

	sqlQuery := "select id, name, created_at, created_by, created_by_name, updated_at, updated_by, updated_by_name, is_deleted" +
		" from m_country where is_deleted = false"
	if name != nil {
		sqlQuery += " and name like %" + name.(string) + "%"
	}

	sqlQuery += " limit " + size.(string) + " offset " + page.(string)
	sqlQuery += " order by id asc"

	var rows *sql.Rows
	var err error

	if name != nil {
		rows, err = tx.QueryContext(ctx, sqlQuery, name, size, page)
	} else {
		rows, err = tx.QueryContext(ctx, sqlQuery, size, page)
	}
	helper.PanicIfError(err)
	defer rows.Close()

	return getCountries(rows), getTotalItem(ctx, tx, name)
}

func (repository *CountryRepositoryImpl) FindAllWithoutPagination(ctx context.Context, tx *sql.Tx, searchBy map[any]any) any {
	name := searchBy["name"]

	sqlQuery := "select id, name, created_at, created_by, created_by_name, updated_at, updated_by, updated_by_name, is_deleted" +
		" from m_country where is_deleted = false"
	if name != nil {
		sqlQuery += " and name like %" + name.(string) + "%"
	}
	sqlQuery += " order by id asc"

	var rows *sql.Rows
	var err error

	if name != nil {
		rows, err = tx.QueryContext(ctx, sqlQuery, name)
	} else {
		rows, err = tx.QueryContext(ctx, sqlQuery)
	}
	helper.PanicIfError(err)
	defer rows.Close()

	return getCountries(rows)
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

func getTotalItem(ctx context.Context, tx *sql.Tx, name any) int64 {
	sqlQuery := "select count(1) as totalItem from m_country where is_deleted = false"
	if name != nil {
		sqlQuery += " and name like %" + name.(string) + "%"
	}

	var rows *sql.Rows
	var err error

	if name != nil {
		rows, err = tx.QueryContext(ctx, sqlQuery, name)
	} else {
		rows, err = tx.QueryContext(ctx, sqlQuery)
	}
	helper.PanicIfError(err)
	defer rows.Close()

	var totalItem int64
	for rows.Next() {
		err := rows.Scan(&totalItem)
		helper.PanicIfError(err)
	}
	return totalItem
}
