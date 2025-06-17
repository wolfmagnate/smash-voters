package domain

import (
	"encoding/json"
	"fmt"
)

type jsonNode struct {
	Argument            string   `json:"argument"`
	IsRebuttal          bool     `json:"is_rebuttal"`
	Importance          []string `json:"importance,omitempty"`
	Uniqueness          []string `json:"uniqueness,omitempty"`
	ImportanceRebuttals []string `json:"importance_rebuttals,omitempty"`
	UniquenessRebuttals []string `json:"uniqueness_rebuttals,omitempty"`
}

func (n *DebateGraphNode) ToJSON() (string, error) {
	if n == nil {
		return "", fmt.Errorf("cannot convert nil DebateGraphNode to JSON")
	}

	jNode := &jsonNode{
		Argument:            n.Argument,
		IsRebuttal:          n.IsRebuttal,
		Importance:          n.Importance,
		Uniqueness:          n.Uniqueness,
		ImportanceRebuttals: n.ImportanceRebuttals,
		UniquenessRebuttals: n.UniquenessRebuttals,
	}

	jsonData, err := json.MarshalIndent(jNode, "", "    ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal DebateGraphNode to JSON: %w", err)
	}

	return string(jsonData), nil
}

type jsonEdge struct {
	Cause               string   `json:"cause"`
	Effect              string   `json:"effect"`
	IsRebuttal          bool     `json:"is_rebuttal"`
	Certainty           []string `json:"certainty,omitempty"`
	Uniqueness          []string `json:"uniqueness,omitempty"`
	CertaintyRebuttal   []string `json:"certainty_rebuttal,omitempty"`
	UniquenessRebuttals []string `json:"uniqueness_rebuttals,omitempty"`
}

type jsonNodeRebuttal struct {
	TargetArgument   string `json:"target_argument"`
	RebuttalType     string `json:"rebuttal_type"`
	RebuttalArgument string `json:"rebuttal_argument"`
}

type jsonEdgeRebuttal struct {
	TargetCauseArgument  string `json:"target_cause_argument"`
	TargetEffectArgument string `json:"target_effect_argument"`
	RebuttalType         string `json:"rebuttal_type"`
	RebuttalArgument     string `json:"rebuttal_argument"`
}

func (e *DebateGraphEdge) ToJSON() (string, error) {
	if e == nil {
		return "", fmt.Errorf("cannot convert nil DebateGraphEdge to JSON")
	}
	if e.Cause == nil || e.Effect == nil {
		return "", fmt.Errorf("cannot marshal edge with nil cause or effect")
	}

	jEdge := &jsonEdge{
		Cause:               e.Cause.Argument,
		Effect:              e.Effect.Argument,
		IsRebuttal:          e.IsRebuttal,
		Certainty:           e.Certainty,
		Uniqueness:          e.Uniqueness,
		CertaintyRebuttal:   e.CertaintyRebuttal,
		UniquenessRebuttals: e.UniquenessRebuttals,
	}

	jsonData, err := json.MarshalIndent(jEdge, "", "    ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal DebateGraphEdge to JSON: %w", err)
	}

	return string(jsonData), nil
}

type jsonCounterArgumentRebuttal struct {
	RebuttalArgument string `json:"rebuttal_argument"`
	TargetArgument   string `json:"target_argument"`
}

type jsonTurnArgumentRebuttal struct {
	RebuttalArgument string `json:"rebuttal_argument"`
}

type jsonGraph struct {
	Nodes                    []*jsonNode                    `json:"nodes"`
	Edges                    []*jsonEdge                    `json:"edges"`
	NodeRebuttals            []*jsonNodeRebuttal            `json:"node_rebuttals,omitempty"`
	EdgeRebuttals            []*jsonEdgeRebuttal            `json:"edge_rebuttals,omitempty"`
	CounterArgumentRebuttals []*jsonCounterArgumentRebuttal `json:"counter_argument_rebuttals,omitempty"`
	TurnArgumentRebuttals    []*jsonTurnArgumentRebuttal    `json:"turn_argument_rebuttals,omitempty"`
}

