// src/pages/result/index.tsx
'use client';
import React, { useEffect, useState } from 'react';

export default function Result() {
  const [data, setData] = useState(null);

  useEffect(() => {
    const raw = sessionStorage.getItem('matchResult');
    if (raw) setData(JSON.parse(raw));
  }, []);

  if (!data) return <p>結果を読み込み中…</p>;

  return (
    <div>
      <h1>あなたと最も一致する政党</h1>
      <h2>{data.top_match.party_name} ({data.top_match.match_rate}%)</h2>

      <h3>全結果</h3>
      <ul>
        {data.results.map((r: any) => (
          <li key={r.party_name}>
            {r.party_name}: {r.match_rate}%
          </li>
        ))}
      </ul>
    </div>
  );
}
