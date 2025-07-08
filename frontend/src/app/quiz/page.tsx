// src/pages/quiz/index.tsx
"use client";
import React, { useState } from "react";
import { useRouter } from "next/navigation";
import Graph from "../../components/graph/ui/graph";
import SelectApproveOrReject from "./_components/SelectApproveOrReject";

const TOTAL_QUESTIONS = 20;

export default function Quiz() {
  const router = useRouter();
  const [current, setCurrent] = useState(0);
  const [answers, setAnswers] = useState<string[]>(
    Array(TOTAL_QUESTIONS).fill("")
  );

  const handleAnswer = (choice: string) => {
    // 回答を保存
    const newAnswers = [...answers];
    newAnswers[current] = choice;
    setAnswers(newAnswers);
    // 次の質問へ、または結果ページへ移動
    if (current < TOTAL_QUESTIONS - 1) {
      setCurrent(current + 1);
    } else {
      //const query = encodeURIComponent(newAnswers.join(','));
      router.push("/result"); //?answers=${query}`);
    }
  };

  return (
    <div className="w-full max-w-7xl mx-auto p-4 bg-transparent">
      <div className="relative h-6 bg-gray-300 rounded-xl overflow-hidden mt-auto mb-10">
        <div
          className="h-full bg-green-500 rounded-xl transition-all duration-300 ease-in-out"
          style={{ width: `${((current + 1) / TOTAL_QUESTIONS) * 100}%` }}
        />
        <span className="absolute top-0 left-1/2 transform -translate-x-1/2 font-bold leading-6 text-black">{`${
          current + 1
        } / ${TOTAL_QUESTIONS}`}</span>
      </div>

      <h3 className="block w-fit mx-auto mb-5 px-3 py-1.5 bg-yellow-200 rounded-3xl text-lg font-extrabold text-black text-center">
        {" "}
        政策に関する法律
      </h3>
      <h1 className="text-3xl font-bold text-black mb-6 text-center">
        問題の説明
      </h1>
      <p className="text-xl text-black leading-relaxed mb-10 text-center">
        ここに政策に関する詳細な説明が入ります。ユーザーに問題の背景や要点を伝えます。
      </p>

      <SelectApproveOrReject onSelect={handleAnswer} className="mb-10" />
      <Graph />
    </div>
  );
}
