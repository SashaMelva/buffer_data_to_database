package app

import (
	"github.com/SashaMelva/buffer_data_to_database/internal/config"
	"github.com/SashaMelva/buffer_data_to_database/pkg/buffer"
	"go.uber.org/zap"
)

type App struct {
	log    *zap.SugaredLogger
	buf    *buffer.Buffer
	Tokens *config.Tokens
}

func New(logger *zap.SugaredLogger, config *config.Tokens, buf *buffer.Buffer) *App {
	return &App{
		log:    logger,
		buf:    buf,
		Tokens: config,
	}
}
