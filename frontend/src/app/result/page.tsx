'use client';
// src/app/result/page.tsx
// ------------------------------------------------------------
// Result UI – 政党マッチング
// ------------------------------------------------------------

import React, { useEffect, useState } from 'react';
import { MatchResult } from '@/types/matching';
const [data, setData] = useState<MatchResult | null>(null);
'use client';
// src/app/result/page.tsx
// ------------------------------------------------------------
// Result UI – 政党マッチング
// ------------------------------------------------------------

import React, { useEffect, useState } from 'react';
import { MatchResult } from '@/types/matching';
const [data, setData] = useState<MatchResult | null>(null);

export default function Result() {
  const [data, setData] = useState(null);

  useEffect(() => {
    const raw = sessionStorage.getItem('matchResult');
    if (raw) setData(JSON.parse(raw));
  }, []);

  if (!data) return <p>結果を読み込み中…</p>;

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
