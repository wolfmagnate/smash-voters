package handler

import (
	"context"
	"math"
	"net/http"
	"sort"

	"github.com/labstack/echo/v4"
	"github.com/wolfmagnate/smash-voters/bff/infra/db"
)

type MatchingHandler struct {
	Queries *db.Queries
}

func NewMatchingHandler(q *db.Queries) *MatchingHandler {
	return &MatchingHandler{Queries: q}
}

type UserAnswer struct {
	QuestionID int32 `json:"question_id"`
	Answer     int32 `json:"answer"`
}

type MatchRequest struct {
	Answers           []UserAnswer `json:"answers"`
	ImportantQuestionIDs []int32      `json:"important_question_ids"`
}

type PartyMatchResult struct {
	PartyName string `json:"party_name"`
	MatchRate int32  `json:"match_rate"`
}

type MatchResponse struct {
	TopMatch PartyMatchResult   `json:"top_match"`
	Results  []PartyMatchResult `json:"results"`
}

func (h *MatchingHandler) CalculateMatch(c echo.Context) error {
	electionID, err := parseID(c, "election_id")
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	var req MatchRequest
	if err := c.Bind(&req); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	// Get all parties
	parties, err := h.Queries.GetAllParties(context.Background())
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	// Get all party stances for the given election
	partyStances, err := h.Queries.GetPartyStancesByElectionID(context.Background(), int32(electionID))
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	// Map party stances for easier lookup
	stanceMap := make(map[int32]map[int32]int32) // questionID -> partyID -> stance
	partyNames := make(map[int32]string)         // partyID -> partyName
	for _, ps := range partyStances {
		if _, ok := stanceMap[ps.QuestionID]; !ok {
			stanceMap[ps.QuestionID] = make(map[int32]int32)
		}
		stanceMap[ps.QuestionID][ps.PartyID] = ps.Stance
		partyNames[ps.PartyID] = ps.PartyName
	}

	var results []PartyMatchResult

	for _, party := range parties {
		var totalDistance float64
		var maxDistance float64

		for _, userAnswer := range req.Answers {
			partyStance, ok := stanceMap[userAnswer.QuestionID][party.ID]
			if !ok {
				// If a party doesn't have a stance for a question, skip it for this calculation
				continue
			}

			distance := math.Abs(float64(userAnswer.Answer - partyStance))

			weight := 1.0
			for _, importantQID := range req.ImportantQuestionIDs {
				if importantQID == userAnswer.QuestionID {
					weight = 2.0
					break
				}
			}

			totalDistance += distance * weight
			maxDistance += 4.0 * weight // Max distance between -2 and 2 is 4
		}

		matchRate := 0.0
		if maxDistance > 0 {
			matchRate = (1 - (totalDistance / maxDistance)) * 100
		}

		results = append(results, PartyMatchResult{
			PartyName: partyNames[party.ID],
			MatchRate: int32(math.Max(0, math.Min(100, matchRate))), // Ensure rate is between 0 and 100
		})
	}

	// Sort results by match rate in descending order
	sort.Slice(results, func(i, j int) bool {
		return results[i].MatchRate > results[j].MatchRate
	})

	var topMatch PartyMatchResult
	if len(results) > 0 {
		topMatch = results[0]
	}

	return c.JSON(http.StatusOK, MatchResponse{
		TopMatch: topMatch,
		Results:  results,
	})
}
