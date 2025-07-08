import dagre from "dagre";
import { Node, Edge } from "reactflow";

export const getLayoutedElements = (nodes: Node[], edges: Edge[]) => {
  return calculateLayout(nodes, edges);
};

export const calculateLayout = (nodes: Node[], edges: Edge[]) => {
  const dagreGraph = new dagre.graphlib.Graph();
  dagreGraph.setDefaultEdgeLabel(() => ({}));
  dagreGraph.setGraph({
    rankdir: "TB",
    nodesep: 150,
    edgesep: 80,
    ranksep: 250,
  });

  nodes.forEach((node) => {
    if (node.type === "argument") {
      const nodeData = node.data as {
        importance?: string[];
        uniqueness?: string[];
      };
      const importanceCount = nodeData?.importance?.length || 0;
      const uniquenessCount = nodeData?.uniqueness?.length || 0;

      const extraHeight = (importanceCount + uniquenessCount) * 40;

      const baseWidth = 500;
      const baseHeight = 120;

      dagreGraph.setNode(node.id, {
        width: baseWidth,
        height: baseHeight + extraHeight,
      });
    }
  });

  edges.forEach((edge) => {
    dagreGraph.setEdge(edge.source, edge.target);
  });

  // レイアウトを計算
  dagre.layout(dagreGraph);

  // 新しい位置を取得
  const layoutedNodes = nodes.map((node) => {
    const nodeWithPosition = dagreGraph.node(node.id);
    return {
      ...node,
      position: {
        x: nodeWithPosition.x - nodeWithPosition.width / 2,
        y: nodeWithPosition.y - nodeWithPosition.height / 2,
      },
    };
  });

  return { nodes: layoutedNodes, edges };
};
