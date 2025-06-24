package domain

import (
	"fmt"
)

type LogicGraphNode struct {
	Argument string
	Causes   []*LogicGraphNode
}

type LogicGraph struct {
	Nodes   []*LogicGraphNode
	NodeMap map[string]*LogicGraphNode
}

func NewLogicGraphNode(argument string) *LogicGraphNode {
	return &LogicGraphNode{
		Argument: argument,
		Causes:   make([]*LogicGraphNode, 0),
	}
}

func NewLogicGraph(initialNodes []*LogicGraphNode) *LogicGraph {
	graph := &LogicGraph{
		Nodes:   make([]*LogicGraphNode, 0),
		NodeMap: make(map[string]*LogicGraphNode),
	}
	for _, node := range initialNodes {
		graph.AddNode(node)
	}
	return graph
}

func (lg *LogicGraph) AddNode(node *LogicGraphNode) {
	if node == nil {
		fmt.Println("Cannot add a nil node.")
		return
	}
	if _, exists := lg.NodeMap[node.Argument]; exists {
		fmt.Printf("Node with argument '%s' already exists. Skipping.\n", node.Argument)
		return
	}
	lg.Nodes = append(lg.Nodes, node)
	lg.NodeMap[node.Argument] = node
}

func ListAllCausalRelationships(graph *LogicGraph) []string {
	var relationships []string

	if graph == nil {
		return relationships // 空のグラフの場合は空のリストを返す
	}

	for _, effect := range graph.Nodes {
		if effect == nil {
			continue // ノードがnilの場合はスキップ
		}
		for _, cause := range effect.Causes {
			if cause == nil {
				continue // 結果ノードがnilの場合はスキップ
			}
			relationship := fmt.Sprintf("- 「%s」であることが「%s」を引き起こす", cause.Argument, effect.Argument)
			relationships = append(relationships, relationship)
		}
	}

	return relationships
}
