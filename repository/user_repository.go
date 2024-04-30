package repository

import (
	"context"
	"database/sql"
	"minang-kos-service/model/domain"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
}
