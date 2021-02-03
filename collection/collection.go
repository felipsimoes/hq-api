package collection

// Collection is a struct that stores a list of HQs
type Collection struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Publisher string `json:"published"`
}
