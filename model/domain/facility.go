package domain

type Facility struct {
	Id           int64
	Name         string
	FacilityType FacilityType
	BaseDomain
}
