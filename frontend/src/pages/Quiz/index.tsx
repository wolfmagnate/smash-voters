// src/pages/Quiz/index.tsx
'use client';
import React, { useState } from 'react';
import { useRouter } from 'next/router';
import styles from './Quiz.module.css';
//import DebateGraph from '../../components/DebateGraph';

const TOTAL_QUESTIONS = 20;
const options = ['反対','やや反対','中立','やや賛成','賛成'];

export default function Quiz() {
  //const router = useRouter();
  const [current, setCurrent] = useState(0);
  const [answers, setAnswers] = useState<string[]>(Array(TOTAL_QUESTIONS).fill(''));
  const [showGraph, setShowGraph] = useState(false);

  const handleAnswer = (choice: string) => {
    // 回答を保存
    const newAnswers = [...answers];
    newAnswers[current] = choice;
    setAnswers(newAnswers);
    // グラフ表示をリセット
    setShowGraph(false);
    // 次の質問へ、または結果ページへ移動
    if (current < TOTAL_QUESTIONS - 1) {
      setCurrent(current + 1);
    } else {
      const query = encodeURIComponent(newAnswers.join(','));
      //router.push(`/result?answers=${query}`);
    }
  };

  return (
    <div className={styles.container}>
      <div className={styles.progressBar}>
        <div
          className={styles.filled}
          style={{ width: `${((current + 1) / TOTAL_QUESTIONS) * 100}%` }}
        />
        <span className={styles.progressText}>{`${current + 1} / ${TOTAL_QUESTIONS}`}</span>
      </div>

      <h3 className={styles.subtitle}> 政策に関する法律</h3>
      <h1 className={styles.title}>問題の説明</h1>
      <p className={styles.description}>
        ここに政策に関する詳細な説明が入ります。ユーザーに問題の背景や要点を伝えます。
      </p>

      <div className={styles.buttonGroup}>
        {options.map((opt) => (
          <button
            key={opt}
            className={styles.button}
            onClick={() => handleAnswer(opt)}
          >
            {opt}
          </button>
        ))}
      </div>

      <button className={styles.generateButton}>
        グラフ生成
      </button>

      <div className={styles.graphContainer}>
        {showGraph}
      </div>
    </div>
  );
}
