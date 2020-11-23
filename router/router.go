package router

import (
	controller "github.com/IvanSamudio/GO2020/Controller/productocontroller"
	"github.com/gin-gonic/gin"
)

// Start ..
func Start() {
	r := gin.Default()
	r.GET("/getproductos", controller.GetVuelos)
	r.Run()

}
