package building

type BuildingRequest struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name" binding:"required"`
	Address string `json:"address,omitempty"`
}