func (dg *DebateGraph) ToJSON() (string, error) {
	if dg == nil {
		return "", fmt.Errorf("cannot convert nil DebateGraph to JSON")
	}

	jGraph := &jsonGraph{
		Nodes:                    make([]*jsonNode, 0, len(dg.Nodes)),
		Edges:                    make([]*jsonEdge, 0, len(dg.edgeMap)),
		NodeRebuttals:            make([]*jsonNodeRebuttal, 0, len(dg.NodeRebuttals)),
		EdgeRebuttals:            make([]*jsonEdgeRebuttal, 0, len(dg.EdgeRebuttals)),
		CounterArgumentRebuttals: make([]*jsonCounterArgumentRebuttal, 0, len(dg.CounterArgumentRebuttals)),
		TurnArgumentRebuttals:    make([]*jsonTurnArgumentRebuttal, 0, len(dg.TurnArgumentRebuttals)),
	}

	// ノードの変換
	for _, node := range dg.Nodes {
		jGraph.Nodes = append(jGraph.Nodes, &jsonNode{
			Argument:            node.Argument,
			IsRebuttal:          node.IsRebuttal,
			Importance:          node.Importance,
			Uniqueness:          node.Uniqueness,
			ImportanceRebuttals: node.ImportanceRebuttals,
			UniquenessRebuttals: node.UniquenessRebuttals,
		})
	}

	// エッジの変換
	for _, edge := range dg.edgeMap {
		jGraph.Edges = append(jGraph.Edges, &jsonEdge{
			Cause:               edge.Cause.Argument,
			Effect:              edge.Effect.Argument,
			IsRebuttal:          edge.IsRebuttal,
			Certainty:           edge.Certainty,
			Uniqueness:          edge.Uniqueness,
			CertaintyRebuttal:   edge.CertaintyRebuttal,
			UniquenessRebuttals: edge.UniquenessRebuttals,
		})
	}

	// ノード反論の変換
	for _, r := range dg.NodeRebuttals {
		jGraph.NodeRebuttals = append(jGraph.NodeRebuttals, &jsonNodeRebuttal{
			TargetArgument:   r.TargetNode.Argument,
			RebuttalType:     r.RebuttalType,
			RebuttalArgument: r.RebuttalNode.Argument,
		})
	}

	// エッジ反論の変換
	for _, r := range dg.EdgeRebuttals {
		jGraph.EdgeRebuttals = append(jGraph.EdgeRebuttals, &jsonEdgeRebuttal{
			TargetCauseArgument:  r.TargetEdge.Cause.Argument,
			TargetEffectArgument: r.TargetEdge.Effect.Argument,
			RebuttalType:         r.RebuttalType,
			RebuttalArgument:     r.RebuttalNode.Argument,
		})
	}

	// 反対意見の変換
	for _, r := range dg.CounterArgumentRebuttals {
		jGraph.CounterArgumentRebuttals = append(jGraph.CounterArgumentRebuttals, &jsonCounterArgumentRebuttal{
			RebuttalArgument: r.RebuttalNode.Argument,
			TargetArgument:   r.TargetNode.Argument,
		})
	}

	// ターンアラウンドの変換
	for _, r := range dg.TurnArgumentRebuttals {
		jGraph.TurnArgumentRebuttals = append(jGraph.TurnArgumentRebuttals, &jsonTurnArgumentRebuttal{
			RebuttalArgument: r.RebuttalNode.Argument,
		})
	}

	jsonData, err := json.MarshalIndent(jGraph, "", "    ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal DebateGraph to JSON: %w", err)
	}

	return string(jsonData), nil
}

