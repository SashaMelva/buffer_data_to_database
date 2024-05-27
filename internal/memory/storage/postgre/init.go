package storage

import (
	"database/sql"
	"sync"

	"go.uber.org/zap"
)

type Storage struct {
	log          *zap.SugaredLogger
	ConnectionDB *sql.DB
	sync.RWMutex
}

func New(connection *sql.DB, log *zap.SugaredLogger) *Storage {
	return &Storage{
		log:          log,
		ConnectionDB: connection,
	}
}
