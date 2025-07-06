package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/wolfmagnate/smash-voters/bff/infra"
	"github.com/wolfmagnate/smash-voters/bff/infra/db"
)

func TestCalculateMatch(t *testing.T) {
	// Initialize database connection
	pool, err := infra.NewPgxPool()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer pool.Close()

	queries := db.New(pool)

	// Create a sample MatchRequest
	// Assuming election ID 1 and question IDs 1, 2, 3 exist from seed data
	matchReq := MatchRequest{
		Answers: []UserAnswer{
			{QuestionID: 1, Answer: 2},   // 現金給付: 賛成
			{QuestionID: 2, Answer: -1},  // 消費税0%: やや反対
			{QuestionID: 3, Answer: 0},   // 大企業課税強化: 中立
		},
		ImportantQuestionIDs: []int32{1, 3}, // 現金給付と大企業課税強化を重要視
	}
	reqBody, _ := json.Marshal(matchReq)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/elections/1/matches", bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set path parameter
	c.SetPath("/elections/:election_id/matches")
	c.SetParamNames("election_id")
	c.SetParamValues("1") // Assuming election ID 1 exists from seed data

	handler := NewMatchingHandler(queries)

	// Assertions
	if assert.NoError(t, handler.CalculateMatch(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response MatchResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.NotEmpty(t, response.Results)
		assert.NotEmpty(t, response.TopMatch.PartyName)
		assert.True(t, response.TopMatch.MatchRate >= 0 && response.TopMatch.MatchRate <= 100)

		// You can add more specific assertions here based on expected match rates
		// For example, if you know a specific party should have a high match rate with these answers:
		// assert.Equal(t, "公明党", response.TopMatch.PartyName) // Example, replace with actual expected party
	}
}
