package product

import "github.com/lib/pq"

type ProductCreateRequest struct {
	Name 				string `json:"name" validate:"required"`
  Description string `json:"description" validate:"required"`
  Images 			pq.StringArray `json:"images" validate:"required,dive,url"`
	Price				float64 `json:"price" validate:"required,gt=0"`
}

type ProductUpdateRequest struct {
	Name 				string `json:"name,omitempty"`
  Description string `json:"description,omitempty"`
  Images 			pq.StringArray `json:"images,omitempty" validate:"dive,url"`
	Price				float64 `json:"price,omitempty" validate:"gt=0"`
}
