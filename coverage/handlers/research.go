package handlers

import (
	"bytes"
	"github.com/wolfmagnate/smash-voters/coverage/domain/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wolfmagnate/smash-voters/coverage/services"
)

// ResearchHandler handles research-related HTTP requests
type ResearchHandler struct {
	researchService *services.ResearchService
}

// NewResearchHandler creates a new ResearchHandler instance
func NewResearchHandler(researchService *services.ResearchService) *ResearchHandler {
	return &ResearchHandler{
		researchService: researchService,
	}
}

// HandleResearch handles the POST /research/{theme}/{is_positive} endpoint
func (rh *ResearchHandler) HandleResearch(c echo.Context) error {
	ctx := c.Request().Context()

	// Extract path parameters
	theme := c.Param("theme")
	isPositive := c.Param("is_positive")

	// Parse request body
	var req models.ResearchRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid request body",
			Code:  http.StatusBadRequest,
		})
	}

	// Validate input
	if req.Query == "" {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Query is required",
			Code:  http.StatusBadRequest,
		})
	}

	// Process research request
	resp, err := rh.researchService.ProcessResearch(ctx, &req, theme, isPositive)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Failed to process research request",
			Details: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}
	defer resp.Body.Close()

	// Read response body
	var respBody bytes.Buffer
	_, err = respBody.ReadFrom(resp.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to read external API response",
			Code:  http.StatusInternalServerError,
		})
	}

	// Return the external API response with the same status code
	return c.JSONBlob(resp.StatusCode, respBody.Bytes())
}
