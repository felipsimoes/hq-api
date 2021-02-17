package collection

import "gorm.io/gorm"

// Collection is a struct that stores a list of HQs
type Collection struct {
	gorm.Model
	// ID        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Publisher string `json:"published"`
}
