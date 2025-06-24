package models

// ResearchRequest represents the incoming request structure for research API
type ResearchRequest struct {
	Query string `json:"query"`
}

// ExternalResearchRequest represents the request sent to external research API
type ExternalResearchRequest struct {
	Query     string `json:"query"`
	DrivePath string `json:"drive_path"`
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