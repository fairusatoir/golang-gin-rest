package routers

import (
	"github.com/fairusatoir/golang-gin-rest/internal/controllers/handlers"
	"github.com/gin-gonic/gin"
)

type ProductRouter struct {
	*handlers.ProductHandler
	*gin.RouterGroup
}

func NewProductRouter(ph *handlers.ProductHandler, g *gin.RouterGroup) *ProductRouter {
	return &ProductRouter{ph, g}
}

func (pr *ProductRouter) AddRoutes() {
	rg := pr.Group("/product")
	rg.GET("/all", pr.GetAllProduct)
}
