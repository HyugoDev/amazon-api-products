package models

//Product este es el modelo que usa el product o producto
type Product struct {
	Category string `json:"Category"`
	Type     string `json:"Type"`
	Brand    string `json:"Brand"`
	Item     string `json:"Item"`
}
