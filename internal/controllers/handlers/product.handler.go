package handlers

import (
	"net/http"

	"github.com/fairusatoir/golang-gin-rest/internal/services"
	"github.com/fairusatoir/golang-gin-rest/pkg/log"
	"github.com/fairusatoir/golang-gin-rest/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ProductHandler struct {
	*services.ProductService
}

func NewProductHandler(ps *services.ProductService) *ProductHandler {
	return &ProductHandler{ps}
}

func (th *ProductHandler) GetAllProduct(ctx *gin.Context) {
	rs := th.FindAll(ctx)
	if rs.Err != nil {
		log.Error(rs.Err.Error(), logrus.Fields{log.Category: log.Config})
		utils.NewErrorResponse(ctx, rs.StatusCode, rs.Err.Error())
		return
	}

	utils.NewSuccessResponse(ctx, rs.StatusCode, http.StatusText(rs.StatusCode), rs.Data)
}
