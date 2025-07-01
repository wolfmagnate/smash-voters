'use client';
import React, { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import styles from './page.module.css';

export default function HomePage() {
  const router = useRouter();
  const [daysLeft, setDaysLeft] = useState(0);

  useEffect(() => {
    const today = new Date();
    const electionDate = new Date(2025, 6, 31);
    const diffMs = electionDate.getTime() - today.getTime();
    const days = Math.ceil(diffMs / (1000 * 60 * 60 * 24));
    setDaysLeft(days > 0 ? days : 0);
  }, []);

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Policy Evaluation System</h1>

      <div className={styles.notice}>
        次回の投票日は7月31日<br />あと{daysLeft}日
      </div>

      <p className={styles.subtitle}>利用ロジック図で自分に合った政策を選ぼう。</p>
      <p className={styles.subtitle}>政策マッチングを試してみてください。</p>

      <button className={styles.startButton} onClick={() => router.push('/quiz')}>
        START！
      </button>
    </div>
  );
}
