package repository

import (
	"context"
	"database/sql"
	"minang-kos-service/helper"
	"minang-kos-service/model/domain"
)

type CountryRepositoryImpl struct {
}

func NewCountryRepository() CountryRepository {
	return &CountryRepositoryImpl{}
}

func (repository *CountryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, domainModel any) any {
	SQL := "insert into m_country(name, created_at, created_by, created_by_name, " +
		"updated_at, updated_by, updated_by_name, is_deleted) values (?,?,?,?,?,?,?,?)"

	country := domainModel.(domain.Country)
	result, err := tx.ExecContext(ctx, SQL, country.Name, country.CreatedAt, country.CreatedBy, country.CreatedByName,
		country.UpdatedAt, country.UpdatedBy, country.UpdatedByName, country.IsDeleted)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	country.Id = id

	return country
}
