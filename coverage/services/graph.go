package services

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/wolfmagnate/smash-voters/coverage/debate_graph_creator"
	"github.com/wolfmagnate/smash-voters/coverage/domain"
	"github.com/wolfmagnate/smash-voters/coverage/logic_graph_creator"
)

// GraphService handles graph creation and local file operations
type GraphService struct {
	logicGraphCreator *logic_graph_creator.LogicGraphCreator
	debateGraphCreator *debate_graph_creator.DebateGraphCreator
}

// NewGraphService creates a new GraphService instance
func NewGraphService() *GraphService {
	// Create components using their proper Create functions
	basicStructureAnalyzer, err := logic_graph_creator.CreateBasicStructureAnalyzer()
	if err != nil {
		panic(fmt.Sprintf("Failed to create BasicStructureAnalyzer: %v", err))
	}

	impactAnalyzer, err := logic_graph_creator.CreateImpactAnalyzer()
	if err != nil {
		panic(fmt.Sprintf("Failed to create ImpactAnalyzer: %v", err))
	}

	benefitHarmConverter, err := logic_graph_creator.CreateBenefitHarmConverter()
	if err != nil {
		panic(fmt.Sprintf("Failed to create BenefitHarmConverter: %v", err))
	}

	causeFinder, err := logic_graph_creator.CreateCauseFinder()
	if err != nil {
		panic(fmt.Sprintf("Failed to create CauseFinder: %v", err))
	}

	newArgumentFinder, err := logic_graph_creator.CreateNewArgumentFinder()
	if err != nil {
		panic(fmt.Sprintf("Failed to create NewArgumentFinder: %v", err))
	}

	debateAnnotationCreator, err := debate_graph_creator.CreateDebateAnnotationCreator()
	if err != nil {
		panic(fmt.Sprintf("Failed to create DebateAnnotationCreator: %v", err))
	}

	documentSplitter, err := debate_graph_creator.CreateDocumentSplitter()
	if err != nil {
		panic(fmt.Sprintf("Failed to create DocumentSplitter: %v", err))
	}

	return &GraphService{
		logicGraphCreator: &logic_graph_creator.LogicGraphCreator{
			BasicStructureAnalyzer: basicStructureAnalyzer,
			ImpactAnalyzer:         impactAnalyzer,
			BenefitHarmConverter:   benefitHarmConverter,
			LogicGraphCompleter: &logic_graph_creator.LogicGraphCompleter{
				CauseFinder:       causeFinder,
				NewArgumentFinder: newArgumentFinder,
			},
		},
		debateGraphCreator: &debate_graph_creator.DebateGraphCreator{
			DebateAnnotationCreator: debateAnnotationCreator,
			DocumentSplitter:        documentSplitter,
		},
	}
}

// ProcessGraph reads document from deployed .txt file, creates logic graph, and saves result
func (gs *GraphService) ProcessGraph(ctx context.Context, resultPath string) (string, error) {
	// Read document from deployed .txt file
	document, err := gs.readFromDeployedFile(ctx, resultPath)
	if err != nil {
		return "", fmt.Errorf("failed to read document from deployed file: %w", err)
	}

	// Create logic graph from document
	logicGraph, err := gs.createLogicGraph(ctx, document)
	if err != nil {
		return "", fmt.Errorf("failed to create logic graph: %w", err)
	}

	// Create debate graph from logic graph
	debateGraph, err := gs.createDebateGraph(ctx, document, logicGraph)
	if err != nil {
		return "", fmt.Errorf("failed to create debate graph: %w", err)
	}

	// Convert debate graph to JSON using domain ToJSON method
	graphJSON, err := debateGraph.ToJSON()
	if err != nil {
		return "", fmt.Errorf("failed to convert debate graph to JSON: %w", err)
	}

	// Save graph JSON to local file
	graphPath, err := gs.saveGraphToFile(ctx, []byte(graphJSON), resultPath)
	if err != nil {
		return "", fmt.Errorf("failed to save graph to file: %w", err)
	}

	return graphPath, nil
}

// readFromDeployedFile reads text content from deployed .txt file URL
func (gs *GraphService) readFromDeployedFile(ctx context.Context, resultPath string) (string, error) {
	// Create HTTP request with context
	req, err := http.NewRequestWithContext(ctx, "GET", resultPath, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Make HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch file from URL: %w", err)
	}
	defer resp.Body.Close()

	// Check HTTP status
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP request failed with status: %d", resp.StatusCode)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(body), nil
}

// createLogicGraph creates a logic graph using the real LogicGraphCreator
func (gs *GraphService) createLogicGraph(ctx context.Context, document string) (*domain.LogicGraph, error) {
	// Use the actual LogicGraphCreator to create the graph
	return gs.logicGraphCreator.CreateLogicGraph(ctx, document)
}

// createDebateGraph creates a debate graph using the real DebateGraphCreator
func (gs *GraphService) createDebateGraph(ctx context.Context, document string, logicGraph *domain.LogicGraph) (*domain.DebateGraph, error) {
	// Use the actual DebateGraphCreator to create the debate graph from logic graph
	return gs.debateGraphCreator.CreateDebateGraph(ctx, document, logicGraph)
}

// saveGraphToFile saves JSON content to a local file
func (gs *GraphService) saveGraphToFile(_ context.Context, jsonContent []byte, originalPath string) (string, error) {
	// Generate output file path
	graphPath := generateGraphPath(originalPath)
	
	// Create output directory if it doesn't exist
	outputDir := filepath.Dir(graphPath)
	if outputDir != "." {
		err := os.MkdirAll(outputDir, 0755)
		if err != nil {
			return "", fmt.Errorf("failed to create output directory: %w", err)
		}
	}
	
	// Write JSON content to file
	err := os.WriteFile(graphPath, jsonContent, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to write graph file: %w", err)
	}
	
	return graphPath, nil
}


// generateGraphPath generates output path for graph JSON file
func generateGraphPath(originalPath string) string {
	// Generate local file path based on original path
	// Create a safe filename from the path
	safeName := strings.ReplaceAll(strings.Trim(originalPath, "/"), "/", "_")
	
	// If the path has a file extension, replace it with .json
	if strings.Contains(safeName, ".") {
		ext := filepath.Ext(safeName)
		safeName = strings.TrimSuffix(safeName, ext)
	}
	
	// Add timestamp to make it unique
	timestamp := time.Now().Format("20060102_150405")
	return fmt.Sprintf("output/%s_debate_graph_%s.json", safeName, timestamp)
}