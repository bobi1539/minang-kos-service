package repository

import (
	"context"
	"database/sql"
	"minang-kos-service/helper"
	"minang-kos-service/model/domain"
)

type KosBedroomRepositoryImpl struct {
}

func NewKosBedroomRepository() KosBedroomRepository {
	return &KosBedroomRepositoryImpl{}
}

func (repository *KosBedroomRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, domainModel any) any {
	kosBedroom := domainModel.(domain.KosBedroom)
	result, err := tx.ExecContext(ctx, sqlSaveKosBedroom(), argsSaveKosBedroom(kosBedroom)...)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	kosBedroom.Id = id
	return kosBedroom
}

func (repository *KosBedroomRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, domainModel any) any {
	panic("imp")
}

func (repository *KosBedroomRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, domainModel any) {
	panic("imp")
}

func (repository *KosBedroomRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int64) (any, error) {
	panic("imp")
}

func (repository *KosBedroomRepositoryImpl) FindAllWithPagination(ctx context.Context, tx *sql.Tx, searchBy map[string]any) any {
	panic("imp")
}

func (repository *KosBedroomRepositoryImpl) FindAllWithoutPagination(ctx context.Context, tx *sql.Tx, searchBy map[string]any) any {
	panic("imp")
}

func (repository *KosBedroomRepositoryImpl) FindTotalItem(ctx context.Context, tx *sql.Tx, searchBy map[string]any) int {
	panic("imp")
}

func sqlSaveKosBedroom() string {
	return "INSERT INTO m_kos_bedroom(" +
		" title," +
		" room_length," +
		" room_width," +
		" unit_length," +
		" is_include_electricity," +
		" price," +
		" street," +
		" images," +
		" kos_type_id," +
		" village_id," +
		" user_id," +
		" created_at," +
		" created_by," +
		" created_by_name," +
		" updated_at," +
		" updated_by," +
		" updated_by_name," +
		" is_deleted) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
}

func argsSaveKosBedroom(kosBedroom domain.KosBedroom) []any {
	var args []any
	args = append(args, kosBedroom.Title)
	args = append(args, kosBedroom.RoomLength)
	args = append(args, kosBedroom.RoomWidth)
	args = append(args, kosBedroom.UnitLength)
	args = append(args, kosBedroom.IsIncludeElectricity)
	args = append(args, kosBedroom.Price)
	args = append(args, kosBedroom.Street)
	args = append(args, kosBedroom.Images)
	args = append(args, kosBedroom.KosType.Id)
	args = append(args, kosBedroom.Village.Id)
	args = append(args, kosBedroom.User.Id)
	args = append(args, kosBedroom.CreatedAt)
	args = append(args, kosBedroom.CreatedBy)
	args = append(args, kosBedroom.CreatedByName)
	args = append(args, kosBedroom.UpdatedAt)
	args = append(args, kosBedroom.UpdatedBy)
	args = append(args, kosBedroom.UpdatedByName)
	args = append(args, kosBedroom.IsDeleted)
	return args
}
