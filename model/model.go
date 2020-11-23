package model

import "gorm.io/gorm"

// Vuelo es la entidad que se va a persistir en la base de datos
type Vuelo struct {
	gorm.Model
	Name    string
	Origin  string
	Destiny string
}
