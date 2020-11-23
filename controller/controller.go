package controller

import (
	"net/http"
	"strconv"

	"github.com/IvanSamudio/GO2020/service"
	"github.com/gin-gonic/gin"
)

// HTTPController ...
type HTTPController struct {
	router  *gin.Engine
	service *service.Vuelo
}

// NewHTTPController ...
func NewHTTPController(s *service.Vuelo) HTTPController {
	r := gin.Default()
	makeEndpoints(r, s) // registro los endpoints
	return HTTPController{r, s}
}

// Run ejecuta el controller
func (c *HTTPController) Run() {
	c.router.Run()
}

func makeEndpoints(r *gin.Engine, s *service.Vuelo) {
	r.GET("/vuelos/:id", makeFindHandler(s))
	r.GET("/vuelos/", makeFindHandler(s))
}

func makeFindHandler(s *service.Vuelo) gin.HandlerFunc {
	// Fijate que aca devuelvo una funcion y no un valor
	return func(c *gin.Context) {
		i, _ := strconv.Atoi(c.Param("id"))

		v := (*s).FindByID(uint(i))

		c.JSON(http.StatusOK, gin.H{
			"vuelo": v,
		})
	}
}
