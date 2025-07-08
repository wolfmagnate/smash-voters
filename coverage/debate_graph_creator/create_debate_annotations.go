package debate_graph_creator

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"log"
	"strings"
	"text/template"

	"github.com/wolfmagnate/smash-voters/coverage/domain"
	"github.com/wolfmagnate/smash-voters/coverage/infra"
)

//go:embed create_debate_annotations_prompt.md
var creteDebateAnnotationsPromptMarkdown string

type DebateAnnotationCreator struct {
	tmpl *template.Template
}

func CreateDebateAnnotationCreator() (*DebateAnnotationCreator, error) {
	tmpl, err := template.New("prompt").Parse(creteDebateAnnotationsPromptMarkdown)

	if err != nil {
		return nil, fmt.Errorf("起動時のテンプレート解析に失敗しました: %w", err)
	}

	return &DebateAnnotationCreator{tmpl: tmpl}, nil
}

type CreateDebateAnnotationTemplateData struct {
	Document        string
	TargetParagraph string
	LogicGraphNodes string
	LogicGraphEdges string
	CiteDocument    string
}

type LogicAnnotations struct {
	Annotations []LogicAnnotation `json:"annotations"` // 分析対象の段落に含まれる全ての論理構造グラフの要素の分析結果
}

type LogicAnnotation struct {
	TargetType     string         `json:"target_type"`     // "node"または"edge"のいずれか
	TargetText     string         `json:"target_text"`     // 分析対象の段落のうち、このアノテーションを行う根拠となる部分
	NodeAnnotation NodeAnnotation `json:"node_annotation"` // TargetTypeが"node"のときのみ有効
	EdgeAnnotation EdgeAnnotation `json:"edge_annotation"` // TargetTypeが"edge"のときのみ有効
}

type NodeAnnotation struct {
	AnnotationType     string            `json:"annotation_type"`               // "argument"または"importance"または"uniqueness"または"importance_rebuttal"または"uniqueness_rebuttal"のいずれか
	Argument           string            `json:"argument"`                      // アノテーションを行う対象の論理構造グラフのノード
	Importance         *domain.Assertion `json:"importance,omitempty"`          // なぜArgumentが重要であるかの理由を表す文章。AnnotationTypeが"importance"のときのみ有効
	Uniqueness         *domain.Assertion `json:"uniqueness,omitempty"`          // なぜArgumentがStatus QuoまたはAffirmative Planでのみ発生するのかの理由を表す文章。AnnotationTypeが"uniqueness"のときのみ有効
	ImportanceRebuttal *domain.Assertion `json:"importance_rebuttal,omitempty"` // なぜArgumentが重要ではないのかの理由を表す文章。AnnotationTypeが"importance_rebuttal"のときのみ有効
	UniquenessRebuttal *domain.Assertion `json:"uniqueness_rebuttal,omitempty"` // なぜArgumentがStatus QuoとAffirmative Planの両方で発生してしまうかの理由を表す文章。AnnotationTypeが
}

type EdgeAnnotation struct {
	AnnotationType     string            `json:"annotation_type"`               // "certainty"または"uniqueness"または"certainty_rebuttal"または"uniqueness_rebuttal"のいずれか
	CauseArgument      string            `json:"cause_argument"`                // エッジの原因に対応する論理構造グラフのノード
	EffectArgument     string            `json:"effect_argument"`               // エッジの結果に対応する論理構造グラフのノード
	Certainty          *domain.Assertion `json:"certainty,omitempty"`           // なぜCauseArgumentがEffectArgumentを引き起こす可能性が高いのかの理由を表す文章。AnnotationTypeが"certainty"のときのみ有効
	Uniqueness         *domain.Assertion `json:"uniqueness,omitempty"`          // なぜCauseArgumentがStatus QuoまたはAffirmative Planでのみ発生するのかの理由を表す文章。AnnotationTypeが"uniqueness"のときのみ有効
	CertaintyRebuttal  *domain.Assertion `json:"certainty_rebuttal,omitempty"`  // なぜCauseArgumentがEffectArgumentを発生させる可能性が低いのかを表す文章。AnnotationTypeが"certainty_rebuttal"のときのみ有効
	UniquenessRebuttal *domain.Assertion `json:"uniqueness_rebuttal,omitempty"` // なぜCauseArgumentがStatus QuoとAffirmative Planの両方でEffectArgumentを引き起こすのかの理由を表す文章。AnnotationTypeが"uniqueness_rebuttal"のときのみ有効
}

func (analyzer *DebateAnnotationCreator) CreateDebateAnnotations(ctx context.Context, document string, targetParagraph string, logicGraph *domain.LogicGraph, citeDocument string) (*LogicAnnotations, error) {
	nodes := make([]string, 0)
	for _, node := range logicGraph.Nodes {
		nodes = append(nodes, node.Argument)
	}

	causalRelationships := domain.ListAllCausalRelationships(logicGraph)

	data := CreateDebateAnnotationTemplateData{
		Document:        document,
		TargetParagraph: targetParagraph,
		LogicGraphNodes: strings.Join(nodes, "\n"),
		LogicGraphEdges: strings.Join(causalRelationships, "\n"),
		CiteDocument:    citeDocument,
	}

	var processedPrompt bytes.Buffer
	err := analyzer.tmpl.Execute(&processedPrompt, data)
	if err != nil {
		log.Printf("テンプレートの実行に失敗しました: %v", err)
		return nil, fmt.Errorf("テンプレートの実行に失敗しました: %w", err)
	}

	promptString := processedPrompt.String()

	thinkingBudget := int32(24_000)
	annotations, _, err := infra.ChatCompletionHandler[LogicAnnotations](ctx, promptString, &thinkingBudget)
	if err != nil {
		return nil, fmt.Errorf("AIモデルの呼び出しに失敗しました: %w", err)
	}

	return annotations, nil
}
