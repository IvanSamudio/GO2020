package router

import (
	controller "GO2020/Controller/productoconntroller"
	"github.com/gin-gonic/gin"
)

// Start ..
func Start() {
	r := gin.Default()
	r.GET("/getproductos", controller.GetVuelos)
	r.Run()

}
