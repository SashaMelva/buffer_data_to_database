package httphandler

import (
	"net/http"

	"github.com/SashaMelva/buffer_data_to_database/internal/entity"
	"github.com/gin-gonic/gin"
)

func (s *Service) AddFactsArray(ctx *gin.Context) {
	var facts entity.Facts
	if err := ctx.ShouldBindJSON(&facts); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	s.log.Debug(facts)

	err := s.app.CreateFacts(&facts)

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}
