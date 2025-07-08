export interface Question {
  id: number;
  title: string;
  question_text: string;
  description: string;
}

interface APIQuestion {
  ID: number;
  Title: string;
  QuestionText: string;
  Description: string | null;
}

export interface QuestionsResponse {
  questions: APIQuestion[];
}

export const questionService = {
  async getQuestions(electionId: number): Promise<Question[]> {
    const response = await fetch(`/api/elections/${electionId}/questions`);

    if (!response.ok) {
      if (response.status === 404) {
        throw new Error("指定されたIDの選挙、または設問が見つかりません");
      }
      throw new Error(`設問の取得に失敗しました: ${response.status}`);
    }
    const data: QuestionsResponse = await response.json();

    // APIレスポンスのデータ形式を変換
    return data.questions.map((q) => ({
      id: q.ID,
      title: q.Title,
      question_text: q.QuestionText,
      description: q.Description || "", // nullの場合は空文字列に変換
    }));
  },
};
