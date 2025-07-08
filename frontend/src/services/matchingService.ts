export interface UserAnswer {
  question_id: number;
  answer: number;
}

export interface MatchRequest {
  answers: UserAnswer[];
  important_question_ids: number[];
}

export interface PartyMatchResult {
  party_name: string;
  match_rate: number;
}

export interface MatchResponse {
  top_match: PartyMatchResult;
  results: PartyMatchResult[];
}

export const matchingService = {
  async getMatches(
    electionId: number,
    request: MatchRequest
  ): Promise<MatchResponse> {
    const response = await fetch(`/api/elections/${electionId}/matches`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(request),
    });

    if (!response.ok) {
      if (response.status === 400) {
        throw new Error("リクエストボディが不正です");
      }
      if (response.status === 404) {
        throw new Error("指定されたIDの選挙が見つかりません");
      }
      throw new Error(`マッチング計算に失敗しました: ${response.status}`);
    }

    return await response.json();
  },
};
