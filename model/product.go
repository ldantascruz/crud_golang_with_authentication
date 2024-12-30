package model

type Product struct {
	ID    string  `json:"id_product"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
