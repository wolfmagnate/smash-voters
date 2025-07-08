// src/app/quiz/page.tsx
"use client";
import React, { useState } from "react";
import { useRouter } from "next/navigation";
import Graph from "../../components/graph/ui/graph";
import SelectApproveOrReject from "./_components/SelectApproveOrReject";
import { useQuestions } from "./hooks/useQuestions";
import { useLatestElection } from "@/hooks/useLatestElection";
import { Spinner } from "@/components/ui/Spinner";

export default function Quiz() {
  const router = useRouter();
  const [current, setCurrent] = useState(0);
  const [answers, setAnswers] = useState<number[]>([]);

  const { election } = useLatestElection();
  const { questions, loading, error } = useQuestions(election?.id || null);

  const handleAnswer = (choice: number) => {
    // 回答を保存
    const newAnswers = [...answers];
    newAnswers[current] = choice;
    setAnswers(newAnswers);

    // 次の質問へ、または重要項目選択ページへ移動
    if (current < questions.length - 1) {
      setCurrent(current + 1);
    } else {
      // 全問回答完了 - 重要項目選択ページへ
      const answersData = questions.map((q, index) => ({
        question_id: q.id,
        answer: newAnswers[index],
      }));
      const queryParams = new URLSearchParams({
        answers: JSON.stringify(answersData),
      });
      router.push(`/quiz/important?${queryParams.toString()}`);
    }
  };

  if (loading || !questions || questions.length === 0) {
    return <Spinner />;
  }

  if (error) {
    return (
      <div className="w-full max-w-7xl mx-auto p-4 text-center">
        <p className="text-xl text-red-500">{error}</p>
      </div>
    );
  }

  const currentQuestion = questions[current];

  return (
    <div className="w-full max-w-7xl mx-auto p-4">
      <div className="relative h-6 bg-gray-300 rounded-xl overflow-hidden mt-auto mb-10">
        <div
          className="h-full bg-green-500 rounded-xl transition-all duration-300 ease-in-out"
          style={{ width: `${((current + 1) / questions.length) * 100}%` }}
        />
        <span className="absolute top-0 left-1/2 transform -translate-x-1/2 font-bold leading-6 text-black">{`${
          current + 1
        } / ${questions.length}`}</span>
      </div>

      <h3 className="block w-fit mx-auto mb-5 px-3 py-1.5 bg-yellow-200 rounded-3xl text-lg font-extrabold text-black text-center">
        {currentQuestion.title}
      </h3>
      <h1 className="text-3xl font-bold text-black mb-6 text-center">
        {currentQuestion.question_text}
      </h1>
      <p className="text-xl text-black leading-relaxed mb-10 text-center">
        {currentQuestion.description}
      </p>

      <SelectApproveOrReject onSelect={handleAnswer} className="mb-10" />
      <Graph />
    </div>
  );
}
