// src/pages/result/index.tsx
"use client";
import React, { useMemo } from "react";
import { useRouter, useSearchParams } from "next/navigation";
import { useLatestElection } from "@/hooks/useLatestElection";
import { useMatching } from "@/hooks/useMatching";
import { UserAnswer } from "@/services/matchingService";
import { Spinner } from "@/components/ui/Spinner";

export default function Result() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const { election } = useLatestElection();

  const matchRequest = useMemo(() => {
    const answersParam = searchParams.get("answers");
    const importantParam = searchParams.get("important");

    if (!answersParam || !importantParam) {
      return null;
    }

    try {
      const answers: UserAnswer[] = JSON.parse(answersParam);
      const important_question_ids: number[] = importantParam
        .split(",")
        .map(Number);

      return {
        answers,
        important_question_ids,
      };
    } catch (error) {
      console.error("Failed to parse URL parameters:", error);
      return null;
    }
  }, [searchParams]);

  const { matchResult, loading, error } = useMatching(
    election?.id || null,
    matchRequest
  );

  if (!matchRequest) {
    return (
      <div className="max-w-4xl mx-auto p-4 text-center">
        <h1 className="text-3xl font-bold text-black mb-6">エラー</h1>
        <p className="text-xl text-red-500 mb-4">回答データが見つかりません</p>
        <button
          onClick={() => router.push("/quiz")}
          className="px-6 py-3 bg-blue-500 text-white rounded-lg hover:bg-blue-600"
        >
          クイズをやり直す
        </button>
      </div>
    );
  }

  if (loading) {
    return (
      <div className="max-w-4xl mx-auto p-4 text-center">
        <Spinner />
        <p className="text-xl mt-4">マッチング結果を計算中...</p>
      </div>
    );
  }

  if (error) {
    return (
      <div className="max-w-4xl mx-auto p-4 text-center">
        <h1 className="text-3xl font-bold text-black mb-6">エラー</h1>
        <p className="text-xl text-red-500 mb-4">{error}</p>
        <button
          onClick={() => router.push("/quiz")}
          className="px-6 py-3 bg-blue-500 text-white rounded-lg hover:bg-blue-600"
        >
          クイズをやり直す
        </button>
      </div>
    );
  }

  if (!matchResult) {
    return (
      <div className="max-w-4xl mx-auto p-4 text-center">
        <p className="text-xl">結果が見つかりません</p>
      </div>
    );
  }

  return (
    <div className="max-w-4xl mx-auto p-4">
      <h1 className="text-4xl font-bold text-black mb-8 text-center">
        マッチング結果
      </h1>

      {/* トップマッチ */}
      <div className="bg-gradient-to-r from-blue-500 to-purple-600 text-white p-8 rounded-xl mb-8 text-center">
        <h2 className="text-2xl font-bold mb-4">あなたに最も適した政党</h2>
        <h3 className="text-3xl font-bold mb-2">
          {matchResult.top_match.party_name}
        </h3>
        <p className="text-xl">マッチ度: {matchResult.top_match.match_rate}%</p>
      </div>

      {/* 全結果 */}
      <div className="space-y-4 mb-8">
        <h2 className="text-2xl font-bold text-black mb-4">全政党との適合度</h2>
        {matchResult.results
          .sort((a, b) => b.match_rate - a.match_rate)
          .map((result, index) => (
            <div
              key={result.party_name}
              className={`p-4 rounded-lg border ${
                index === 0 ? "border-blue-500 bg-blue-50" : "border-gray-300"
              }`}
            >
              <div className="flex justify-between items-center">
                <span className="text-lg font-semibold">
                  {result.party_name}
                </span>
                <span className="text-xl font-bold">{result.match_rate}%</span>
              </div>
              <div className="w-full bg-gray-200 rounded-full h-2 mt-2">
                <div
                  className={`h-2 rounded-full ${
                    index === 0 ? "bg-blue-500" : "bg-gray-400"
                  }`}
                  style={{ width: `${result.match_rate}%` }}
                />
              </div>
            </div>
          ))}
      </div>

      {/* アクション */}
      <div className="text-center">
        <button
          onClick={() => router.push("/quiz")}
          className="px-6 py-3 bg-green-500 text-white rounded-lg hover:bg-green-600 mr-4"
        >
          もう一度クイズを受ける
        </button>
        <button
          onClick={() => router.push("/")}
          className="px-6 py-3 bg-gray-500 text-white rounded-lg hover:bg-gray-600"
        >
          ホームに戻る
        </button>
      </div>
    </div>
  );
}
