import { useState, useEffect } from "react";
import { electionService, Election } from "../services/electionService";

interface UseLatestElectionResult {
  election: Election | null;
  loading: boolean;
  error: string | null;
}

export const useLatestElection = (): UseLatestElectionResult => {
  const [election, setElection] = useState<Election | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchLatestElection = async () => {
      try {
        setLoading(true);
        setError(null);
        const data = await electionService.getLatestElection();
        setElection(data);
      } catch (err) {
        setError(
          err instanceof Error ? err.message : "選挙情報の取得に失敗しました"
        );
      } finally {
        setLoading(false);
      }
    };

    fetchLatestElection();
  }, []);

  return { election, loading, error };
};