func NewDebateGraphFromJSON(jsonData string) (*DebateGraph, error) {
	var jGraph jsonGraph
	if err := json.Unmarshal([]byte(jsonData), &jGraph); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON to DebateGraph: %w", err)
	}

	dg := NewDebateGraph()

	// 1. ノードをすべて構築
	for _, jNode := range jGraph.Nodes {
		node := NewDebateGraphNode(jNode.Argument, jNode.IsRebuttal)
		node.Importance = jNode.Importance
		node.Uniqueness = jNode.Uniqueness
		node.ImportanceRebuttals = jNode.ImportanceRebuttals
		node.UniquenessRebuttals = jNode.UniquenessRebuttals
		if err := dg.AddNode(node); err != nil {
			return nil, fmt.Errorf("failed to add node '%s' from JSON: %w", jNode.Argument, err)
		}
	}

	// 2. エッジをすべて構築
	for _, jEdge := range jGraph.Edges {
		causeNode, causeExists := dg.GetNode(jEdge.Cause)
		if !causeExists {
			return nil, fmt.Errorf("cause node '%s' for edge not found in graph", jEdge.Cause)
		}
		effectNode, effectExists := dg.GetNode(jEdge.Effect)
		if !effectExists {
			return nil, fmt.Errorf("effect node '%s' for edge not found in graph", jEdge.Effect)
		}

		edge := NewDebateGraphEdge(causeNode, effectNode, jEdge.IsRebuttal)
		edge.Certainty = jEdge.Certainty
		edge.Uniqueness = jEdge.Uniqueness
		edge.CertaintyRebuttal = jEdge.CertaintyRebuttal
		edge.UniquenessRebuttals = jEdge.UniquenessRebuttals

		if err := dg.AddEdge(edge); err != nil {
			return nil, fmt.Errorf("failed to add edge '%s -> %s' from JSON: %w", jEdge.Cause, jEdge.Effect, err)
		}
	}

	// 3. ノード反論を再構築
	for _, jRebuttal := range jGraph.NodeRebuttals {
		targetNode, exists := dg.GetNode(jRebuttal.TargetArgument)
		if !exists {
			return nil, fmt.Errorf("target node '%s' for node rebuttal not found", jRebuttal.TargetArgument)
		}
		rebuttalNode, exists := dg.GetNode(jRebuttal.RebuttalArgument)
		if !exists {
			return nil, fmt.Errorf("rebuttal node '%s' for node rebuttal not found", jRebuttal.RebuttalArgument)
		}

		rebuttal := &DebateGraphNodeRebuttal{
			TargetNode:   targetNode,
			RebuttalType: jRebuttal.RebuttalType,
			RebuttalNode: rebuttalNode,
		}
		dg.NodeRebuttals = append(dg.NodeRebuttals, rebuttal)
	}

	// 4. エッジ反論を再構築
	for _, jRebuttal := range jGraph.EdgeRebuttals {
		targetEdge, exists := dg.GetEdge(jRebuttal.TargetCauseArgument, jRebuttal.TargetEffectArgument)
		if !exists {
			return nil, fmt.Errorf("target edge '%s -> %s' for edge rebuttal not found", jRebuttal.TargetCauseArgument, jRebuttal.TargetEffectArgument)
		}
		rebuttalNode, exists := dg.GetNode(jRebuttal.RebuttalArgument)
		if !exists {
			return nil, fmt.Errorf("rebuttal node '%s' for edge rebuttal not found", jRebuttal.RebuttalArgument)
		}

		rebuttal := &DebateGraphEdgeRebuttal{
			TargetEdge:   targetEdge,
			RebuttalType: jRebuttal.RebuttalType,
			RebuttalNode: rebuttalNode,
		}
		dg.EdgeRebuttals = append(dg.EdgeRebuttals, rebuttal)
	}

	// 5. 反対意見を再構築
	for _, jRebuttal := range jGraph.CounterArgumentRebuttals {
		rebuttalNode, exists := dg.GetNode(jRebuttal.RebuttalArgument)
		if !exists {
			return nil, fmt.Errorf("rebuttal node '%s' for counter argument rebuttal not found", jRebuttal.RebuttalArgument)
		}

		targetNode, exiexists := dg.GetNode(jRebuttal.TargetArgument)
		if !exiexists {
			return nil, fmt.Errorf("target node '%s' for counter argument rebuttal not found", jRebuttal.TargetArgument)
		}

		rebuttal := &CounterArgumentRebuttal{
			RebuttalNode: rebuttalNode,
			TargetNode:   targetNode,
		}
		dg.CounterArgumentRebuttals = append(dg.CounterArgumentRebuttals, rebuttal)
	}

	// 6. ターンアラウンドを再構築
	for _, jRebuttal := range jGraph.TurnArgumentRebuttals {
		rebuttalNode, exists := dg.GetNode(jRebuttal.RebuttalArgument)
		if !exists {
			return nil, fmt.Errorf("rebuttal node '%s' for turn argument rebuttal not found", jRebuttal.RebuttalArgument)
		}

		rebuttal := &TurnArgumentRebuttal{
			RebuttalNode: rebuttalNode,
		}
		dg.TurnArgumentRebuttals = append(dg.TurnArgumentRebuttals, rebuttal)
	}

	return dg, nil
}
