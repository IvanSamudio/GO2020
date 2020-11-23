package database

import (
	"github.com/IvanSamudio/GO2020/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Database ...
type Database struct {
	db *gorm.DB
}

// NewDatabase ...
func NewDatabase(dsn string) *Database { //dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return &Database{db}
}

// FindAll devuelve todos los vuelos registrados
func (database *Database) FindAll() []model.Vuelo {
	var vuelos []model.Vuelo
	database.db.Find(&vuelos)
	return vuelos
}

// CreateSchema create el esquema a partir de las entidades.
func (database *Database) CreateSchema() {
	database.db.AutoMigrate(&model.Vuelo{})
}

// Add inserta un vuelo
func (database *Database) Add(v *model.Vuelo) {
	database.db.Create(v)
}

// FindByID ...
func (database *Database) FindByID(ID uint) *model.Vuelo {
	var v model.Vuelo
	database.db.Find(&v, ID)
	return &v
}
