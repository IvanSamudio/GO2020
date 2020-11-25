package router

import (
	controller "github.com/IvanSamudio/GO2020/Controller/productocontroller"
	"github.com/gin-gonic/gin"
)

// Start ..
func Start() {
	r := gin.Default()
	r.GET("/getproductos", controller.GetProductos)
	r.PUT("/updateproducto/:id", controller.UpdateProducto)
	r.GET("/getproducto/:id", controller.GetProducto)
	r.DELETE("/deleteproducto/:id", controller.DeleteProducto)
	r.POST("/insertproducto/", controller.InsertProducto)
	r.Run()
}
