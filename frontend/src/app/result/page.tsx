// src/pages/result/index.tsx
"use client";

export default function Result() {
  return (
    <div className="max-w-4xl mx-auto p-4 text-center">
      <h1 className="font-bold text-black mb-6 text-center">結果ページ</h1>
      <p className="text-2xl text-black">政策マッチングの結果を表示します。</p>
      {/* TODO: ここにマッチ度が最も高い政党を表示 */}
    </div>
  );
}
