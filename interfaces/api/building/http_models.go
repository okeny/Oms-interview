package building

type Request struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name" binding:"required"`
	Address string `json:"address,omitempty"`
}
