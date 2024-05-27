package app

import (
	"github.com/SashaMelva/buffer_data_to_database/internal/config"
	storage "github.com/SashaMelva/buffer_data_to_database/internal/memory/storage/postgre"
	"github.com/SashaMelva/buffer_data_to_database/pkg/buffer"
	"go.uber.org/zap"
)

type App struct {
	log     *zap.SugaredLogger
	buf     *buffer.Buffer
	Tokens  *config.Tokens
	Storage *storage.Storage
}

func New(logger *zap.SugaredLogger, config *config.Tokens, buf *buffer.Buffer, storage *storage.Storage) *App {
	return &App{
		log:     logger,
		buf:     buf,
		Tokens:  config,
		Storage: storage,
	}
}
