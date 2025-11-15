package handler

import (
	"net/http"

	"github.com/ettory-automation/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// CreateOpeningHandler godoc
// @Summary Create a new job opening
// @Description Creates a new job opening record in the database
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Opening data"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(c *gin.Context) {
	request := CreateOpeningRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role: request.Role,
		Company: request.Company,
		Location: request.Location,
		Remote: *request.Remote,
		Link: request.Link,
		Salary: request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	sendSuccess(c, "create-opening", opening)
}