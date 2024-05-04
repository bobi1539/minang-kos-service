package repository

import (
	"context"
	"database/sql"
	"errors"
	"minang-kos-service/constant"
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
	sqlQuery := sqlSelectKosBedroom() + " AND mkb.id = ?"
	rows, err := tx.QueryContext(ctx, sqlQuery, id)
	helper.PanicIfError(err)
	defer rows.Close()

	kosBedroom := domain.KosBedroom{}
	for rows.Next() {
		scanKosBedroom(rows, &kosBedroom)
		return kosBedroom, nil
	}
	return nil, errors.New(constant.DATA_NOT_FOUND)
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

func sqlSelectKosBedroom() string {
	return "SELECT mkb.id," +
		" mkb.title," +
		" mkb.room_length," +
		" mkb.room_width," +
		" mkb.unit_length," +
		" mkb.is_include_electricity," +
		" mkb.price," +
		" mkb.street," +
		" mkb.images," +
		" mkb.kos_type_id," +
		" mkb.village_id," +
		" mkb.user_id," +
		" mkb.created_at," +
		" mkb.created_by," +
		" mkb.created_by_name," +
		" mkb.updated_at," +
		" mkb.updated_by," +
		" mkb.updated_by_name," +
		" mkb.is_deleted," +
		" mkt.name," +
		" mu.email," +
		" mu.name," +
		" mu.phone_number," +
		" mv.name," +
		" md.name," +
		" mct.name," +
		" mp.name," +
		" mc.name" +
		" FROM m_kos_bedroom mkb" +
		" LEFT JOIN m_kos_type mkt ON mkb.kos_type_id = mkt.id" +
		" LEFT JOIN m_user mu ON mkb.user_id = mu.id" +
		" LEFT JOIN m_village mv ON mkb.village_id = mv.id" +
		" LEFT JOIN m_district md ON mv.district_id = md.id" +
		" LEFT JOIN m_city mct ON md.city_id = mct.id" +
		" LEFT JOIN m_province mp ON mct.province_id = mp.id" +
		" LEFT JOIN m_country mc ON mp.country_id = mc.id" +
		" WHERE mkb.is_deleted = false"
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

func scanKosBedroom(rows *sql.Rows, kosBedroom *domain.KosBedroom) {
	err := rows.Scan(
		&kosBedroom.Id,
		&kosBedroom.Title,
		&kosBedroom.RoomLength,
		&kosBedroom.RoomWidth,
		&kosBedroom.UnitLength,
		&kosBedroom.IsIncludeElectricity,
		&kosBedroom.Price,
		&kosBedroom.Street,
		&kosBedroom.Images,
		&kosBedroom.KosType.Id,
		&kosBedroom.Village.Id,
		&kosBedroom.User.Id,
		&kosBedroom.CreatedAt,
		&kosBedroom.CreatedBy,
		&kosBedroom.CreatedByName,
		&kosBedroom.UpdatedAt,
		&kosBedroom.UpdatedBy,
		&kosBedroom.UpdatedByName,
		&kosBedroom.IsDeleted,
		&kosBedroom.KosType.Name,
		&kosBedroom.User.Email,
		&kosBedroom.User.Name,
		&kosBedroom.User.PhoneNumber,
		&kosBedroom.Village.Name,
		&kosBedroom.Village.District.Name,
		&kosBedroom.Village.District.City.Name,
		&kosBedroom.Village.District.City.Province.Name,
		&kosBedroom.Village.District.City.Province.Country.Name,
	)
	helper.PanicIfError(err)
}
