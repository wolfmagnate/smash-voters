package domain

import (
	"fmt"
)

// DebateGraphNode は、ディベートの論理構造グラフにおける個々の「主張」や「事象」を表すノードです。
type DebateGraphNode struct {
	// Argument は、このノードが表現する具体的な主張の内容を保持します。
	Argument *Assertion

	// Causes は、このノード（結果）を引き起こす原因となった因果関係（エッジ）のスライス（リスト）です。
	// グラフ構造を結果側から原因側へ辿る際に使用されます。
	// このノードを Effect とするエッジがこのリストに含まれます。
	Causes []*DebateGraphEdge

	// Importance は、このノードの主張がなぜ重要なのかを補強するための論拠のリストです。
	Importance []*Assertion

	// Uniqueness は、この主張がなぜ特定のプランでのみ発生するのかを補強するための論拠のリストです。
	Uniqueness []*Assertion

	// ImportanceRebuttals は、「重要性」に対する反論のリストです。
	ImportanceRebuttals []*Assertion

	// UniquenessRebuttals は、「独自性」に対する反論のリストです。
	UniquenessRebuttals []*Assertion

	// IsRebuttal は、このノード自体が相手の議論に対する反論として提示されたものかどうかを示す真偽値です。
	// trueの場合、このノードは反対意見、新たなデメリットの提示、各種反論など、反駁のために構築された主張であることを意味します。
	IsRebuttal bool
}

// DebateGraphEdge は、2つのノード（主張）間の「因果関係」を表す有向エッジです。
type DebateGraphEdge struct {
	// Cause は、因果関係における「原因」となるノードへのポインタです。
	Cause *DebateGraphNode

	// Effect は、因果関係における「結果」となるノードへのポインタです。
	Effect *DebateGraphNode

	// Certainty は、この因果関係が確実であることを補強するための論拠のリストです。
	Certainty []*Assertion

	// Uniqueness は、この因果関係の独自性を補強するための論拠のリストです。
	Uniqueness []*Assertion

	// CertaintyRebuttal は、「確実性」に対する反論のリストです。
	CertaintyRebuttal []*Assertion

	// UniquenessRebuttals は、「独自性」に対する反論のリストです。
	UniquenessRebuttals []*Assertion

	// IsRebuttal は、このエッジ自体が相手の議論に対する反論の一部として提示されたものかどうかを示す真偽値です。
	// 特に、相手の主張（ノード）を途中まで認め、そこから新たなデメリットや意図せざる結果を導き出す「ターンアラウンド」反論などで使用されます。
	IsRebuttal bool
}

func NewDebateGraphNode(argument *Assertion, isRebuttal bool) *DebateGraphNode {
	return &DebateGraphNode{
		Argument:            argument,
		Causes:              make([]*DebateGraphEdge, 0),
		Importance:          make([]*Assertion, 0),
		Uniqueness:          make([]*Assertion, 0),
		ImportanceRebuttals: make([]*Assertion, 0),
		UniquenessRebuttals: make([]*Assertion, 0),
		IsRebuttal:          isRebuttal,
	}
}

func NewDebateGraphEdge(cause, effect *DebateGraphNode, isRebuttal bool) *DebateGraphEdge {
	return &DebateGraphEdge{
		Cause:               cause,
		Effect:              effect,
		Certainty:           make([]*Assertion, 0),
		Uniqueness:          make([]*Assertion, 0),
		CertaintyRebuttal:   make([]*Assertion, 0),
		UniquenessRebuttals: make([]*Assertion, 0),
		IsRebuttal:          isRebuttal,
	}
}

type DebateGraph struct {
	Nodes                    []*DebateGraphNode
	NodeRebuttals            []*DebateGraphNodeRebuttal
	EdgeRebuttals            []*DebateGraphEdgeRebuttal
	CounterArgumentRebuttals []*CounterArgumentRebuttal
	TurnArgumentRebuttals    []*TurnArgumentRebuttal

	nodeMap map[string]*DebateGraphNode // 小文字で非公開にし、メソッド経由でアクセス
	edgeMap map[string]*DebateGraphEdge // キー: "CauseArgument->EffectArgument"
}

func NewDebateGraph() *DebateGraph {
	return &DebateGraph{
		Nodes:         make([]*DebateGraphNode, 0),
		NodeRebuttals: make([]*DebateGraphNodeRebuttal, 0),
		EdgeRebuttals: make([]*DebateGraphEdgeRebuttal, 0),
		nodeMap:       make(map[string]*DebateGraphNode),
		edgeMap:       make(map[string]*DebateGraphEdge),
	}
}

