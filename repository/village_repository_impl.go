package repository

import (
	"context"
	"database/sql"
	"errors"
	"minang-kos-service/constant"
	"minang-kos-service/helper"
	"minang-kos-service/model/domain"
)

type VillageRepositoryImpl struct {
}

func NewVillageRepository() VillageRepository {
	return &VillageRepositoryImpl{}
}

func (repository *VillageRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, domainModel any) any {
	village := domainModel.(domain.Village)
	result, err := tx.ExecContext(
		ctx, sqlSaveVillage(), village.Name, village.District.Id, village.CreatedAt, village.CreatedBy, village.CreatedByName,
		village.UpdatedAt, village.UpdatedBy, village.UpdatedByName, village.IsDeleted,
	)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	village.Id = id
	return village
}

func (repository *VillageRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, domainModel any) any {
	village := domainModel.(domain.Village)
	_, err := tx.ExecContext(
		ctx, sqlUpdateVillage(), village.Name, village.District.Id,
		village.UpdatedAt, village.UpdatedBy, village.UpdatedByName, village.Id,
	)
	helper.PanicIfError(err)

	return village
}

func (repository *VillageRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, domainModel any) {
	village := domainModel.(domain.Village)
	_, err := tx.ExecContext(
		ctx, sqlDeleteVillage(), village.UpdatedAt, village.UpdatedBy,
		village.UpdatedByName, village.IsDeleted, village.Id,
	)
	helper.PanicIfError(err)
}

func (repository *VillageRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int64) (any, error) {
	sqlQuery := sqlSelectVillage() + " AND mv.id = ?"
	rows, err := tx.QueryContext(ctx, sqlQuery, id)
	helper.PanicIfError(err)
	defer rows.Close()

	village := domain.Village{}
	if rows.Next() {
		scanVillage(rows, &village)
		return village, nil
	}
	return village, errors.New(constant.DATA_NOT_FOUND)
}

func (repository *VillageRepositoryImpl) FindAllWithPagination(ctx context.Context, tx *sql.Tx, searchBy map[string]any) any {
	name := searchBy["name"].(string)
	districtId := searchBy["districtId"].(int64)
	page := searchBy["page"].(int)
	size := searchBy["size"].(int)

	page = helper.GetSqlOffset(page, size)

	sqlSearch, args := sqlSearchByVillage(name, districtId)
	sqlQuery := sqlSelectVillage() + sqlSearch + " ORDER BY mv.id ASC LIMIT ? OFFSET ?"
	args = append(args, size, page)

	rows := FetchRows(ctx, tx, sqlQuery, args)
	defer rows.Close()

	return getVillages(rows)
}

func (repository *VillageRepositoryImpl) FindAllWithoutPagination(ctx context.Context, tx *sql.Tx, searchBy map[string]any) any {
	name := searchBy["name"].(string)
	districtId := searchBy["districtId"].(int64)

	sqlSearch, args := sqlSearchByVillage(name, districtId)
	sqlQuery := sqlSelectVillage() + sqlSearch + " ORDER BY mv.id ASC"

	rows := FetchRows(ctx, tx, sqlQuery, args)
	defer rows.Close()

	return getVillages(rows)
}

func (repository *VillageRepositoryImpl) FindTotalItem(ctx context.Context, tx *sql.Tx, searchBy map[string]any) int {
	name := searchBy["name"].(string)
	districtId := searchBy["districtId"].(int64)

	sqlSearch, args := sqlSearchByVillage(name, districtId)
	sqlQuery := sqlFindTotalVillage() + sqlSearch

	rows := FetchRows(ctx, tx, sqlQuery, args)
	defer rows.Close()

	return ScanTotalItem(rows)
}

func sqlSaveVillage() string {
	return "INSERT INTO m_village(name," +
		" district_id," +
		" created_at," +
		" created_by," +
		" created_by_name," +
		" updated_at," +
		" updated_by," +
		" updated_by_name," +
		" is_deleted) VALUES (?,?,?,?,?,?,?,?,?)"
}

