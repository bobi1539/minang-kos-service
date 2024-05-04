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

type CityRepositoryImpl struct {
}

func NewCityRepository() CityRepository {
	return &CityRepositoryImpl{}
}

func (repository *CityRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, domainModel any) any {
	city := domainModel.(domain.City)
	result, err := tx.ExecContext(
		ctx, sqlSaveCity(), city.Name, city.Province.Id, city.CreatedAt, city.CreatedBy, city.CreatedByName,
		city.UpdatedAt, city.UpdatedBy, city.UpdatedByName, city.IsDeleted,
	)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	city.Id = id
	return city
}

func (repository *CityRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, domainModel any) any {
	city := domainModel.(domain.City)
	_, err := tx.ExecContext(
		ctx, sqlUpdateCity(), city.Name, city.Province.Id,
		city.UpdatedAt, city.UpdatedBy, city.UpdatedByName, city.Id,
	)
	helper.PanicIfError(err)

	return city
}

func (repository *CityRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, domainModel any) {
	city := domainModel.(domain.City)
	_, err := tx.ExecContext(
		ctx, sqlDeleteCity(), city.UpdatedAt, city.UpdatedBy,
		city.UpdatedByName, city.IsDeleted, city.Id,
	)
	helper.PanicIfError(err)
}

func (repository *CityRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int64) (any, error) {
	sqlQuery := sqlSelectCity() + " AND mct.id = ?"
	rows, err := tx.QueryContext(ctx, sqlQuery, id)
	helper.PanicIfError(err)
	defer rows.Close()

	city := domain.City{}
	if rows.Next() {
		scanCity(rows, &city)
		return city, nil
	}
	return city, errors.New(constant.DATA_NOT_FOUND)
}

func (repository *CityRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, searchBy any) any {
	citySearch := searchBy.(search.CitySearch)

	sqlSearch, args := sqlSearchCityBy(citySearch.Name, citySearch.ProvinceId)
	sqlQuery := sqlSelectCity() + sqlSearch

	if citySearch.Page > 0 {
		sqlQuery += " LIMIT ? OFFSET ?"
		offset := helper.GetSqlOffset(citySearch.Page, citySearch.Size)
		args = append(args, citySearch.Size, offset)
	}

	rows := FetchRows(ctx, tx, sqlQuery, args)
	defer rows.Close()

	return getCities(rows)
}

func (repository *CityRepositoryImpl) FindTotalItem(ctx context.Context, tx *sql.Tx, searchBy any) int {
	citySearch := searchBy.(search.CitySearch)

	sqlSearch, args := sqlSearchCityBy(citySearch.Name, citySearch.ProvinceId)
	sqlQuery := sqlFindTotalCity() + sqlSearch

	rows := FetchRows(ctx, tx, sqlQuery, args)
	defer rows.Close()

	return ScanTotalItem(rows)
}

func sqlSaveCity() string {
	return "INSERT INTO m_city(name," +
		" province_id," +
		" created_at," +
		" created_by," +
		" created_by_name," +
		" updated_at," +
		" updated_by," +
		" updated_by_name," +
		" is_deleted) VALUES (?,?,?,?,?,?,?,?,?)"
}

func sqlUpdateCity() string {
	return "UPDATE m_city SET name = ?," +
		" province_id = ?," +
		" updated_at = ?," +
		" updated_by = ?," +
		" updated_by_name = ? WHERE id = ?"
}

func sqlDeleteCity() string {
	return "UPDATE m_city SET updated_at = ?," +
		" updated_by = ?," +
		" updated_by_name = ?," +
		" is_deleted = ?" +
		" WHERE id = ?"
}

func sqlSelectCity() string {
	return "SELECT mct.id," +
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
		" mc.is_deleted FROM m_city mct" +
		" LEFT JOIN m_province mp ON mct.province_id = mp.id" +
		" LEFT JOIN m_country mc ON mp.country_id = mc.id" +
		" WHERE mct.is_deleted = false"
}

func sqlFindTotalCity() string {
	return "SELECT COUNT(1) AS totalItem FROM m_city mct WHERE mct.is_deleted = false"
}

func sqlSearchCityBy(name string, provinceId int64) (string, []any) {
	var args []any
	sqlQuery := ""

	if len(name) != 0 {
		sqlQuery += " AND LOWER(mct.name) LIKE ?"
		args = append(args, helper.StringQueryLike(name))
	}

	if provinceId != 0 {
		sqlQuery += " AND mct.province_id = ?"
		args = append(args, provinceId)
	}

	sqlQuery += " ORDER BY mct.id ASC"
	return sqlQuery, args
}

func scanCity(rows *sql.Rows, city *domain.City) {
	err := rows.Scan(
		&city.Id,
		&city.Name,
		&city.Province.Id,
		&city.CreatedAt,
		&city.CreatedBy,
		&city.CreatedByName,
		&city.UpdatedAt,
		&city.UpdatedBy,
		&city.UpdatedByName,
		&city.IsDeleted,
		&city.Province.Name,
		&city.Province.Country.Id,
		&city.Province.CreatedAt,
		&city.Province.CreatedBy,
		&city.Province.CreatedByName,
		&city.Province.UpdatedAt,
		&city.Province.UpdatedBy,
		&city.Province.UpdatedByName,
		&city.Province.IsDeleted,
		&city.Province.Country.Name,
		&city.Province.Country.CreatedAt,
		&city.Province.Country.CreatedBy,
		&city.Province.Country.CreatedByName,
		&city.Province.Country.UpdatedAt,
		&city.Province.Country.UpdatedBy,
		&city.Province.Country.UpdatedByName,
		&city.Province.Country.IsDeleted,
	)
	helper.PanicIfError(err)
}

func getCities(rows *sql.Rows) []domain.City {
	var cities []domain.City
	for rows.Next() {
		city := domain.City{}
		scanCity(rows, &city)
		cities = append(cities, city)
	}
	return cities
}
