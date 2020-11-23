package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/IvanSamudio/GO2020/controller"
	"github.com/IvanSamudio/GO2020/database"
	"github.com/IvanSamudio/GO2020/model"
	"github.com/IvanSamudio/GO2020/service"
)

func main() {
	// Creo una instancia de la base de datos
	db := database.NewDatabase("database.db") // aca esta hardcoded el nombre de la base de datos
	// Creo el esquema
	db.CreateSchema()

	// creo el servicio y le paso la base de datos
	service := service.NewInstance(db)

	// ------- ESTO ES A MODO DE EJEMPLO DE COMO USAR EL SERVICIO -----------
	// genero un vuelo a insertar cuyo nombre es un numero
	vuelo := model.Vuelo{Name: strconv.Itoa(time.Now().Nanosecond()), Origin: "Ezeiza", Destiny: "Chicago"}

	// agrego el vuelo
	service.Add(&vuelo)

	// pongo en la variable vuelo el contenido del punntero que me devuelve la funcion findByID
	vuelo = *service.FindByID(vuelo.ID)
	//      ^
	//      |
	//      Esto me devuelve el valor que aloja el puntero que me devuelve la
	//      funcion service.FindByID.

	fmt.Printf("\tID=%v\n", vuelo.ID)
	fmt.Printf("\tName=%v\n", vuelo.Name)
	fmt.Printf("\tOrigin=%v\n", vuelo.Origin)
	fmt.Printf("\tDestiby=%v\n", vuelo.Destiny)
	fmt.Printf("\tCreatedAt=%v\n", vuelo.CreatedAt)
	fmt.Printf("\tUpdatedAt=%v\n", vuelo.UpdatedAt)
	fmt.Printf("\tDeletedAt=%v\n", vuelo.DeletedAt)
	// -------- ACA TERMINA EL EJEMPLO DE COMO USAR EL SERVICIO -------

	// Ahora voy a crear el servicio rest (http)
	http := controller.NewHTTPController(&service)
	http.Run()
}
