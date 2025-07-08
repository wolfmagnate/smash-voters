"use client";
import React, { useState, useEffect } from "react";
import { useRouter } from "next/navigation";

export default function HomePage() {
  const router = useRouter();
  const [daysLeft, setDaysLeft] = useState(0);

  useEffect(() => {
    const today = new Date();
    const electionDate = new Date(2025, 6, 29); // ← 6 ＝ 7
    const diffMs = electionDate.getTime() - today.getTime();
    const days = Math.ceil(diffMs / (1000 * 60 * 60 * 24));
    setDaysLeft(days > 0 ? days : 0);
  }, []);

  return (
    <div className="w-full max-w-4xl mx-auto px-4 py-8 text-center bg-white">
      <h1 className="text-5xl font-bold text-black mb-15">
        Policy Evaluation System（仮）
      </h1>

      <div className="max-w-4xl border-t-[5px] border-b-[5px] border-gray-300 py-6 mx-0 my-4 text-2xl font-extrabold text-gray-700 mb-15">
        {daysLeft === 0 ? (
          <>今日は投票日です！</>
        ) : (
          <>
            次回の投票日は7月29日（仮）
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
