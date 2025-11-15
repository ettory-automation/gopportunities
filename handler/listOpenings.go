package handler

import (
	"net/http"

	"github.com/ettory-automation/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// ListOpeningsHandler godoc
// @Summary List all job openings
// @Description List all job openings from the database
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(c *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(c, "list-openings", openings)
}
