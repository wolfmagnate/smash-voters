// src/pages/result/index.tsx
'use client';
import React from 'react';
import styles from './result.module.css';

export default function Result() {
  return (
    <div className={styles.container}>
      <h1 className={styles.title}>結果ページ</h1>
      <p className={styles.text}>政策マッチングの結果を表示します。</p>
      {/* TODO: ここにマッチ度が最も高い政党を表示 */}
    </div>
  );
}