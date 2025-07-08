'use client';
// src/pages/result/index.tsx
// ------------------------------------------------------------
// 政党マッチング – Result ページ
// ------------------------------------------------------------

import React, { useMemo } from 'react';
import { useRouter, useSearchParams } from 'next/navigation';
import { useLatestElection } from '@/hooks/useLatestElection';
import { useMatching } from '@/hooks/useMatching';
import { UserAnswer } from '@/services/matchingService';
import { Spinner } from '@/components/ui/Spinner';
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  Tooltip,
  ResponsiveContainer,
  LabelList,
  Cell,
} from 'recharts';

// ------------------ モックデータ ------------------
const SAMPLE_RESULT: MatchResult = {
  top_match: { party_name: '立憲民主党', match_rate: 78 },
  results: [
    { party_name: '立憲民主党', match_rate: 78 },
    { party_name: 'みんなでつくる党', match_rate: 73 },
    { party_name: '公明党', match_rate: 72 },
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

// ------------------ 政党カラー対応表 ------------------
const PARTY_COLOR_MAP: Record<string, string> = {
  立憲民主党: '#184589',
  みんなでつくる党: '#F8EA0D',
  公明党: '#F55881',
  再生の道: '#5E005E',
  国民民主党: '#F8BC00',
  日本維新の会: '#6FBA2C',
  日本共産党: '#DB001C',
  れいわ新選組: '#E4027E',
  自由民主党: '#3CA324',
  社会民主党: '#01A8EC',
  日本保守党: '#0A82DC',
  無所属: '#DFDFDF',
  参政党: '#D85D0F',
  諸派: '#D3D3D3',
  NHK党: '#ffef00',
};


interface PartyStat {
  party_name: string;
  match_rate: number;
}

export default function Result() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const { election } = useLatestElection();

  /* -------- URL パラメータ → リクエスト -------- */
  const matchRequest = useMemo(() => {
    const answersParam = searchParams.get('answers');
    const importantParam = searchParams.get('important');
    if (!answersParam || !importantParam) return null;
    try {
      const answers: UserAnswer[] = JSON.parse(answersParam);
      const important_question_ids = importantParam.split(',').map(Number);
      return { answers, important_question_ids };
    } catch {
      return null;
    }
  }, [searchParams]);

  /* -------- API 呼出 -------- */
  const { matchResult, loading, error } = useMatching(
    election?.id ?? null,
    matchRequest,
  );

  const dataToShow: MatchResult = matchResult ?? SAMPLE_RESULT;
  const sortedResults: PartyStat[] = [...dataToShow.results]
    .map((r) => ({ ...r, party_name: r.party_name.trim() }))
    .sort((a, b) => b.match_rate - a.match_rate);

  /* -------- 早期リターン -------- */
  if (!matchRequest)
    return <Error message="回答データが見つかりません" onRetry={() => router.push('/quiz')} />;
  if (loading)
    return (
      <div className="max-w-4xl mx-auto p-4 text-center">
        <Spinner />
        <p className="text-xl mt-4">マッチング結果を計算中...</p>
      </div>
    );
  if (error) return <Error message={error} onRetry={() => router.push('/quiz')} />;

  const best = dataToShow.top_match;

  return (
    <div className="max-w-4xl mx-auto p-4">
      {/* Top Match */}
      <section className="text-center space-y-6 mb-10">
        <h1 className="text-4xl font-bold">マッチング結果</h1>
        <div className="bg-gradient-to-r from-blue-500 to-purple-600 text-white p-8 rounded-xl inline-block shadow-lg">
          <h2 className="text-2xl font-bold mb-2">あなたに最も適した政党</h2>
          <h3 className="text-3xl font-bold mb-1">{best.party_name}</h3>
          <p className="text-xl">マッチ度: {best.match_rate}%</p>
        </div>
      </section>

      {/* 個人結果グラフ */}
      <section className="mb-16">
        <h2 className="text-2xl font-bold mb-4 text-center">全政党との適合度</h2>
        <ResponsiveContainer width="100%" height={700}>
          <BarChart data={sortedResults} layout="vertical" margin={{ top: 40, right: 20, left: 20, bottom: 5 }} barCategoryGap={5}>
            <XAxis type="number" domain={[0, 100]} hide />
            <YAxis interval={0} type="category" dataKey="party_name" width={180} />
            <Tooltip formatter={(v: number) => `${v}%`} />
            <Bar dataKey="match_rate">
              {sortedResults.map((r, idx) => (
                <Cell key={idx} fill={PARTY_COLOR_MAP[r.party_name] ?? '#8884d8'}/>
              ))}
              <LabelList dataKey="match_rate" position="right" formatter={(v: number) => `${v}%`} />
            </Bar>
          </BarChart>
        </ResponsiveContainer>
      </section>

      {/* アクション */}
      <div className="text-center">
        <button onClick={() => router.push('/quiz')} className="px-6 py-3 bg-green-500 text-white rounded-lg hover:bg-green-600 mr-4">もう一度クイズを受ける</button>
        <button onClick={() => router.push('/')} className="px-6 py-3 bg-gray-500 text-white rounded-lg hover:bg-gray-600">ホームに戻る</button>
      </div>
    </div>
  );
}

/* -------- 補助 -------- */
function Error({ message, onRetry }: { message: string; onRetry: () => void }) {
  return (
    <div className="max-w-4xl mx-auto p-4 text-center">
      <h1 className="text-3xl font-bold mb-6">エラー</h1>
      <p className="text-xl text-red-500 mb-4">{message}</p>
      <button onClick={onRetry} className="px-6 py-3 bg-blue-500 text-white rounded-lg hover:bg-blue-600">クイズをやり直す</button>
    </div>
  );
}
