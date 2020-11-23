package productocontroller

import (
	productModel "github.com/IvanSamudio/GO2020/model/productomodel"
	"net/http"

	"github.com/gin-gonic/gin"
)

var modelo *productModel.Database

// NewController ..
func NewController(model *productModel.Database) {
	modelo = model
}

// GetProductos ..
func GetProductos(ctx *gin.Context) {
	p := modelo.GetProductos()
	ctx.JSON(http.StatusCreated, gin.H{
		"Object": p,
	})
}
