package productomodel

import "database/sql"

// comment

// Database ..
type Database struct {
	DB *sql.DB
}

// Producto ..
type Producto struct {
	ID     int
	Nombre string
	Precio int
}

// NewDataBase ..
func NewDataBase() Database {
	db, err := sql.Open("mysql", "root:@/producto")

	if err != nil {
		panic("Failed to connect to database!")
	}
	err = db.Ping()
	if err != nil {
		panic("error al conectar")
	}
	return Database{db}
}

// GetProductos ..
func (db *Database) GetProductos() []*Producto {
	var v []*Producto

	resultado, err := db.DB.Query("SELECT *FROM producto")

	if err != nil {
		panic(err.Error())
	}

	for resultado.Next() {
		var p Producto
		resultado.Scan(&p.ID, &p.Nombre, &p.Precio)
		v = append(v, &p)
	}
	defer resultado.Close()
	return v
}
