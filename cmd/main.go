package main

import (
	"context"
	"flag"
	"os/signal"
	"syscall"
	"time"

	"github.com/SashaMelva/buffer_data_to_database/internal/app"
	"github.com/SashaMelva/buffer_data_to_database/internal/config"
	"github.com/SashaMelva/buffer_data_to_database/internal/entity"
	"github.com/SashaMelva/buffer_data_to_database/internal/logger"
	"github.com/SashaMelva/buffer_data_to_database/internal/memory/connection"
	storage "github.com/SashaMelva/buffer_data_to_database/internal/memory/storage/postgre"
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
	connectionDB := connection.New(config.DataBase, log)

	memstorage := storage.New(connectionDB.StorageDb, log)
	app := app.New(log, config.Tokens, buffer, memstorage)

	httpServer := http.NewServer(log, app, config.HttpServer)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		httpServer.Stop(ctx)
	}()
	go func() {
		log.Debug("Start listen")
		var fact *entity.Fact
		for {
			fact = buffer.Read()
			ok, err := app.SaveFact(fact)

			if err != nil {
				log.Debug(err)
				break
			}

			if !ok {
				buffer.Add(fact)
			} else {
				log.Debug("Fact Saving :", fact)
			}
		}
	}()
	log.Debug("Services is running...")
	httpServer.Start(ctx)
}
