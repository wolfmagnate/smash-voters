package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wolfmagnate/smash-voters/bff/infra/db"
)

type ElectionHandler struct {
	Queries *db.Queries
}

func NewElectionHandler(q *db.Queries) *ElectionHandler {
	return &ElectionHandler{Queries: q}
}

func (h *ElectionHandler) GetLatestElection(c echo.Context) error {
	election, err := h.Queries.GetLatestElection(context.Background())
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, election)
}

func (h *ElectionHandler) GetQuestionsByElectionID(c echo.Context) error {
	electionID, err := parseID(c, "election_id")
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	questions, err := h.Queries.GetQuestionsByElectionID(context.Background(), int32(electionID))
	if err != nil || len(questions) == 0 {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"questions": questions})
}
