package product

import (
	"gorm.io/gorm"
	"github.com/lib/pq"
)

type Product struct {
    gorm.Model
    Name        string
    Description string
    Images      pq.StringArray
    Price			 float64
}

func NewProduct(name string, description string, images []string, price float64) *Product {
    return &Product{
        Name: name,
        Description: description,
        Images: images,
        Price: price,
    }
}