package repository

import (
	"context"
	"database/sql"
	"errors"
	"minang-kos-service/constant"
	"minang-kos-service/helper"
	"minang-kos-service/model/domain"
	"minang-kos-service/model/web/search"
)

type KosTypeRepositorImpl struct {
}

func NewKosTypeRepository() KosTypeRepository {
	return &KosTypeRepositorImpl{}
}

func (repository *KosTypeRepositorImpl) Save(ctx context.Context, tx *sql.Tx, domainModel any) any {
	kosType := domainModel.(domain.KosType)
	result, err := tx.ExecContext(
		ctx, sqlSaveKosType(), kosType.Name, kosType.CreatedAt, kosType.CreatedBy, kosType.CreatedByName,
		kosType.UpdatedAt, kosType.UpdatedBy, kosType.UpdatedByName, kosType.IsDeleted,
	)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	kosType.Id = id
	return kosType
}

func (repository *KosTypeRepositorImpl) Update(ctx context.Context, tx *sql.Tx, domainModel any) any {
	kosType := domainModel.(domain.KosType)
	_, err := tx.ExecContext(ctx, sqlUpdateKosType(), kosType.Name, kosType.UpdatedAt, kosType.UpdatedBy, kosType.UpdatedByName, kosType.Id)
	helper.PanicIfError(err)

	return kosType
}

func (repository *KosTypeRepositorImpl) Delete(ctx context.Context, tx *sql.Tx, domainModel any) {
	kosType := domainModel.(domain.KosType)
	_, err := tx.ExecContext(
		ctx, sqlDeleteKosType(), kosType.UpdatedAt, kosType.UpdatedBy, kosType.UpdatedByName, kosType.IsDeleted, kosType.Id,
	)
	helper.PanicIfError(err)
}

func (repository *KosTypeRepositorImpl) FindById(ctx context.Context, tx *sql.Tx, id int64) (any, error) {
	sqlQuery := sqlSelectKosType() + " AND id = ?"
	rows, err := tx.QueryContext(ctx, sqlQuery, id)
	helper.PanicIfError(err)
	defer rows.Close()

	kosType := domain.KosType{}

	if rows.Next() {
		scanKosType(rows, &kosType)
		return kosType, nil
	}
	return kosType, errors.New(constant.DATA_NOT_FOUND)
}

func (repository *KosTypeRepositorImpl) FindAll(ctx context.Context, tx *sql.Tx, searchBy any) any {
	kosTypeSearch := searchBy.(search.KosTypeSearch)

	sqlSearch, args := sqlSearchKosTypeBy(kosTypeSearch.Name)
	sqlQuery := sqlSelectKosType() + sqlSearch

	if kosTypeSearch.Page > 0 {
		sqlQuery += " LIMIT ? OFFSET ?"
		offset := helper.GetSqlOffset(kosTypeSearch.Page, kosTypeSearch.Size)
		args = append(args, kosTypeSearch.Size, offset)
	}

	rows := FetchRows(ctx, tx, sqlQuery, args)
	defer rows.Close()

	return getKosTypes(rows)
}

func (repository *KosTypeRepositorImpl) FindTotalItem(ctx context.Context, tx *sql.Tx, searchBy any) int {
	kosTypeSearch := searchBy.(search.KosTypeSearch)

	sqlSearch, args := sqlSearchKosTypeBy(kosTypeSearch.Name)
	sqlQuery := sqlFindTotalKosType() + sqlSearch

	rows := FetchRows(ctx, tx, sqlQuery, args)
	defer rows.Close()

	return ScanTotalItem(rows)
}

func sqlSaveKosType() string {
	return "INSERT INTO m_kos_type(name," +
		" created_at," +
		" created_by," +
		" created_by_name," +
		" updated_at," +
		" updated_by," +
		" updated_by_name," +
		" is_deleted) VALUES (?,?,?,?,?,?,?,?)"
}

func sqlUpdateKosType() string {
	return "UPDATE m_kos_type SET name = ?," +
		" updated_at = ?," +
		" updated_by = ?," +
		" updated_by_name = ?" +
		" WHERE id = ?"
}

func sqlDeleteKosType() string {
	return "UPDATE m_kos_type SET updated_at = ?," +
		" updated_by = ?," +
		" updated_by_name = ?," +
		" is_deleted = ?" +
		" WHERE id = ?"
}

func sqlSelectKosType() string {
	return "SELECT id," +
		" name," +
		" created_at," +
		" created_by," +
		" created_by_name," +
		" updated_at," +
		" updated_by," +
		" updated_by_name," +
		" is_deleted FROM m_kos_type" +
		" WHERE is_deleted = false"
}

func sqlFindTotalKosType() string {
	return "SELECT COUNT(1) AS totalItem FROM m_kos_type WHERE is_deleted = false"
}

func sqlSearchKosTypeBy(name string) (string, []any) {
	var args []any
	sqlQuery := ""

	if len(name) != 0 {
		sqlQuery += " AND LOWER(name) LIKE ?"
		args = append(args, helper.StringQueryLike(name))
	}

	sqlQuery += " ORDER BY id ASC"
	return sqlQuery, args
}

func getKosTypes(rows *sql.Rows) []domain.KosType {
	var kosTypes []domain.KosType
	for rows.Next() {
		kosType := domain.KosType{}
		scanKosType(rows, &kosType)
		kosTypes = append(kosTypes, kosType)
	}
	return kosTypes
}

func scanKosType(rows *sql.Rows, kosType *domain.KosType) {
	err := rows.Scan(
		&kosType.Id,
		&kosType.Name,
		&kosType.CreatedAt,
		&kosType.CreatedBy,
		&kosType.CreatedByName,
		&kosType.UpdatedAt,
		&kosType.UpdatedBy,
		&kosType.UpdatedByName,
		&kosType.IsDeleted,
	)
	helper.PanicIfError(err)
}
