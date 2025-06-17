package domain

import (
	"fmt"
)

type DebateGraphNode struct {
	Argument            string
	Causes              []*DebateGraphEdge
	Importance          []string
	Uniqueness          []string
	ImportanceRebuttals []string
	UniquenessRebuttals []string

	IsRebuttal bool
}

type DebateGraphEdge struct {
	Cause               *DebateGraphNode
	Effect              *DebateGraphNode
	Certainty           []string
	Uniqueness          []string
	CertaintyRebuttal   []string
	UniquenessRebuttals []string

	IsRebuttal bool
}

func NewDebateGraphNode(argument string, isRebuttal bool) *DebateGraphNode {
	return &DebateGraphNode{
		Argument:            argument,
		Causes:              make([]*DebateGraphEdge, 0),
		Importance:          make([]string, 0),
		Uniqueness:          make([]string, 0),
		ImportanceRebuttals: make([]string, 0),
		UniquenessRebuttals: make([]string, 0),
		IsRebuttal:          isRebuttal,
	}
}

func NewDebateGraphEdge(cause, effect *DebateGraphNode, isRebuttal bool) *DebateGraphEdge {
	return &DebateGraphEdge{
		Cause:               cause,
		Effect:              effect,
		Certainty:           make([]string, 0),
		Uniqueness:          make([]string, 0),
		CertaintyRebuttal:   make([]string, 0),
		UniquenessRebuttals: make([]string, 0),
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
	if node == nil {
		return fmt.Errorf("cannot add a nil node to DebateGraph")
	}
	if _, exists := dg.nodeMap[node.Argument]; exists {
		// 既に存在する場合、エラーを返すか、既存ノードを返すか、何もしないかは設計次第。
		// ここではエラーとして、呼び出し元に重複を通知します。
		return fmt.Errorf("node with argument '%s' already exists in DebateGraph", node.Argument)
	}
	dg.Nodes = append(dg.Nodes, node)
	dg.nodeMap[node.Argument] = node
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
	if edge.Cause == nil || edge.Effect == nil {
		return fmt.Errorf("edge must have valid cause and effect nodes")
	}

	// エッジが参照するノードがグラフに存在することを確認
	if _, exists := dg.nodeMap[edge.Cause.Argument]; !exists {
		return fmt.Errorf("cause node '%s' of the edge is not in the graph", edge.Cause.Argument)
	}
	effectNodeInMap, effectNodeExists := dg.nodeMap[edge.Effect.Argument]
	if !effectNodeExists {
		return fmt.Errorf("effect node '%s' of the edge is not in the graph", edge.Effect.Argument)
	}
	// edge.Effectがマップ内のインスタンスと同じであることを保証（通常、呼び出し側が正しく構築すれば問題ない）
	if edge.Effect != effectNodeInMap {
		return fmt.Errorf("edge's effect node instance does not match the instance in the graph's nodeMap for argument '%s'", edge.Effect.Argument)
	}

	edgeKey := generateEdgeKey(edge.Cause.Argument, edge.Effect.Argument)
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
