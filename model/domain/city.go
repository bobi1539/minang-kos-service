package domain

type City struct {
	Id       int64
	Name     string
	Province Province
	BaseDomain
}
