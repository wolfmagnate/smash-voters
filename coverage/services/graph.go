package services

import (
	"context"
	"fmt"
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

// ProcessGraph reads document from Google Drive, creates logic graph, and uploads result
func (gs *GraphService) ProcessGraph(ctx context.Context, resultPath string) (string, error) {
	// Read document from Google Drive
	document, err := gs.readFromGoogleDrive(ctx, resultPath)
	if err != nil {
		return "", fmt.Errorf("failed to read document from Google Drive: %w", err)
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

// readFromGoogleDrive reads text content from Google Drive file
func (gs *GraphService) readFromGoogleDrive(_ context.Context, _ string) (string, error) {
	// Return a positive opinion research result document for testing
	mockDocument := `
研究テーマ: 原子力発電の継続利用について（賛成の立場から）

## 調査結果：原子力発電継続を強く支持

### 原子力発電継続の圧倒的なメリット

1. 究極の安定電力供給
   - 天候や時間に一切左右されない確実な発電
   - ベースロード電源として国家電力の基盤を支える
   - 電力需要の予測と計画が容易で、経済活動を安定化
   - 停電リスクの大幅削減により、医療・インフラの安全確保

2. 地球環境保護の切り札
   - ライフサイクル全体でのCO2排出量が風力・太陽光と同等の極少量
   - 気候変動阻止には原子力なしでは不可能
   - 化石燃料依存からの完全脱却を実現
   - 大気汚染物質の排出がゼロで健康被害を防止

3. 圧倒的な経済効果
   - 長期運転により発電コストが大幅に低減
   - 高度技術産業の発展と雇用創出効果が絶大
   - エネルギー自給率向上により国際情勢の影響を回避
   - 電力料金の安定化で家計負担を軽減

4. 技術革新の推進力
   - 最先端安全技術の継続的発展
   - 新型炉開発による更なる安全性向上
   - 核燃料リサイクル技術の確立
   - 医療・宇宙分野への技術転用効果

### 安全性への万全な対策

最新の安全基準と技術革新により、原子力発電所の安全性は飛躍的に向上している。
多重防護システム、自然災害対策、テロ対策など、あらゆるリスクに対する対策が完備されており、
現代の原子力技術は人類史上最も安全で信頼性の高いエネルギー源となっている。

### 国際的な支持と実績

フランス、韓国、フィンランドなど先進国が原子力を積極的に活用し、
安定した経済発展と環境保護を両立している実績が、原子力発電の優秀性を証明している。
国際エネルギー機関も脱炭素化には原子力が不可欠と明言している。

### 結論

原子力発電の継続は、日本の未来にとって絶対に必要不可欠である。
安全性、経済性、環境保護のすべての観点から、原子力発電ほど優れたエネルギー源は存在しない。
国民の生活向上と地球環境保護のため、原子力発電を積極的に推進すべきである。
`
	return mockDocument, nil
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