func sqlUpdateVillage() string {
	return "UPDATE m_village SET name = ?," +
		" district_id = ?," +
		" updated_at = ?," +
		" updated_by = ?," +
		" updated_by_name = ? WHERE id = ?"
}

func sqlDeleteVillage() string {
	return "UPDATE m_village SET updated_at = ?," +
		" updated_by = ?," +
		" updated_by_name = ?," +
		" is_deleted = ?" +
		" WHERE id = ?"
}

func sqlSelectVillage() string {
	return "SELECT mv.id," +
		" mv.name," +
		" mv.district_id," +
		" mv.created_at," +
		" mv.created_by," +
		" mv.created_by_name," +
		" mv.updated_at," +
		" mv.updated_by," +
		" mv.updated_by_name," +
		" mv.is_deleted," +
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
		" mc.is_deleted FROM m_village mv" +
		" LEFT JOIN m_district md ON mv.district_id = md.id" +
		" LEFT JOIN m_city mct ON md.city_id = mct.id" +
		" LEFT JOIN m_province mp ON mct.province_id = mp.id" +
		" LEFT JOIN m_country mc ON mp.country_id = mc.id" +
		" WHERE mv.is_deleted = false"
}

func sqlFindTotalVillage() string {
	return "SELECT COUNT(1) AS totalItem FROM m_village mv WHERE mv.is_deleted = false"
}

func sqlSearchByVillage(name string, districtId int64) (string, []any) {
	var args []any
	sqlQuery := ""

	if len(name) != 0 {
		sqlQuery += " AND LOWER(mv.name) LIKE ?"
		args = append(args, helper.StringQueryLike(name))
	}

	if districtId != 0 {
		sqlQuery += " AND mv.district_id = ?"
		args = append(args, districtId)
	}
	return sqlQuery, args
}

func scanVillage(rows *sql.Rows, village *domain.Village) {
	err := rows.Scan(
		&village.Id,
		&village.Name,
		&village.District.Id,
		&village.CreatedAt,
		&village.CreatedBy,
		&village.CreatedByName,
		&village.UpdatedAt,
		&village.UpdatedBy,
		&village.UpdatedByName,
		&village.IsDeleted,
		&village.District.Name,
		&village.District.City.Id,
		&village.District.CreatedAt,
		&village.District.CreatedBy,
		&village.District.CreatedByName,
		&village.District.UpdatedAt,
		&village.District.UpdatedBy,
		&village.District.UpdatedByName,
		&village.District.IsDeleted,
		&village.District.City.Name,
		&village.District.City.Province.Id,
		&village.District.City.CreatedAt,
		&village.District.City.CreatedBy,
		&village.District.City.CreatedByName,
		&village.District.City.UpdatedAt,
		&village.District.City.UpdatedBy,
		&village.District.City.UpdatedByName,
		&village.District.City.IsDeleted,
		&village.District.City.Province.Name,
		&village.District.City.Province.Country.Id,
		&village.District.City.Province.CreatedAt,
		&village.District.City.Province.CreatedBy,
		&village.District.City.Province.CreatedByName,
		&village.District.City.Province.UpdatedAt,
		&village.District.City.Province.UpdatedBy,
		&village.District.City.Province.UpdatedByName,
		&village.District.City.Province.IsDeleted,
		&village.District.City.Province.Country.Name,
		&village.District.City.Province.Country.CreatedAt,
		&village.District.City.Province.Country.CreatedBy,
		&village.District.City.Province.Country.CreatedByName,
		&village.District.City.Province.Country.UpdatedAt,
		&village.District.City.Province.Country.UpdatedBy,
		&village.District.City.Province.Country.UpdatedByName,
		&village.District.City.Province.Country.IsDeleted,
	)
	helper.PanicIfError(err)
}

func getVillages(rows *sql.Rows) []domain.Village {
	var villages []domain.Village
	for rows.Next() {
		village := domain.Village{}
		scanVillage(rows, &village)
		villages = append(villages, village)
	}
	return villages
}
