package repository

import (
	"context"
	"database/sql"
)

type CrudRepository interface {
	Save(ctx context.Context, tx *sql.Tx, domainModel any) any
}
