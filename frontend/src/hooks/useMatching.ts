import { useState, useEffect } from "react";
import {
  matchingService,
  MatchRequest,
  MatchResponse,
} from "@/services/matchingService";

interface UseMatchingResult {
  matchResult: MatchResponse | null;
  loading: boolean;
  error: string | null;
}

export const useMatching = (
  electionId: number | null,
  request: MatchRequest | null
): UseMatchingResult => {
  const [matchResult, setMatchResult] = useState<MatchResponse | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (
      !electionId ||
      !request ||
      !request.answers.length ||
      !request.important_question_ids.length
    ) {
      return;
    }

    const fetchMatching = async () => {
      try {
        setLoading(true);
        setError(null);
        const result = await matchingService.getMatches(electionId, request);
        setMatchResult(result);
      } catch (err) {
        setError(
          err instanceof Error
            ? err.message
            : "マッチング結果の取得に失敗しました"
        );
      } finally {
        setLoading(false);
      }
    };

    fetchMatching();
  }, [electionId, request]);

  return { matchResult, loading, error };
};
