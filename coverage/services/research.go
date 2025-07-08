package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/wolfmagnate/smash-voters/coverage/domain/models"
	"net/http"
	"time"
)

// ResearchService handles research-related business logic
type ResearchService struct {
	externalAPIURL string
	httpClient     *http.Client
}

// NewResearchService creates a new ResearchService instance
func NewResearchService(externalAPIURL string) *ResearchService {
	return &ResearchService{
		externalAPIURL: externalAPIURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// ProcessResearch processes the research request and calls external API
func (rs *ResearchService) ProcessResearch(ctx context.Context, req *models.ResearchRequest, theme, isPositive string) (*http.Response, error) {
	// Create external API request
	externalReq := models.ExternalResearchRequest{
		Query:      req.Query,
		DrivePath:  fmt.Sprintf("/%s/%s", theme, isPositive),
		WebhookURL: req.WebhookURL,
	}

	// Marshal request to JSON
	reqBody, err := json.Marshal(externalReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal external request: %w", err)
	}

	// Create HTTP request with context
	httpReq, err := http.NewRequestWithContext(ctx, "POST", rs.externalAPIURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	// Send request to external API
	resp, err := rs.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to call external API: %w", err)
	}

	return resp, nil
}
