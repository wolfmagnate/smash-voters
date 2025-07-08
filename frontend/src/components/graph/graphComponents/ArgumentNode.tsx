import React from "react";
import { Handle, Position, useViewport } from "reactflow";
import { ZOOM_THRESHOLDS } from "../constants/graphConstants";

interface ArgumentNodeData {
  label: string;
  isRebuttal: boolean;
  importance?: string[];
  uniqueness?: string[];
}

const ArgumentNode = ({ data }: { data: ArgumentNodeData }) => {
  const { zoom } = useViewport();

  const shouldShowDetails = zoom > ZOOM_THRESHOLDS.SHOW_NODE_DETAILS;

  return (
    <div className="relative">
      <div
        className="bg-white border-2 border-blue-500 rounded-lg px-4 py-3 shadow-lg w-[500px] relative transition-all duration-500 ease-out"
        style={{
          minHeight: shouldShowDetails ? "auto" : "200px",
          display: "flex",
          flexDirection: "column",
          justifyContent: "center",
        }}
      >
        <Handle type="target" position={Position.Top} />
        <Handle type="source" position={Position.Bottom} />

        <div
          className="font-black text-center text-gray-800 break-words transition-all duration-500 ease-out"
          style={{
            fontSize: shouldShowDetails ? "1.5rem" : "3rem",
            lineHeight: shouldShowDetails ? "1.2" : "1.3",
            padding: shouldShowDetails ? "0" : "1rem 0",
          }}
        >
          {data.label}
        </div>

        {data.isRebuttal && (
          <div className="text-xs text-red-600 text-center mt-1 font-semibold">
            反論
          </div>
        )}
      </div>

      {data.importance && data.importance.length > 0 && (
        <div
          className="mt-2 border-2 border-green-300 rounded-lg p-2 bg-green-50 w-[500px] transition-all duration-500 ease-out"
          style={{
            opacity: shouldShowDetails ? 1 : 0,
            visibility: shouldShowDetails ? "visible" : "hidden",
            transform: shouldShowDetails
              ? "translateY(0) scale(1)"
              : "translateY(-10px) scale(0.95)",
            filter: shouldShowDetails ? "blur(0px)" : "blur(2px)",
          }}
        >
          <div className="text-xs font-semibold text-green-700 mb-1 text-center">
            重要性:
          </div>
          {data.importance.map((imp, index) => (
            <div
              key={index}
              className="text-xs text-green-600 px-2 py-1 rounded mb-1 break-words"
            >
              • {imp}
            </div>
          ))}
        </div>
      )}

      {data.uniqueness && data.uniqueness.length > 0 && (
        <div
          className="mt-2 border-2 border-purple-300 rounded-lg p-2 bg-purple-50 w-[500px] transition-all duration-500 ease-out"
          style={{
            opacity: shouldShowDetails ? 1 : 0,
            visibility: shouldShowDetails ? "visible" : "hidden",
            transform: shouldShowDetails
              ? "translateY(0) scale(1)"
              : "translateY(-10px) scale(0.95)",
            filter: shouldShowDetails ? "blur(0px)" : "blur(2px)",
            transitionDelay: shouldShowDetails ? "100ms" : "0ms",
          }}
        >
          <div className="text-xs font-semibold text-purple-700 mb-1 text-center">
            独自性:
          </div>
          {data.uniqueness.map((uniq, index) => (
            <div
              key={index}
              className="text-xs text-purple-600 px-2 py-1 rounded mb-1 break-words"
            >
              • {uniq}
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default ArgumentNode;
export type { ArgumentNodeData };
