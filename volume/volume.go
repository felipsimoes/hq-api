package volume

// Volume is a struct that represents one HQ
type Volume struct {
	ID           int    `json:"id"`
	CollectionID int    `json:"collection_id"`
	Name         string `json:"name"`
	Edition      string `json:"edition"`
	Pages        int    `json:"pages"`
}
