package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/wolfmagnate/smash-voters/bff/infra"
	"github.com/wolfmagnate/smash-voters/bff/infra/db"
)

func TestGetLatestElection(t *testing.T) {
	

	// Initialize database connection
	pool, err := infra.NewPgxPool()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer pool.Close()

	queries := db.New(pool)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/elections/latest", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := NewElectionHandler(queries)

	// Assertions
	if assert.NoError(t, handler.GetLatestElection(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var election db.GetLatestElectionRow
		err := json.Unmarshal(rec.Body.Bytes(), &election)
		assert.NoError(t, err)

		// Assuming seed_data.json has "2024年衆議院議員選挙" as the latest election
		assert.Equal(t, "2024年衆議院議員選挙", election.Name)
	}
}

func TestGetQuestionsByElectionID(t *testing.T) {
	

	// Initialize database connection
	pool, err := infra.NewPgxPool()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer pool.Close()

	queries := db.New(pool)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/elections/1/questions", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set path parameter
	c.SetPath("/elections/:election_id/questions")
	c.SetParamNames("election_id")
	c.SetParamValues("1") // Assuming election ID 1 exists from seed data

	handler := NewElectionHandler(queries)

	// Assertions
	if assert.NoError(t, handler.GetQuestionsByElectionID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response struct {
			Questions []db.GetQuestionsByElectionIDRow `json:"questions"`
		}
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.NotEmpty(t, response.Questions)
		assert.Equal(t, "現金給付", response.Questions[0].Title) // Assuming "現金給付" is the first question for election ID 1
	}
}
