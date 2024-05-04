package request

type KosBedroomCreateRequest struct {
	Title                string   `validate:"required,max=255,min=1" json:"title"`
	RoomLength           float32  `validate:"required" json:"roomLength"`
	RoomWidth            float32  `validate:"required" json:"roomWidth"`
	IsIncludeElectricity bool     `validate:"required" json:"isIncludeElectricity"`
	Price                float32  `validate:"required" json:"price"`
	Street               string   `validate:"required" json:"street"`
	KosTypeId            int64    `validate:"required" json:"kosTypeId"`
	VillageId            int64    `validate:"required" json:"villageId"`
	FacilityIds          []int64  `validate:"required" json:"facilityIds"`
	Images               []string `validate:"required" json:"images"`
}
