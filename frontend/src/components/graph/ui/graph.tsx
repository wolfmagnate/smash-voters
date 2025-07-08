"use client";

import React from "react";
import ReactFlow, {
  NodeTypes,
  EdgeTypes,
  Controls,
  Background,
  BackgroundVariant,
} from "reactflow";
import { ReactFlowProvider } from "reactflow";
import "reactflow/dist/style.css";
import ArgumentNode from "../graphComponents/ArgumentNode";
import CustomEdge from "../graphComponents/CustomEdge";
import { useGraph } from "../hooks/useGraph";

const nodeTypes: NodeTypes = {
  argument: ArgumentNode,
};

const edgeTypes: EdgeTypes = {
  custom: CustomEdge,
};

function GraphComponent() {
  const { data, loading, error, nodes, edges, onConnect, onNodeClick } =
    useGraph();

  if (loading) {
    return (
      <div className="flex items-center justify-center h-[600px] bg-white">
        <div className="relative w-12 h-12">
          <div className="absolute w-12 h-12 border-4 border-gray-200 rounded-full"></div>
          <div className="absolute w-12 h-12 border-4 border-blue-500 rounded-full border-t-transparent animate-spin"></div>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex items-center justify-center h-[600px] bg-white">
        <div className="text-lg text-red-600">エラー: {error}</div>
      </div>
    );
  }

  if (!data) {
    return (
      <div className="flex items-center justify-center h-[600px] bg-white">
        <div className="text-lg">データがありません</div>
      </div>
    );
  }

  return (
    <div style={{ width: "100%", height: "600px" }}>
      <ReactFlow
        nodes={nodes}
        edges={edges}
        nodesDraggable={false}
        nodesConnectable={false}
        elementsSelectable={false}
        onConnect={onConnect}
        onNodeClick={onNodeClick}
        nodeTypes={nodeTypes}
        edgeTypes={edgeTypes}
        fitView
        fitViewOptions={{
          padding: 0.1,
          includeHiddenNodes: false,
          minZoom: 0.1,
          maxZoom: 2,
        }}
        defaultViewport={{ x: 0, y: 0, zoom: 1.0 }}
        minZoom={0.1}
        maxZoom={2}
        attributionPosition="bottom-left"
      >
        <Controls showInteractive={false} />
        <Background variant={BackgroundVariant.Dots} gap={12} size={1} />
      </ReactFlow>
    </div>
  );
}

export default function Graph() {
  return (
    <div className="min-h-screen p-8 bg-gray-100">
      <div className="max-w-7xl mx-auto">
        <div className="bg-white rounded-lg shadow-lg p-2">
          <ReactFlowProvider>
            <GraphComponent />
          </ReactFlowProvider>
        </div>
      </div>
    </div>
  );
}
