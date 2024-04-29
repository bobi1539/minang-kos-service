package repository

import (
	"context"
	"database/sql"
)

type CrudRepository interface {
	Save(ctx context.Context, tx *sql.Tx, domainModel any) any
	Update(ctx context.Context, tx *sql.Tx, domainModel any) any
	Delete(ctx context.Context, tx *sql.Tx, domainModel any)
	FindById(ctx context.Context, tx *sql.Tx, id int64) (any, error)
	FindAllWithPagination(ctx context.Context, tx *sql.Tx, searchBy map[string]any) any
	FindAllWithoutPagination(ctx context.Context, tx *sql.Tx, searchBy map[string]any) any
	FindTotalItem(ctx context.Context, tx *sql.Tx, searchBy map[string]any) int
}
