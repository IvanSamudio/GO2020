package productomodel

import (
	"database/sql"
	"log"
	// comment
)

// Database ..
type Database struct {
	DB *sql.DB
}

// Producto ..
type Producto struct {
	ID     int    `json:"id_producto"`
	Nombre string `json:"nombre"`
	Precio int    `json:"precio"`
}

// NewDataBase ..
func NewDataBase() Database {
	db, err := sql.Open("mysql", "root:@/productos")
	if err != nil {
		panic("Failed to connect to database!")
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return Database{db}
}

// GetProductos ..
func (db *Database) GetProductos() []*Producto {
	var arrp []*Producto

	resultado, err := db.DB.Query("SELECT *FROM producto")
	log.Printf("pepito", resultado)
	if err != nil {
		panic(err.Error())
	}

	for resultado.Next() {
		var p Producto
		resultado.Scan(&p.ID, &p.Nombre, &p.Precio)
		arrp = append(arrp, &p)
	}
	defer resultado.Close()
	return arrp
}

//func (db *Database) getProducto(id int) *Producto {
//	var v *Producto
//	selecte, err := db.DB.Query("SELECT *FROM vuelo WHERE id_producto = ?", id)
//	if err != nil {
//		panic(err.Error())
//	}
//	v = append(selecte)
//	defer selecte.close()
//	return selecte
//}
//GetProducto ...
func (db *Database) GetProducto(id int) []*Producto {
	var arrp []*Producto
	resultado, err := db.DB.Query("SELECT * FROM producto WHERE id_producto=?", id)
	if err != nil {
		panic(err.Error())
	}

	for resultado.Next() {
		var p Producto
		resultado.Scan(&p.ID, &p.Nombre, &p.Precio)
		arrp = append(arrp, &p)
	}
	defer resultado.Close()
	return arrp
}

// UpdateProducto ..
func (db *Database) UpdateProducto(id, precio int, nombre string) {
	log.Printf("id", id)
	log.Printf("nombre", nombre)
	log.Printf("precio", precio)
	actualizar, err := db.DB.Query("UPDATE  producto SET nombre = ?, precio=? WHERE id_producto = ?", nombre, precio, id)
	if err != nil {
		panic(err.Error())
	}
	defer actualizar.Close()
}

// DeleteProducto ..
func (db *Database) DeleteProducto(id int) {
	delete, err := db.DB.Query("DELETE FROM producto WHERE id_producto = ?", id)
	if err != nil {
		panic(err.Error())
	}
	defer delete.Close()
}

// InsertProducto ..
func (db *Database) InsertProducto(id, precio int, nombre string) {
	log.Printf("id", id)
	log.Printf("nombre", nombre)
	log.Printf("precio", precio)
	actualizar, err := db.DB.Query("INSERT INTO `producto`(`id_producto`, `nombre`, `precio`) VALUES (?,?,?)", id, nombre, precio)
	if err != nil {
		panic(err.Error())
	}
	defer actualizar.Close()
}
