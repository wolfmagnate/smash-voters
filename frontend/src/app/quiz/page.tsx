'use client';
// src/app/quiz/page.tsx
// ------------------------------------------------------------
// Quiz UI – 政党マッチング
// ------------------------------------------------------------

import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import styles from './page.module.css';

import { Question } from '@/types/matching';
import { fetchLatestElectionId, fetchQuestions, postMatch } from '@/lib/api';

// 回答オプション
const OPTIONS = [
  { label: '反対', value: -2 },
  { label: 'やや反対', value: -1 },
  { label: '中立', value: 0 },
  { label: 'やや賛成', value: 1 },
  { label: '賛成', value: 2 },
];

export default function Quiz() {
  const router = useRouter();

  const [questions, setQuestions] = useState<Question[]>([]);
  const [answers, setAnswers] = useState<number[]>([]);
  const [currentIdx, setCurrentIdx] = useState(0);
  const [electionId, setElectionId] = useState<number | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  // ------------------ データ取得 ------------------
  useEffect(() => {
    (async () => {
      try {
        const id = await fetchLatestElectionId();
        setElectionId(id);
        const qs = await fetchQuestions(id);
        setQuestions(qs);
        setAnswers(Array(qs.length).fill(NaN));
      } catch (e: any) {
        console.error(e);
        setError(e.message ?? 'API Error');
      } finally {
        setLoading(false);
      }
    })();
  }, []);

  // ------------------ 回答処理 ------------------
  const handleAnswer = async (val: number) => {
    // 回答を保存
    const newAns = [...answers];
    newAns[currentIdx] = val;
    setAnswers(newAns);

    // まだ次の設問がある場合
    if (currentIdx < questions.length - 1) {
      setCurrentIdx(currentIdx + 1);
      return;
    }

    // 全問終了 → マッチングAPI へ
    if (electionId == null) {
      setError('選挙ID未取得');
      return;
    }

    try {
      const result = await postMatch(
        electionId,
        questions.map((q, idx) => ({ question_id: q.id, answer: newAns[idx] })),
      );
      sessionStorage.setItem('matchResult', JSON.stringify(result));
      router.push('/result');
    } catch (e: any) {
      console.error(e);
      setError(e.message ?? 'マッチングAPIエラー');
    }
  };

  // ------------------ UI ------------------
  if (loading) return <p className={styles.loading}>読み込み中...</p>;
  if (error) return <p className={styles.error}>{error}</p>;
  if (!questions.length) return <p className={styles.loading}>設問がありません</p>;

  const q = questions[currentIdx];

  return (
    <div className={styles.container}>
      {/* 進捗バー */}
      <div className={styles.progressBar}>
        <div
          className={styles.filled}
          style={{ width: `${((currentIdx + 1) / questions.length) * 100}%` }}
        />
        <span className={styles.progressText}>{`${currentIdx + 1} / ${questions.length}`}</span>
      </div>

      {/* 設問表示 */}
      <h3 className={styles.subtitle}>{q.title}</h3>
      <h1 className={styles.title}>{q.text}</h1>
      {q.desc && <p className={styles.description}>{q.desc}</p>}

      {/* 回答ボタン */}
      <div className={styles.buttonGroup}>
        {OPTIONS.map((opt) => (
          <button
            key={opt.label}
            className={styles.button}
            onClick={() => handleAnswer(opt.value)}
          >
            {opt.label}
          </button>
        ))}
      </div>
    </div>
  );
}
