package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wolfmagnate/smash-voters/coverage/domain/models"
	"github.com/wolfmagnate/smash-voters/coverage/services"
)

// GraphHandler handles graph-related HTTP requests
type GraphHandler struct {
	graphService *services.GraphService
}

// NewGraphHandler creates a new GraphHandler instance
func NewGraphHandler(graphService *services.GraphService) *GraphHandler {
	return &GraphHandler{
		graphService: graphService,
	}
}

// HandleGraph handles the GET /graph endpoint with query parameters
func (gh *GraphHandler) HandleGraph(c echo.Context) error {
	ctx := c.Request().Context()

	// Extract query parameters
	status := c.QueryParam("status")
	resultPath := c.QueryParam("result_path")

	// Validate query parameters
	if status == "" {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "status query parameter is required",
			Code:  http.StatusBadRequest,
		})
	}

	if resultPath == "" {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "result_path query parameter is required",
			Code:  http.StatusBadRequest,
		})
	}

	if status != "completed" && status != "failed" {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "status must be either 'completed' or 'failed'",
			Code:  http.StatusBadRequest,
		})
	}

	// If status is failed, return early
	if status == "failed" {
		return c.JSON(http.StatusOK, models.GraphResponse{
			Status:  "failed",
			Message: "Research process failed, no graph generated",
		})
	}

	// Process graph creation request
	graphPath, err := gh.graphService.ProcessGraph(ctx, resultPath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Failed to process graph creation",
			Details: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, models.GraphResponse{
		Status:    "success",
		Message:   "Logic graph created successfully",
		GraphPath: graphPath,
	})
}