package main

import (
	controller "GO2020/Controller/productoconntroller"
	model "GO2020/model/productoModel"
	router "GO2020/router/router"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := model.NewDataBase()
	defer db.DB.Close()
	// le paso la instancia de la base de datos
	controller.NewController(&db)
	router.Start()
}
