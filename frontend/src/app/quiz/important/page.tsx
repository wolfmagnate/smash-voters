"use client";
import React, { useState } from "react";
import { useRouter, useSearchParams } from "next/navigation";
import { useQuestions } from "../hooks/useQuestions";
import { useLatestElection } from "@/hooks/useLatestElection";
import { Spinner } from "@/components/ui/Spinner";

export default function ImportantQuestions() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const { election } = useLatestElection();
  const { questions, loading, error } = useQuestions(election?.id || null);
  const [selectedQuestions, setSelectedQuestions] = useState<number[]>([]);

  const answers = searchParams.get("answers");

  const handleQuestionToggle = (questionId: number) => {
    setSelectedQuestions((prev) => {
      if (prev.includes(questionId)) {
        return prev.filter((id) => id !== questionId);
      } else if (prev.length < 3) {
        return [...prev, questionId];
      }
      return prev;
    });
  };

  const handleSubmit = () => {
    if (selectedQuestions.length === 3) {
      const queryParams = new URLSearchParams({
        answers: answers || "",
        important: selectedQuestions.join(","),
      });
      router.push(`/result?${queryParams.toString()}`);
    }
  };

  if (loading || !questions || questions.length === 0 || !answers) {
    return <Spinner />;
  }

  if (error) {
    return (
      <div className="w-full max-w-7xl mx-auto p-4 text-center">
        <p className="text-xl text-red-500">{error}</p>
      </div>
    );
  }

  return (
    <div className="w-full max-w-4xl mx-auto px-4 py-8">
      <h1 className="text-3xl font-bold text-black mb-6 text-center">
        重要項目を3つ選択してください
      </h1>

      <p className="text-xl text-gray-600 mb-8 text-center">
        あなたが特に重視する政策分野を3つ選んでください。
        選択した項目はマッチング結果により強く反映されます。
      </p>

      <div className="space-y-4 mb-8">
        {questions.map((question) => (
          <button
            key={question.id}
            onClick={() => handleQuestionToggle(question.id)}
            className={`w-full p-4 rounded-lg border-2 text-left transition-all duration-200 ${
              selectedQuestions.includes(question.id)
                ? "border-blue-500 bg-blue-50"
                : "border-gray-300 hover:border-gray-400"
            }`}
            disabled={
              !selectedQuestions.includes(question.id) &&
              selectedQuestions.length >= 3
            }
          >
            <h3 className="text-lg font-bold mb-2">{question.title}</h3>
            <p className="text-gray-600">{question.question_text}</p>
          </button>
        ))}
      </div>

      <div className="text-center">
        <p className="mb-4 text-lg">選択済み: {selectedQuestions.length} / 3</p>

        <button
          onClick={handleSubmit}
          disabled={selectedQuestions.length !== 3}
          className={`px-8 py-3 rounded-lg text-xl font-bold transition-all duration-200 ${
            selectedQuestions.length === 3
              ? "bg-green-500 text-white hover:bg-green-600"
              : "bg-gray-300 text-gray-500 cursor-not-allowed"
          }`}
        >
          結果を見る
        </button>
      </div>
    </div>
  );
}
