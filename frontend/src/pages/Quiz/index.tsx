// src/pages/Quiz/index.tsx
import React, { useState } from 'react';
import styles from './Quiz.module.css';
// グラフ生成用コンポーネント（後で実装）
import DebateGraph from '../../components/DebateGraph';

const TOTAL_QUESTIONS = 12;
const options = [
  '反対',
  'やや反対',
  '中立',
  'やや賛成',
  '賛成'
];

export default function Quiz() {
  /** 現在の質問番号 */
  const [current, setCurrent] = useState(0);
  /** ユーザーの回答 */
  const [answers, setAnswers] = useState<string[]>(Array(TOTAL_QUESTIONS).fill(''));
  /** グラフ生成フラグ */
  const [showGraph, setShowGraph] = useState(false);

  const handleAnswer = (choice: string) => {
    const newAnswers = [...answers];
    newAnswers[current] = choice;
    setAnswers(newAnswers);
    if (current < TOTAL_QUESTIONS - 1) {
      setCurrent(current + 1);
    }
    // 回答時にグラフをリセット
    setShowGraph(false);
  };

  const handleGenerate = () => {
    // ここでDebateGraph生成の処理を呼び出し
    setShowGraph(true);
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

      <h3 className={styles.subtitle}>政策に関する法律</h3>
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

      {/* グラフ生成ボタン */}
      <button className={styles.generateButton} onClick={handleGenerate}>
        グラフ生成
      </button>

      <div className={styles.graphContainer}>
        {showGraph ? (
          <DebateGraph data={answers} questionIndex={current} />
        ) : null}
      </div>
    </div>
  );
}
