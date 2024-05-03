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

func sqlSaveKosFacility() string {
	return "INSERT INTO m_kos_facility(" +
		" kos_bedroom_id," +
		" facility_id) VALUES (?,?)"
}

func argsSaveKosFacility(kosFacility domain.KosFacility) []any {
	var args []any
	args = append(args, kosFacility.KosBedroom.Id)
	args = append(args, kosFacility.Facility.Id)
	return args
}
