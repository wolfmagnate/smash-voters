// src/types/matching.ts
// --------------------------- 型定義 ---------------------------

export interface Question {
    id: number;
    title: string;
    text: string;
    desc?: string | null;
  }
  
  export interface MatchResult {
    top_match: {
      party_name: string;
      match_rate: number;
    };
    results: {
      party_name: string;
      match_rate: number;
    }[];
  }
  