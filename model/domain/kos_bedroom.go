package domain

type KosBedroom struct {
	Id                   int64
	Title                string
	RoomLength           float32
	RoomWidth            float32
	UnitLength           string
	IsIncludeElectricity bool
	Price                float32
	Images               string
	KosType              KosType
	Village              Village
	User                 User
	BaseDomain
}
