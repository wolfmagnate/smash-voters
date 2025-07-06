package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"github.com/wolfmagnate/smash-voters/bff/handler"
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

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	// Initialize Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Serve OpenAPI spec and Swagger UI
	e.File("/openapi.yml", "handler/openapi.yml") // Serve the OpenAPI spec file directly
	e.GET("/swagger/*", echo.WrapHandler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+port+"/openapi.yml"), // Point Swagger UI to the direct spec file
	)))

	// Handlers
	electionHandler := handler.NewElectionHandler(queries)
	matchingHandler := handler.NewMatchingHandler(queries)

	// Routes
	e.GET("/elections/latest", electionHandler.GetLatestElection)
	e.GET("/elections/:election_id/questions", electionHandler.GetQuestionsByElectionID)
	e.POST("/elections/:election_id/matches", matchingHandler.CalculateMatch)

	log.Printf("Server starting on port %s", port)
	log.Fatal(e.Start(":" + port))
}
