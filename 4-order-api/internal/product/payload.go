package product

import "github.com/lib/pq"

type ProductCreateRequest struct {
	Url 				string `json:"url" validate:"required,url"`
	Name 				string `json:"name" validate:"required,string"`
  Description string `json:"description" validate:"required,string"`
  Images 			pq.StringArray `json:"images" validate:"required,dive,url"`
	Price				float64 `json:"price" validate:"required,gt=0"`
}

type ProductUpdateRequest struct {
	Url 				string `json:"url,omitempty" validate:"url"`
	Name 				string `json:"name,omitempty" validate:"string"`
  Description string `json:"description,omitempty" validate:"string"`
  Images 			pq.StringArray `json:"images,omitempty" validate:"dive,url"`
	Price				float64 `json:"price,omitempty" validate:"gt=0"`
}
