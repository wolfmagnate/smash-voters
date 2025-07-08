import { useState, useEffect, useCallback, useMemo } from "react";
import { DebateGraph } from "../../../types/debate-graph";
import { debateGraphService } from "../services/GraphService";
import {
  Node,
  Connection,
  addEdge,
  useNodesState,
  useEdgesState,
  useReactFlow,
} from "reactflow";
import { getLayoutedElements } from "../utils/layoutUtils";
import { handleNodeClick } from "../utils/nodeUtils";
import { createFlowElements } from "../utils/flowUtils";
import { ZOOM_LEVEL } from "../constants/graphConstants";

export function useGraph() {
  const [data, setData] = useState<DebateGraph | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const { setViewport } = useReactFlow();

  const { nodes: flowNodes, edges: flowEdges } = useMemo(() => {
    if (!data) return { nodes: [], edges: [] };
    return createFlowElements(data, getLayoutedElements);
  }, [data]);

  const [nodes, setNodes] = useNodesState(flowNodes);
  const [edges, setEdges] = useEdgesState(flowEdges);

  useEffect(() => {
    setNodes(flowNodes);
    setEdges(flowEdges);
  }, [flowNodes, flowEdges, setNodes, setEdges]);

  const onConnect = useCallback(
    (params: Connection) => setEdges((eds) => addEdge(params, eds)),
    [setEdges]
  );

  const onNodeClick = useCallback(
    (_event: React.MouseEvent, node: Node) => {
      handleNodeClick(node, setViewport, ZOOM_LEVEL);
    },
    [setViewport]
  );

  const fetchData = async () => {
    try {
      setLoading(true);
      setError(null);
      const result = await debateGraphService.getDebateGraph();
      setData(result);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Unknown error occurred");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchData();
  }, []);

  return {
    data,
    loading,
    error,
    refetch: fetchData,
    nodes,
    edges,
    onConnect,
    onNodeClick,
  };
}
