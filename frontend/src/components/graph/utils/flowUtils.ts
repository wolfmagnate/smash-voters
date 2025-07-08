import { Node, Edge } from "reactflow";
import { DebateGraph } from "../../../types/debate-graph";

export const createFlowElements = (
  data: DebateGraph,
  getLayoutedElements: (
    nodes: Node[],
    edges: Edge[]
  ) => { nodes: Node[]; edges: Edge[] }
) => {
  const nodes: Node[] = [];
  const edges: Edge[] = [];

  const argumentNodeMap = new Map<string, string>();

  data.nodes.forEach((node) => {
    const nodeId_flow = node.id;
    argumentNodeMap.set(node.id, nodeId_flow);

    nodes.push({
      id: nodeId_flow,
      type: "argument",
      position: { x: 0, y: 0 }, // dagreが後で調整する
      data: {
        label: node.argument.statement,
        isRebuttal: node.is_rebuttal,
        importance: node.importance,
        uniqueness: node.uniqueness,
      },
    });
  });

  data.edges.forEach((edge) => {
    const sourceId = argumentNodeMap.get(edge.cause_id);
    const targetId = argumentNodeMap.get(edge.effect_id);

    if (sourceId && targetId) {
      let edgeLabel = "";
      if (edge.is_rebuttal) {
        edgeLabel = "反論";
      }

      edges.push({
        id: `${sourceId}-${targetId}`,
        source: sourceId,
        target: targetId,
        type: "custom",
        style: {
          stroke: edge.is_rebuttal ? "#ef4444" : "#3b82f6",
          strokeWidth: 3,
        },
        animated: true,
        data: {
          label: edgeLabel,
          isRebuttal: edge.is_rebuttal,
        },
      });
    }
  });

  // dagreを使用してレイアウトを自動調整
  const { nodes: layoutedNodes } = getLayoutedElements(nodes, edges);

  return { nodes: layoutedNodes, edges };
};
