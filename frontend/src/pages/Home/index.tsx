// src/pages/home/index.tsx
'use client';
import React, { useState } from 'react';
import { useRouter } from 'next/navigation';
import styles from './home.module.css';

export default function Home() {
    const router = useRouter();
  
    return (
      <div className={styles.container}>
        <h1 className={styles.title}>Policy Evaluation System</h1>
        <p className={styles.subtitle}>ロジック図を通して、自分に合った政策を選ぼう！</p>
        <p className={styles.subtitle}>政策マッチングを試してみてください</p>
        <button
          className={styles.startButton}
          onClick={() => router.push('/quiz')}
        >
          START！
        </button>
      </div>
    );
  }