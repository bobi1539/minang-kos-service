package domain

type Facility struct {
	Id           int64
	Name         string
	Icon         string
	FacilityType FacilityType
	BaseDomain
}
