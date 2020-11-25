package objeto

// ProductoJSON ..
type ProductoJSON struct {
	ID     int    `json:"id" param:"id"`
	Nombre string `json:"nombre" form:"nombre"`
	Precio int    `json:"precio" form:"precio"`
}
