package handler

import (
	"fmt"
	"net/http"

	"github.com/ettory-automation/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// DeleteOpeningHandler godoc
// @Summary Delete a new job opening
// @Description Deletes an existing job opening from the database
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query int true "Opening identification"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	//Find Opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(c, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	//Delete Opening
	if err := db.Delete(&opening).Error; err != nil {
		sendError(c, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}

	sendSuccess(c, "delete-opening", opening)
}
