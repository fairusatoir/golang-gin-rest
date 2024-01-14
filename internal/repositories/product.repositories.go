package repositories

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/fairusatoir/golang-gin-rest/cmd/api/config"
	"github.com/fairusatoir/golang-gin-rest/internal/models"
	"github.com/fairusatoir/golang-gin-rest/pkg/constants"
	"github.com/fairusatoir/golang-gin-rest/pkg/utils"
	"github.com/pkg/errors"
)

type ProductRepo struct {
	*config.Config
}

func NewProductRepo(cfg *config.Config) *ProductRepo {
	return &ProductRepo{cfg}
}

func (r *ProductRepo) All(ctx context.Context, tx *sql.Tx) *utils.Result {
	rs, err := tx.QueryContext(ctx, r.Q_PRODUCT_ALL)
	if err != nil {
		return utils.NewResult(nil, http.StatusInternalServerError, errors.Wrap(err, constants.ErrExecQuery.Error()))
	}
	defer rs.Close()

	var ts []models.Product
	for rs.Next() {
		var t models.Product
		if err := rs.Scan(&t.Id, &t.Name, &t.Price); err != nil {
			return utils.NewResult(nil, http.StatusInternalServerError, errors.Wrap(err, constants.ErrScanQuery.Error()))
		}
		ts = append(ts, t)
	}

	if err = rs.Err(); err != nil {
		return utils.NewResult(nil, http.StatusInternalServerError, errors.Wrap(err, constants.ErrResultQuery.Error()))
	}

	return utils.NewResult(ts, http.StatusOK, nil)
}
