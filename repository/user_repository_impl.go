package repository

import (
	"context"
	"database/sql"
	"errors"
	"minang-kos-service/constant"
	"minang-kos-service/helper"
	"minang-kos-service/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	sqlQuery := sqlSelectUser() + " AND mu.email = ?"

	rows, err := tx.QueryContext(ctx, sqlQuery, email)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		scanUser(rows, &user)
		return user, nil
	}
	return user, errors.New(constant.DATA_NOT_FOUND)
}

func sqlSelectUser() string {
	return "SELECT mu.id," +
		" mu.email," +
		" mu.password," +
		" mu.name," +
		" mu.phone_number," +
		" mu.role_id," +
		" mu.created_at," +
		" mu.created_by," +
		" mu.created_by_name," +
		" mu.updated_at," +
		" mu.updated_by," +
		" mu.updated_by_name," +
		" mu.is_deleted," +
		" mr.name," +
		" mr.created_at," +
		" mr.created_by," +
		" mr.created_by_name," +
		" mr.updated_at," +
		" mr.updated_by," +
		" mr.updated_by_name," +
		" mr.is_deleted FROM m_user mu" +
		" LEFT JOIN m_role mr ON mu.role_id = mr.id" +
		" WHERE mu.is_deleted = false"
}

func scanUser(rows *sql.Rows, user *domain.User) {
	err := rows.Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.PhoneNumber,
		&user.Role.Id,
		&user.CreatedAt,
		&user.CreatedBy,
		&user.CreatedByName,
		&user.UpdatedAt,
		&user.UpdatedBy,
		&user.UpdatedByName,
		&user.IsDeleted,
		&user.Role.Name,
		&user.Role.CreatedAt,
		&user.Role.CreatedBy,
		&user.Role.CreatedByName,
		&user.Role.UpdatedAt,
		&user.Role.UpdatedBy,
		&user.Role.UpdatedByName,
		&user.Role.IsDeleted,
	)
	helper.PanicIfError(err)
}
