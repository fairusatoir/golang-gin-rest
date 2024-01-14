package utils

import (
	"database/sql"

	"github.com/fairusatoir/golang-gin-rest/pkg/constants"
	"github.com/pkg/errors"
)

func CommitOrCallback(tx *sql.Tx) error {
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return errors.Wrap(err, constants.ErrTxCommit.Error())
	}
	return nil
}
