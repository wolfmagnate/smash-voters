// src/lib/api.ts
// ------------------------------------------------------------
// API ラッパ – 政党マッチング
// HTTP 通信を担当
// ------------------------------------------------------------

import { Question, MatchResult } from '@/types/matching';

const API_BASE = process.env.NEXT_PUBLIC_API_BASE ?? 'http://localhost:8123';

// ------------------ 最新選挙ID取得 ------------------
export async function fetchLatestElectionId(): Promise<number> {
  const res = await fetch(`${API_BASE}/elections/latest`, { mode: 'cors' });
  if (!res.ok) throw new Error(`latest fetch → ${res.status}`);
  const data = await res.json();
  const id: number | undefined = data.ID ?? data.id;
  if (!id) throw new Error('選挙IDが取得できません');
  return id;
}

// ------------------ 設問一覧取得 ------------------
export async function fetchQuestions(electionId: number): Promise<Question[]> {
  const res = await fetch(`${API_BASE}/elections/${electionId}/questions`, { mode: 'cors' });
  if (!res.ok) throw new Error(`questions fetch → ${res.status}`);
  const json = await res.json();
  if (!json.questions || !Array.isArray(json.questions) || json.questions.length === 0) {
    throw new Error('設問が存在しません');
  }
  return json.questions.map((q: any): Question => ({
    id: q.ID ?? q.id,
    title: q.Title ?? q.title,
    text: q.QuestionText ?? q.question_text,
    desc: q.Description ?? q.description ?? null,
  }));
}

// ------------------ マッチング計算 ------------------
export async function postMatch(
  electionId: number,
  answers: { question_id: number; answer: number }[],
): Promise<MatchResult> {
  const payload = { answers, important_question_ids: [] };
  const res = await fetch(`${API_BASE}/elections/${electionId}/matches`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload),
    mode: 'cors',
  });
  if (!res.ok) throw new Error(`matches fetch → ${res.status}`);
  return res.json();
}
