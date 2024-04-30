package repository

import (
	"context"
	"database/sql"
	"minang-kos-service/helper"
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

func ScanTotalItem(rows *sql.Rows) int {
	var totalItem int
	for rows.Next() {
		err := rows.Scan(&totalItem)
		helper.PanicIfError(err)
	}
	return totalItem
}

func FetchRows(ctx context.Context, tx *sql.Tx, sqlQuery string, args []any) *sql.Rows {
	rows, err := tx.QueryContext(ctx, sqlQuery, args...)
	helper.PanicIfError(err)
	return rows
}
