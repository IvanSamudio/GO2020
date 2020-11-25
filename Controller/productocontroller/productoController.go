package productocontroller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	productoModel "github.com/IvanSamudio/GO2020/model/productomodel"
	"github.com/IvanSamudio/GO2020/objeto"
	"github.com/gin-gonic/gin"
)

var modelo *productoModel.Database

// NewController ..
func NewController(model *productoModel.Database) {
	modelo = model
}

// GetProductos ..
func GetProductos(ctx *gin.Context) {
	p := modelo.GetProductos()
	ctx.JSON(http.StatusCreated, gin.H{
		"Object": p,
	})
}

// GetProducto ..
func GetProducto(ctx *gin.Context) {
	id := ctx.Param("id")
	var valor int
	valor, err := strconv.Atoi(id)
	if err != nil {
		panic("No es un integer")
	}
	var param objeto.ProductoJSON
	ctx.BindJSON(&param)
	p := modelo.GetProducto(valor)
	ctx.JSON(http.StatusOK, gin.H{
		"Object": p,
	})
}

// UpdateProducto ..
func UpdateProducto(c *gin.Context) {
	student := objeto.ProductoJSON{}
	err := c.BindJSON(&student)
	id := c.Param("id")
	var valor int
	valor, err = strconv.Atoi(id)
	if err != nil {
		panic("No es un integer")
	}
	modelo.UpdateProducto(valor, student.Precio, student.Nombre)
	exception := ""
	if err != nil {
		exception = err.Error()
	}
	c.JSON(200, gin.H{"exception": exception, "data": student})
}

// DeleteProducto ..
func DeleteProducto(contexto *gin.Context) {
	id := contexto.Param("id")
	var valor int
	valor, err := strconv.Atoi(id)
	if err != nil {
		panic("No es un integer")
	}

	fmt.Println(valor)
	modelo.DeleteProducto(valor)
	contexto.JSON(http.StatusCreated, gin.H{
		"Objeto borrado": valor,
	})
}

// InsertProducto ..
func InsertProducto(c *gin.Context) {
	id := c.Request.GetBody
	log.Printf("asdas", id)
}
