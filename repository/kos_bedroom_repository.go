package repository

import (
	"context"
	"database/sql"
)

type KosBedroomRepository interface {
	CrudRepository
	FindAll(ctx context.Context, tx *sql.Tx, searchBy any) any
}
