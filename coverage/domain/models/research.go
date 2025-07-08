package models

// ResearchRequest represents the incoming request structure for research API
type ResearchRequest struct {
	Query      string `json:"query"`
	WebhookURL string `json:"webhook_url"`
}

// ExternalResearchRequest represents the request sent to external research API
type ExternalResearchRequest struct {
	Query      string `json:"query"`
	DrivePath  string `json:"drive_path"`
	WebhookURL string `json:"webhook_url"`
}

// DebateAnalysisRequest represents a request for debate graph creation
type DebateAnalysisRequest struct {
	Document string `json:"document"`
	Theme    string `json:"theme"`
}

// LogicAnalysisRequest represents a request for logic graph creation
type LogicAnalysisRequest struct {
	Document string `json:"document"`
}

// GraphRequest represents the incoming request structure for graph API
type GraphRequest struct {
	Status     string `json:"status" validate:"required,oneof=completed failed"`
	ResultPath string `json:"result_path" validate:"required"`
}

// GraphResponse represents the response structure for graph API
type GraphResponse struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	GraphPath string `json:"graph_path,omitempty"`
}