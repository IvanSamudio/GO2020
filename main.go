package main

import (
	controller "github.com/IvanSamudio/GO2020/Controller/productocontroller"
	modelo "github.com/IvanSamudio/GO2020/model/productomodel"
	"github.com/IvanSamudio/GO2020/router"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := modelo.NewDataBase()
	defer db.DB.Close()
	// le paso la instancia de la base de datos
	controller.NewController(&db)
	router.Start()
}
