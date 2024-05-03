package response

type KosBedroomResponse struct {
	Id                   int64           `json:"id"`
	Title                string          `json:"title"`
	RoomLength           float32         `json:"roomLength"`
	RoomWidth            float32         `json:"roomWidth"`
	UnitLength           string          `json:"unitLength"`
	IsIncludeElectricity bool            `json:"isIncludeElectricity"`
	Price                float32         `json:"price"`
	Images               []string        `json:"images"`
	KosType              KosTypeResponse `json:"kosType"`
	Village              VillageResponse `json:"village"`
	User                 UserResponse    `json:"user"`
	BaseDomainResponse
}
