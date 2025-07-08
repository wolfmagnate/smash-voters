package main

import (
	"log"
	"os"

	"github.com/wolfmagnate/smash-voters/bff/infra"
	"github.com/wolfmagnate/smash-voters/bff/infra/db"
)

func main() {
	// Initialize database connection
	pool, err := infra.NewPgxPool()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer pool.Close()

	queries := db.New(pool)

	seedFilePath := "seed_data.json"
	if len(os.Args) > 1 {
		seedFilePath = os.Args[1]
	}

	if err := infra.Seed(queries, seedFilePath); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}

	log.Println("Seeding completed.")
}
