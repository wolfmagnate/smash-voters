export interface Election {
  id: number;
  name: string;
}

export const electionService = {
  async getLatestElection(): Promise<Election> {
    const response = await fetch("/api/elections/latest");

    if (!response.ok) {
      if (response.status === 404) {
        throw new Error("アクティブな選挙情報が見つかりません");
      }
      throw new Error(`選挙情報の取得に失敗しました: ${response.status}`);
    }

    const data = await response.json();
    return {
      id: data.ID,
      name: data.Name,
    };
  },
};
