'use client';
// src/app/result/page.tsx
// ------------------------------------------------------------
// Result UI – 政党マッチング (モックデータ対応版)
// ------------------------------------------------------------

import React, { useEffect, useState } from 'react';
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
  LabelList,
} from 'recharts';

import { MatchResult } from '@/types/matching';

// ------------------ モックデータ ------------------
const SAMPLE_RESULT: MatchResult = {
  top_match: { party_name: '立憲民主党', match_rate: 78 },
  results: [
    { party_name: '立憲民主党', match_rate: 78 },
    { party_name: 'みんなでつくる党', match_rate: 73 },
    { party_name: ' 公明党', match_rate: 72 },
    { party_name: '再生の道', match_rate: 71 },
    { party_name: '国民民主党', match_rate: 67 },
    { party_name: '日本維新の会', match_rate: 66 },
    { party_name: '日本共産党', match_rate: 65 },
    { party_name: 'れいわ新選組', match_rate: 63 },
    { party_name: '自由民主党', match_rate: 61 },
    { party_name: '社会民主党', match_rate: 60 },
    { party_name: '日本保守党', match_rate: 57 },
    { party_name: '無所属連合', match_rate: 57 },
    { party_name: '参政党', match_rate: 55 },
    { party_name: 'NHK党', match_rate: 52 },
  ],
};

// 全ユーザー統計も同じ配列を使ってデモ
const SAMPLE_STATS = SAMPLE_RESULT.results;

interface PartyStat {
  party_name: string;
  match_rate: number;
}

export default function Result() {
  const [data, setData] = useState<MatchResult | null>(null);
  const [stats, setStats] = useState<PartyStat[] | null>(null);

  useEffect(() => {
    // 個人分
    const raw = sessionStorage.getItem('matchResult');
    if (raw) {
      setData(JSON.parse(raw));
    } else {
      // セッションに無い場合はモックを使用
      setData(SAMPLE_RESULT);
    }

    // 全ユーザ統計（モック）
    const agg = sessionStorage.getItem('matchStats');
    if (agg) {
      setStats(JSON.parse(agg));
    } else {
      setStats(SAMPLE_STATS);
    }
  }, []);

  if (!data) return <p className="text-center mt-10">結果を読み込み中…</p>;

  const best = data.top_match;

  return (
    <div className="max-w-4xl mx-auto p-4">
      {/* ------- matching result (個人) ------- */}
      <section className="text-center space-y-6">
        <h1 className="font-bold text-3xl mb-2">政策マッチング結果！</h1>
        <p className="text-xl">あなたの考えに最も近い政党は…</p>

        <div className="bg-gray-100 rounded-2xl p-6 shadow-md inline-block">
          <h2 className="text-2xl font-semibold text-indigo-700 mb-2">{best.party_name}</h2>
          <p className="text-xl">マッチング度</p>
          <p className="text-5xl font-bold text-indigo-600">
            {best.match_rate}
            <small className="text-2xl ml-1">%</small>
          </p>
        </div>
        <p className="text-xl">でした</p>
      </section>

      {/* ------- matching result details (個人) ------- */}
      <section className="mt-12">
        <h3 className="text-xl font-semibold mb-4 text-center">あなたの全政党マッチング度</h3>
        <ResponsiveContainer width="100%" height={300}>
          <BarChart data={data.results} layout="vertical" margin={{ top: 5, right: 30, left: 20, bottom: 5 }}>
            <CartesianGrid strokeDasharray="3 3" />
            <XAxis type="number" domain={[0, 100]} hide />
            <YAxis type="category" dataKey="party_name" width={120} />
            <Tooltip formatter={(v: number) => `${v}%`} />
            <Bar dataKey="match_rate" fill="#6366f1">
              <LabelList dataKey="match_rate" position="right" formatter={(v: number) => `${v}%`} />
            </Bar>
          </BarChart>
        </ResponsiveContainer>
      </section>

      {/* ------- all users stats ------- */}
      {stats && (
        <section className="mt-16">
          <h3 className="text-xl font-semibold mb-4 text-center">みんなのマッチング結果</h3>
          <ResponsiveContainer width="100%" height={300}>
            <BarChart data={stats} layout="vertical" margin={{ top: 5, right: 30, left: 20, bottom: 5 }}>
              <CartesianGrid strokeDasharray="3 3" />
              <XAxis type="number" domain={[0, 100]} hide />
              <YAxis type="category" dataKey="party_name" width={120} />
              <Tooltip formatter={(v: number) => `${v}%`} />
              <Bar dataKey="match_rate" fill="#10b981">
                <LabelList dataKey="match_rate" position="right" formatter={(v: number) => `${v}%`} />
              </Bar>
            </BarChart>
          </ResponsiveContainer>
        </section>
      )}
    </div>
  );
}
