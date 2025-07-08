import { Node } from "reactflow";

export const handleNodeClick = (
  node: Node,
  setViewport: (
    viewport: { x: number; y: number; zoom: number },
    options?: { duration?: number }
  ) => void,
  zoomLevel: number
) => {
  const nodeTop = {
    x: node.position.x + (node.width || 500) / 2,
    y: node.position.y,
  };

  const reactFlowElement = document.querySelector(".react-flow__viewport")
    ?.parentElement as HTMLElement;
  const containerWidth = reactFlowElement?.offsetWidth || window.innerWidth;
  const containerHeight = reactFlowElement?.offsetHeight || 600;

  const newViewport = {
    x: containerWidth / 2 - nodeTop.x * zoomLevel,
    y: containerHeight / 4 - nodeTop.y * zoomLevel,
    zoom: zoomLevel,
  };

  setViewport(newViewport, { duration: 600 });
};
