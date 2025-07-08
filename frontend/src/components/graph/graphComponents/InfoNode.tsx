/*
import React from "react";
import { Handle, Position, useViewport } from "reactflow";

interface InfoNodeData {
  type: "importance" | "uniqueness";
  content: string;
}

const InfoNode = ({ data }: { data: InfoNodeData }) => {
  const { zoom } = useViewport();

  const opacity = Math.max(0, Math.min(1, (zoom - 0.8) * 5));
  const scale = Math.max(0.8, Math.min(1, zoom));

  return (
    <div
      className={`border-2 rounded-lg px-3 py-2 shadow-md min-w-[150px] max-w-[250px] relative transition-all duration-300 ease-out ${
        data.type === "importance"
          ? "bg-green-50 border-green-500"
          : "bg-purple-50 border-purple-500"
      }`}
      style={{
        opacity,
        transform: `scale(${scale})`,
        transformOrigin: "left center",
      }}
    >
      <Handle type="target" position={Position.Left} />
      <div
        className={`text-xs font-bold mb-1 ${
          data.type === "importance" ? "text-green-700" : "text-purple-700"
        }`}
      >
        {data.type === "importance" ? "IMPORTANCE" : "UNIQUENESS"}
      </div>
      <div className="text-xs text-gray-700 break-words">{data.content}</div>
    </div>
  );
};

export default InfoNode;
export type { InfoNodeData };
*/
