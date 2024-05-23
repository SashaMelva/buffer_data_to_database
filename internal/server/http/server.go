package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/SashaMelva/buffer_data_to_database/internal/app"
	"github.com/SashaMelva/buffer_data_to_database/internal/config"
	"github.com/SashaMelva/buffer_data_to_database/internal/handler/httphandler"
	"github.com/SashaMelva/buffer_data_to_database/pkg"
)

type Server struct {
	srv *http.Server
	log *zap.SugaredLogger
}

func NewServer(log *zap.SugaredLogger, app *app.App, config *config.ConfigHttpServer) *Server {
	log.Debug("URL: " + config.Host + ":" + config.Port)
	router := gin.Default()
	handler := httphandler.NewHendler(log, app)

	router.GET("/", func(ctx *gin.Context) {
		fmt.Println("Hellow world)")
	})

	router.POST("/add_new_facts", handler.AddFacts)

	return &Server{
		srv: &http.Server{
			Addr:    config.Host + ":" + config.Port,
			Handler: router,
		},
		log: log,
	}
}

func (s *Server) Start(ctx context.Context) {
	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.log.Fatalf("listen: %s\n", err)
	}
}

func (s *Server) Stop(ctx context.Context) {
	if err := s.srv.Shutdown(ctx); err != nil {
		s.log.Fatal("Server forced to shutdown: ", err)
	}

	os.Exit(1)
}

func AuthMiddleware(log *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		log.Debug(token)

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		sub, err := pkg.ParseAccessToken(token, "secretJWT")

		log.Debug(sub)
		if err != nil {
			log.Error(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		intAcc, err := strconv.Atoi(sub)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			log.Error(err)
			return
		}

		c.Set("accountId", intAcc)
		c.Next()
	}
}
