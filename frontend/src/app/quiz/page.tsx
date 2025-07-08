'use client';
// ------------------------------------------------------------
// 政党マッチングクイズ（質問表示＋回答送信）
// ------------------------------------------------------------
// ✅ 改善点
// 1) PascalCase / camelCase どちらのレスポンスでも動くようキー名を変換
// 2) fetch エラー時のメッセージ表示（CORS, 404 等）
// 3) mode:"cors" を明示
// 4) Optional chaining/ガードで API 異常時のクラッシュを防止
// ------------------------------------------------------------

import React, { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import styles from './page.module.css';

// --------------------------- 型定義 ---------------------------
interface Question {
  id: number;
  title: string;
  text: string;
  desc?: string | null;
}

interface MatchResult {
  top_match: {
    party_name: string;
    match_rate: number;
  };
  results: {
    party_name: string;
    match_rate: number;
  }[];
}

// 選択肢
const OPTIONS = [
  { label: '反対', value: -2 },
  { label: 'やや反対', value: -1 },
  { label: '中立', value: 0 },
  { label: 'やや賛成', value: 1 },
  { label: '賛成', value: 2 },
];

// API ベース URL
const API_BASE = process.env.NEXT_PUBLIC_API_BASE ?? 'http://localhost:8123';

export default function Quiz() {
  const router = useRouter();

  const [questions, setQuestions] = useState<Question[]>([]);
  const [answers, setAnswers] = useState<number[]>([]);
  const [currentIdx, setCurrentIdx] = useState(0);
  const [electionId, setElectionId] = useState<number | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  // ------------------ 最新選挙＆設問取得 ------------------
  useEffect(() => {
    (async () => {
      try {
        // ① 最新選挙
        const resElection = await fetch(`${API_BASE}/elections/latest`, { mode: 'cors' });
        if (!resElection.ok) throw new Error(`latest fetch → ${resElection.status}`);
        const latestJson = await resElection.json();
        const id: number | undefined = latestJson.ID ?? latestJson.id;
        if (!id) throw new Error('選挙IDが取得できません');
        setElectionId(id);

        // ② 設問
        const resQ = await fetch(`${API_BASE}/elections/${id}/questions`, { mode: 'cors' });
        if (!resQ.ok) throw new Error(`questions fetch → ${resQ.status}`);
        const qsJson = await resQ.json();
        if (!qsJson.questions || !Array.isArray(qsJson.questions) || qsJson.questions.length === 0) {
          throw new Error('設問が存在しません');
        }
        const qs: Question[] = qsJson.questions.map((q: any): Question => ({
          id: q.ID ?? q.id,
          title: q.Title ?? q.title,
          text: q.QuestionText ?? q.question_text,
          desc: q.Description ?? q.description ?? null,
        }));
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
    const newAns = [...answers];
    newAns[currentIdx] = val;
    setAnswers(newAns);

    if (currentIdx < questions.length - 1) {
      setCurrentIdx(currentIdx + 1);
      return;
    }

    // ---- 全質問終了: マッチングAPIへ ----
    if (electionId == null) {
      setError('選挙ID未取得');
      return;
    }

    const payload = {
      answers: questions.map((q, idx) => ({
        question_id: q.id,
        answer: newAns[idx],
      })),
      important_question_ids: [], // 今後のUI拡張用
    };

    try {
      const res = await fetch(`${API_BASE}/elections/${electionId}/matches`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload),
        mode: 'cors',
      });
      if (!res.ok) throw new Error(`matches fetch → ${res.status}`);
      const matchJson: MatchResult = await res.json();
      sessionStorage.setItem('matchResult', JSON.stringify(matchJson));
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

      {/* 設問 */}
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
