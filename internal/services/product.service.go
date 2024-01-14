package services

import (
	"context"
	"net/http"

	"github.com/fairusatoir/golang-gin-rest/internal/clients"
	"github.com/fairusatoir/golang-gin-rest/internal/repositories"
	"github.com/fairusatoir/golang-gin-rest/pkg/constants"
	"github.com/fairusatoir/golang-gin-rest/pkg/utils"
	"github.com/pkg/errors"
)

type ProductService struct {
	*repositories.ProductRepo
	*clients.Datamaster
}

func NewProductService(rp *repositories.ProductRepo, dm *clients.Datamaster) *ProductService {
	return &ProductService{rp, dm}
}

func (ps *ProductService) FindAll(ctx context.Context) *utils.Result {
	tx, err := ps.Begin()
	if err != nil {
		return utils.NewResult(nil, http.StatusInternalServerError, errors.Wrap(err, constants.ErrTxStart.Error()))
	}
	defer utils.CommitOrCallback(tx)

	return ps.All(ctx, tx)
}
