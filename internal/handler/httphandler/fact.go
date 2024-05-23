package httphandler

import (
	"net/http"

	"github.com/SashaMelva/buffer_data_to_database/internal/entity"
	"github.com/gin-gonic/gin"
)

func (s *Service) AddFacts(ctx *gin.Context) {
	var facts entity.Fact
	if err := ctx.ShouldBindJSON(&facts); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	s.log.Debug(facts)

	id, err := s.app.CreateFact(&facts)

	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Id": id})
}
