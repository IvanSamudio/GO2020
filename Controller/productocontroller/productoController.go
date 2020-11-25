package productocontroller

import (
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
	var id = convertirValorEntero(ctx.Param("id"))
	p := modelo.GetProducto(id)
	ctx.JSON(http.StatusOK, gin.H{
		"Object": p,
	})
}

// UpdateProducto ..
func UpdateProducto(c *gin.Context) {
	var producto objeto.ProductoJSON
	producto.Nombre = c.Request.FormValue("nombre")
	producto.ID = convertirValorEntero(c.Param("id"))
	producto.Precio = convertirValorEntero(c.Request.FormValue("precio"))
	c.BindJSON(&producto)
	modelo.UpdateProducto(producto.ID, producto.Precio, producto.Nombre)
	c.JSON(http.StatusCreated, gin.H{
		"Objeto editado": &producto,
	})
}

// DeleteProducto ..
func DeleteProducto(contexto *gin.Context) {
	var ID = convertirValorEntero(contexto.Param("id"))
	modelo.DeleteProducto(ID)
	contexto.JSON(http.StatusCreated, gin.H{
		"Objeto borrado": ID,
	})
}

// InsertProducto ...
func InsertProducto(c *gin.Context) {
	var producto objeto.ProductoJSON
	producto.Nombre = c.Request.FormValue("nombre")
	producto.Precio = convertirValorEntero(c.Request.FormValue("precio"))
	c.BindJSON(&producto)
	modelo.InsertProducto(producto.Precio, producto.Nombre)
	producto.ID = modelo.GetLastID()
	c.JSON(http.StatusCreated, gin.H{
		"Objeto Creado": &producto,
	})

}

func convertirValorEntero(convertir string) int {
	var valor int
	valor, _ = strconv.Atoi(convertir)
	return valor
}
