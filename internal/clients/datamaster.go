package clients

import (
	"database/sql"

	"github.com/fairusatoir/golang-gin-rest/cmd/api/config"
	"github.com/fairusatoir/golang-gin-rest/pkg/constants"
	"github.com/fairusatoir/golang-gin-rest/pkg/log"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

type Datamaster struct {
	*sql.DB
}

func Connect(cfg *config.Config) (*Datamaster, error) {
	db, err := sql.Open(cfg.DM_Driver, cfg.DM_URL)
	if err != nil {
		return nil, errors.Wrap(err, constants.ErrConnDb.Error())
	}

	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, constants.ErrPingDb.Error())
	}

	return &Datamaster{db}, nil
}

func InitDatamaster() {
	cfg, _ := config.Load()
	_, err := Connect(cfg)
	if err != nil {
		log.Fatal(err.Error(), logrus.Fields{log.Category: log.Database})
	}
	log.Info(constants.InfoDbConfig, logrus.Fields{log.Category: log.Database})
}
