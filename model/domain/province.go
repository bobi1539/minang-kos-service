package domain

type Province struct {
	Id      int64
	Name    string
	Country Country
	BaseDomain
}
