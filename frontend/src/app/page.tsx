"use client";
import React, { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { useLatestElection } from "@/hooks/useLatestElection";
import { Spinner } from "@/components/ui/Spinner";

export default function HomePage() {
  const router = useRouter();
  const [daysLeft, setDaysLeft] = useState(0);
  const { election, loading, error } = useLatestElection();

  useEffect(() => {
    const today = new Date();
    const electionDate = new Date(2025, 6, 29); // ← 6 ＝ 7
    const diffMs = electionDate.getTime() - today.getTime();
    const days = Math.ceil(diffMs / (1000 * 60 * 60 * 24));
    setDaysLeft(days > 0 ? days : 0);
  }, []);

  if (loading) {
    return <Spinner />;
  }

  return (
    <div className="w-full mx-auto px-4 py-8 text-center">
      <h1 className="text-5xl font-bold text-black mb-5">
        {error
          ? "Policy Evaluation System（仮）"
          : election?.name || "Policy Evaluation System（仮）"}
      </h1>

      <div className="bg-gradient-to-r from-blue-500 to-purple-600 py-15 mx-0 my-20 text-2xl font-extrabold text-white p-15 rounded-xl inline-block shadow-lg text-center">
        {daysLeft === 0 ? (
          <>今日は投票日です！</>
        ) : (
          <>
            次回の投票日は7月29日
            <br />
            あと{daysLeft}日
          </>
        )}
      </div>

      <p className="block w-fit mx-auto mb-10 px-3 py-1.5 text-2xl font-extrabold text-black text-center">
        ロジックグラフによって、自分に合った政策を選ぼう！
      </p>
      <p className="block w-fit mx-auto mb-10 px-3 py-1.5 text-2xl font-extrabold text-black text-center">
        政策マッチングを試してみてください
      </p>

      <button
        className="mt-6 px-6 py-3 text-2xl font-bold text-center text-white bg-green-500 border-2 border-transparent rounded-lg cursor-pointer transition-colors duration-200 hover:text-green-500 hover:bg-white hover:border-green-500"
        onClick={() => router.push("/quiz")}
      >
        START！
      </button>
    </div>
  );
}