// AddNode はグラフにノードを追加します。
// 同じArgumentを持つノードが既に存在する場合はエラーを返します。
func (dg *DebateGraph) AddNode(node *DebateGraphNode) error {
	if node == nil || node.Argument == nil {
		return fmt.Errorf("cannot add a nil node or a node with a nil argument to DebateGraph")
	}
	if _, exists := dg.nodeMap[node.Argument.Statement]; exists {
		// 既に存在する場合、エラーを返すか、既存ノードを返すか、何もしないかは設計次第。
		// ここではエラーとして、呼び出し元に重複を通知します。
		return fmt.Errorf("node with argument '%s' already exists in DebateGraph", node.Argument.Statement)
	}
	dg.Nodes = append(dg.Nodes, node)
	dg.nodeMap[node.Argument.Statement] = node
	return nil
}

// GetNode はArgument文字列によってノードを取得します。
func (dg *DebateGraph) GetNode(argument string) (*DebateGraphNode, bool) {
	node, exists := dg.nodeMap[argument]
	return node, exists
}

// generateEdgeKey はエッジマップ用のキーを生成する内部ヘルパー関数です。
func generateEdgeKey(causeArgument, effectArgument string) string {
	return fmt.Sprintf("%s->%s", causeArgument, effectArgument)
}

// AddEdge はグラフにエッジを追加します。
// エッジのCauseノードとEffectノードは事前にグラフに追加されている必要があります。
// EffectノードのCausesリストも更新します。
func (dg *DebateGraph) AddEdge(edge *DebateGraphEdge) error {
	if edge == nil {
		return fmt.Errorf("cannot add a nil edge to DebateGraph")
	}
	if edge.Cause == nil || edge.Effect == nil || edge.Cause.Argument == nil || edge.Effect.Argument == nil {
		return fmt.Errorf("edge must have valid cause and effect nodes with non-nil arguments")
	}

	// エッジが参照するノードがグラフに存在することを確認
	if _, exists := dg.nodeMap[edge.Cause.Argument.Statement]; !exists {
		return fmt.Errorf("cause node '%s' of the edge is not in the graph", edge.Cause.Argument.Statement)
	}
	effectNodeInMap, effectNodeExists := dg.nodeMap[edge.Effect.Argument.Statement]
	if !effectNodeExists {
		return fmt.Errorf("effect node '%s' of the edge is not in the graph", edge.Effect.Argument.Statement)
	}
	// edge.Effectがマップ内のインスタンスと同じであることを保証（通常、呼び出し側が正しく構築すれば問題ない）
	if edge.Effect != effectNodeInMap {
		return fmt.Errorf("edge's effect node instance does not match the instance in the graph's nodeMap for argument '%s'", edge.Effect.Argument.Statement)
	}

	edgeKey := generateEdgeKey(edge.Cause.Argument.Statement, edge.Effect.Argument.Statement)
	if _, exists := dg.edgeMap[edgeKey]; exists {
		return nil
	}

	dg.edgeMap[edgeKey] = edge
	// EffectノードのCausesリストにこのエッジを追加
	// edge.Effect はグラフ内の正しいインスタンスである前提
	edge.Effect.Causes = append(edge.Effect.Causes, edge)
	return nil
}

func (dg *DebateGraph) RemoveEdge(causeArgument, effectArgument string) error {
	edgeKey := generateEdgeKey(causeArgument, effectArgument)
	edge, exists := dg.edgeMap[edgeKey]
	if !exists {
		return fmt.Errorf("削除対象のエッジ '%s' がグラフ内に見つかりません", edgeKey)
	}

	// 1. edgeMapからエッジを削除します。
	delete(dg.edgeMap, edgeKey)

	// 2. EffectノードのCausesスライスから該当するエッジを削除します。
	//    スライスをループ処理し、削除対象以外の要素で新しいスライスを構築し直します。
	effectNode := edge.Effect
	newCauses := make([]*DebateGraphEdge, 0, len(effectNode.Causes)-1)
	for _, e := range effectNode.Causes {
		// ポインタを比較して、削除対象のエッジと同一でないものだけを新しいスライスに追加します。
		if e != edge {
			newCauses = append(newCauses, e)
		}
	}
	// EffectノードのCausesスライスを、更新された新しいスライスで上書きします。
	effectNode.Causes = newCauses

	return nil
}

func (dg *DebateGraph) GetEdge(causeArgument string, effectArgument string) (*DebateGraphEdge, bool) {
	edgeKey := generateEdgeKey(causeArgument, effectArgument)
	edge, exists := dg.edgeMap[edgeKey]
	return edge, exists
}

func (dg *DebateGraph) GetAllEdges() []*DebateGraphEdge {
	edges := make([]*DebateGraphEdge, 0, len(dg.edgeMap))
	for _, edge := range dg.edgeMap {
		edges = append(edges, edge)
	}
	return edges
}
