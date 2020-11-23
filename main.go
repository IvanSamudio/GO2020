package main

import (
	controller "github.com/IvanSamudio/GO2020/Controller/productocontroller"
	model "github.com/IvanSamudio/GO2020/model/productomodel"
	router "github.com/IvanSamudio/GO2020/router/router"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := model.NewDataBase()
	defer db.DB.Close()
	// le paso la instancia de la base de datos
	controller.NewController(&db)
	router.Start()
}
