package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/wolfmagnate/smash-voters/bff/infra/db"
)

type SeedData struct {
	Elections    []ElectionSeed    `json:"elections"`
	Parties      []PartySeed       `json:"parties"`
	Questions    []QuestionSeed    `json:"questions"`
	PartyStances []PartyStanceSeed `json:"party_stances"`
}

type ElectionSeed struct {
	Name         string `json:"name"`
	ElectionDate string `json:"election_date"`
}

type PartySeed struct {
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
}

type QuestionSeed struct {
	ElectionName string `json:"election_name"`
	Title        string `json:"title"`
	QuestionText string `json:"question_text"`
	Description  string `json:"description"`
	DisplayOrder int32  `json:"display_order"`
}

type PartyStanceSeed struct {
	PartyName     string `json:"party_name"`
	QuestionTitle string `json:"question_title"`
	ElectionName  string `json:"election_name"`
	Stance        int32  `json:"stance"`
}

// ClearData truncates all relevant tables to ensure a clean slate for seeding.
func ClearData(queries *db.Queries) error {
	ctx := context.Background()
	log.Println("Clearing existing data...")
	// TRUNCATE CASCADE will also reset serial sequences
	err := queries.TruncateAllTables(ctx)
	if err != nil {
		return fmt.Errorf("failed to truncate tables: %w", err)
	}
	log.Println("Existing data cleared.")
	return nil
}

func Seed(queries *db.Queries, seedFilePath string) error {
	// Clear existing data before seeding
	if err := ClearData(queries); err != nil {
		return fmt.Errorf("failed to clear data: %w", err)
	}

	// Read seed data from JSON file
	jsonFile, err := os.Open(seedFilePath)
	if err != nil {
		return fmt.Errorf("Error opening seed_data.json: %v", err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var seedData SeedData
	json.Unmarshal(byteValue, &seedData)

	ctx := context.Background()

	// Seed Elections
	for _, election := range seedData.Elections {
		electionDate, err := time.Parse("2006-01-02", election.ElectionDate)
		if err != nil {
			log.Printf("Error parsing election date %s: %v", election.ElectionDate, err)
			continue
		}
		_, err = queries.CreateElection(ctx, db.CreateElectionParams{
			Name:         election.Name,
			ElectionDate: pgtype.Date{Time: electionDate, Valid: true},
		})
		if err != nil {
			log.Printf("Error seeding election %s: %v", election.Name, err)
		}
	}

	// Seed Parties
	for _, party := range seedData.Parties {
		_, err = queries.CreateParty(ctx, db.CreatePartyParams{
			Name:      party.Name,
			ShortName: pgtype.Text{String: party.ShortName, Valid: party.ShortName != ""},
		})
		if err != nil {
			log.Printf("Error seeding party %s: %v", party.Name, err)
		}
	}

	// Seed Questions
	for _, question := range seedData.Questions {
		electionID, err := queries.GetElectionByName(ctx, question.ElectionName)
		if err != nil {
			log.Printf("Error getting election ID for %s: %v", question.ElectionName, err)
			continue
		}
		_, err = queries.CreateQuestion(ctx, db.CreateQuestionParams{
			ElectionID:   electionID,
			Title:        question.Title,
			QuestionText: question.QuestionText,
			Description:  pgtype.Text{String: question.Description, Valid: question.Description != ""},
			DisplayOrder: question.DisplayOrder,
		})
		if err != nil {
			log.Printf("Error seeding question %s for election %s: %v", question.Title, question.ElectionName, err)
		}
	}

	// Seed Party Stances
	for _, ps := range seedData.PartyStances {
		electionID, err := queries.GetElectionByName(ctx, ps.ElectionName)
		if err != nil {
			log.Printf("Error getting election ID for %s: %v", ps.ElectionName, err)
			continue
		}

		partyID, err := queries.GetPartyByName(ctx, ps.PartyName)
		if err != nil {
			log.Printf("Error getting party ID for %s: %v", ps.PartyName, err)
			continue
		}

		questionID, err := queries.GetQuestionByTitleAndElectionID(ctx, db.GetQuestionByTitleAndElectionIDParams{
			Title:      ps.QuestionTitle,
			ElectionID: electionID,
		})
		if err != nil {
			log.Printf("Error getting question ID for %s in election %s: %v", ps.QuestionTitle, ps.ElectionName, err)
			continue
		}

		_, err = queries.CreatePartyStance(ctx, db.CreatePartyStanceParams{
			PartyID:    partyID,
			QuestionID: questionID,
			Stance:     ps.Stance,
		})
		if err != nil {
			log.Printf("Error seeding party stance for %s - %s: %v", ps.PartyName, ps.QuestionTitle, err)
		}
	}

	return nil
}
