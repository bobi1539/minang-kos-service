package repository

import (
	"context"
	"database/sql"
	"minang-kos-service/model/domain"
)

type KosFacilityRepository interface {
	Save(ctx context.Context, tx *sql.Tx, kosFacility domain.KosFacility) domain.KosFacility
	FindAllWithoutPagination(ctx context.Context, tx *sql.Tx, searchBy map[string]any) []domain.KosFacility
}
