package constants

import (
	"errors"
)

var (
	// Config
	ErrLoadConfig      = errors.New("failed to load config")
	ErrUnmarshalConfig = errors.New("failed to unmarshal config")

	// Database
	ErrConnDb  = errors.New("failed to open connection database")
	ErrPingDb  = errors.New("failed to pinging database")
	ErrCloseDb = errors.New("failed to close database")

	// Repository
	ErrExecQuery   = errors.New("failed to run query")
	ErrScanQuery   = errors.New("failed to fetch data from database")
	ErrResultQuery = errors.New("failed to read query results")

	// Transaction
	ErrTxStart  = errors.New("failed to start database transaction")
	ErrTxCommit = errors.New("failed to commit database transaction")
)
