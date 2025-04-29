package product

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Images      pq.StringArray `json:"images" gorm:"type:text[]"`
}
type AllProduct struct {
	AllProd []Product `json:"allproduct"`
}

func NewProduct(name, des string, images []string) *Product {
	return &Product{
		Name:        name,
		Description: des,
		Images:      pq.StringArray(images),
	}
}
