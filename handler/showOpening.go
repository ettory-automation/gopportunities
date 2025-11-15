package handler

import (
	"net/http"

	"github.com/ettory-automation/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// ShowOpeningHandler godoc
// @Summary Show a job opening
// @Description Show an existing job opening from the database
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query int true "Opening identification"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
func ShowOpeningHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(c, http.StatusNotFound, "opening not found")
		return
	}

	sendSuccess(c, "show-opening", opening)
}
