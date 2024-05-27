package httphandler

import (
	"net/http"

	"github.com/SashaMelva/buffer_data_to_database/internal/entity"
	"github.com/gin-gonic/gin"
)

// @Summary Add Facts
// @Tags facts
// @Description add facts for buffer
// @ID add-facts
// @Accept  json
// @Produce  json
// @Param input body entity.Facts true "facts info"
// @Success 200
// @Failure 400,404
// @Failure 500
// @Failure default
// @Router /add_new_facts_array [post]
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
