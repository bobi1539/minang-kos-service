package repository

import (
	"context"
	"database/sql"
	"minang-kos-service/model/domain"
	"minang-kos-service/model/web/search"
)

type KosFacilityRepository interface {
	Save(ctx context.Context, tx *sql.Tx, kosFacility domain.KosFacility) domain.KosFacility
	FindAll(ctx context.Context, tx *sql.Tx, searchBy search.KosFacilitySearch) []domain.KosFacility
}
