package apartment

type ApartmentRequest struct {
	ID         int    `json:"id,omitempty"`
	BuildingID int    `json:"building_id" binding:"required"`
	Number     string `json:"number" binding:"required"`
	Floor      int    `json:"floor,omitempty"`
	SQMeters   int    `json:"sq_meters,omitempty"`
}
