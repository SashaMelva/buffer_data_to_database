package main

import (
	"context"
	"flag"
	"os/signal"
	"syscall"
	"time"

	"github.com/SashaMelva/buffer_data_to_database/internal/app"
	"github.com/SashaMelva/buffer_data_to_database/internal/config"
	"github.com/SashaMelva/buffer_data_to_database/internal/logger"
	"github.com/SashaMelva/buffer_data_to_database/internal/server/http"
	"github.com/SashaMelva/buffer_data_to_database/pkg/buffer"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "../configs/", "Path to configuration file")
}

func main() {

	config := config.New(configFile)
	log := logger.New(config.Logger, "../logs/")

	buffer := buffer.New()
	app := app.New(log, config.Tokens, buffer)

	httpServer := http.NewServer(log, app, config.HttpServer)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		fact := <-buffer.Buf
		log.Debug("get ", fact)
	}()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		httpServer.Stop(ctx)
	}()
}
