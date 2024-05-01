package domain

type Village struct {
	Id       int64
	Name     string
	District District
	BaseDomain
}
