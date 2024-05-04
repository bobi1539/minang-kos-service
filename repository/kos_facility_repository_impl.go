package repository

import (
	"context"
	"database/sql"
	"minang-kos-service/helper"
	"minang-kos-service/model/domain"
)

type KosFacilityRepositoryImpl struct {
}

func NewKosFacilityRepository() KosFacilityRepository {
	return &KosFacilityRepositoryImpl{}
}

func (repository *KosFacilityRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, kosFacility domain.KosFacility) domain.KosFacility {
	result, err := tx.ExecContext(ctx, sqlSaveKosFacility(), argsSaveKosFacility(kosFacility)...)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	kosFacility.Id = id
	return kosFacility
}

func (repository *KosFacilityRepositoryImpl) FindAllWithoutPagination(ctx context.Context, tx *sql.Tx, searchBy map[string]any) []domain.KosFacility {
	kosBedroomId := searchBy["kosBedroomId"].(int64)
	facilityId := searchBy["facilityId"].(int64)

	sqlSearch, args := sqlSearchByKosFacility(kosBedroomId, facilityId)
	sqlQuery := sqlSelectKosFacility() + sqlSearch

	rows := FetchRows(ctx, tx, sqlQuery, args)
	defer rows.Close()

	return getKosFacilities(rows)
}

func sqlSaveKosFacility() string {
	return "INSERT INTO m_kos_facility(" +
		" kos_bedroom_id," +
		" facility_id) VALUES (?,?)"
}

func sqlSelectKosFacility() string {
	return "SELECT mkf.id," +
		" mkf.kos_bedroom_id," +
		" mkf.facility_id," +
		" mf.name," +
		" mf.icon," +
		" mf.facility_type_id," +
		" mft.name" +
		" FROM m_kos_facility mkf" +
		" LEFT JOIN m_facility mf ON mkf.facility_id = mf.id" +
		" LEFT JOIN m_facility_type mft ON mf.facility_type_id = mft.id" +
		" WHERE 1 = 1"
}

func sqlSearchByKosFacility(kosBedroomId int64, facilityId int64) (string, []any) {
	var args []any
	sqlQuery := ""

	if kosBedroomId != 0 {
		sqlQuery += " AND mkf.kos_bedroom_id =  ?"
		args = append(args, kosBedroomId)
	}

	if facilityId != 0 {
		sqlQuery += " AND mkf.facility_id = ?"
		args = append(args, facilityId)
	}
	return sqlQuery, args
}

func argsSaveKosFacility(kosFacility domain.KosFacility) []any {
	var args []any
	args = append(args, kosFacility.KosBedroom.Id)
	args = append(args, kosFacility.Facility.Id)
	return args
}

func getKosFacilities(rows *sql.Rows) []domain.KosFacility {
	var kosFacilities []domain.KosFacility
	for rows.Next() {
		kosFacility := domain.KosFacility{}
		scanKosFacility(rows, &kosFacility)
		kosFacilities = append(kosFacilities, kosFacility)
	}
	return kosFacilities
}

func scanKosFacility(rows *sql.Rows, kosFacility *domain.KosFacility) {
	err := rows.Scan(
		&kosFacility.Id,
		&kosFacility.KosBedroom.Id,
		&kosFacility.Facility.Id,
		&kosFacility.Facility.Name,
		&kosFacility.Facility.Icon,
		&kosFacility.Facility.FacilityType.Id,
		&kosFacility.Facility.FacilityType.Name,
	)
	helper.PanicIfError(err)
}
