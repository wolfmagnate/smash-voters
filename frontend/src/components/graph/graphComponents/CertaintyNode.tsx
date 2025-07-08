// import React from "react";
// import { Handle, Position, useViewport } from "reactflow";

// // カスタム確実性ノードコンポーネント
// interface CertaintyNodeData {
//   label: string;
//   isRebuttal: boolean;
//   certainty: string[];
// }

// const CertaintyNode = ({ data }: { data: CertaintyNodeData }) => {
//   const { zoom } = useViewport();

//   // ズームレベルに応じて詳細情報を表示するかどうかを判定
//   const shouldShowDetails = zoom > 0.6;

//   return (
//     <div className="relative">
//       {/* メインの確実性ノード（ズーム時のみ表示） */}

//       <div className="bg-white border-2 border-orange-500 rounded-lg px-2 py-1 shadow-lg min-w-[200px] max-w-[300px] relative">
//         {/* 上下のハンドル（中間ノード用） */}
//         <Handle type="target" position={Position.Top} />
//         <Handle type="source" position={Position.Bottom} />

//         {/* メインの確実性テキスト */}
//         <div className="text-xs font-medium text-center text-gray-800 break-words">
//           {data.label}
//         </div>

//         {data.isRebuttal && (
//           <div className="text-xs text-red-600 text-center mt-1 font-semibold">
//             反論
//           </div>
//         )}
//       </div>

//       {/* ズーム時でない場合は小さな点として表示 */}
//       {!shouldShowDetails && (
//         <div className="bg-orange-500 rounded-full w-4 h-4 shadow-lg relative">
//           <Handle type="target" position={Position.Top} />
//           <Handle type="source" position={Position.Bottom} />
//         </div>
//       )}

//       {/* Certainty セクション（ノードの真下、枠の外、ズーム時のみ表示） */}
//     </div>
//   );
// };

// export default CertaintyNode;
// export type { CertaintyNodeData };
