package dto

type CreateProductrRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
