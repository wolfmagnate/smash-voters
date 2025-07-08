import { useState, useEffect } from "react";
import { questionService, Question } from "../services/questionService";

interface UseQuestionsResult {
  questions: Question[];
  loading: boolean;
  error: string | null;
}

export const useQuestions = (electionId: number | null): UseQuestionsResult => {
  const [questions, setQuestions] = useState<Question[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!electionId) {
      setLoading(false);
      return;
    }

    const fetchQuestions = async () => {
      try {
        setLoading(true);
        setError(null);
        const data = await questionService.getQuestions(electionId);
        setQuestions(data);
      } catch (err) {
        setError(
          err instanceof Error ? err.message : "設問の取得に失敗しました"
        );
      } finally {
        setLoading(false);
      }
    };

    fetchQuestions();
  }, [electionId]);

  return { questions, loading, error };
};
