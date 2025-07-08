/**
 * ディベートグラフのノード（主張）を表す型定義
 */
export interface DebateGraphNode {
  /** ノードの一意識別子 */
  id: string;
  /** 主張の内容 */
  argument: string;
  /** 主張がなぜ重要なのかを補強する論拠 */
  importance?: string[];
  /** その主張が、なぜ自分たちのプランで固有に発生するのかを補強する論拠 */
  uniqueness?: string[];
  /** このノード自体が、相手への反論として提示されたものかを示す */
  is_rebuttal: boolean;
}

/**
 * ディベートグラフのエッジ（因果関係）を表す型定義
 */
export interface DebateGraphEdge {
  /** 因果関係における「原因」となるノードのID */
  cause: string;
  /** 因果関係における「結果」となるノードのID */
  effect: string;
  /** なぜその因果関係が確実に発生するのかを補強する論拠 */
  certainty?: string[];
  /** この因果関係自体が、なぜ特定のプランでのみ強く成立するのかを補強する論拠 */
  uniqueness?: string[];
  /** このエッジ自体が相手の議論に対する反論の一部として提示されたものかを示す */
  is_rebuttal: boolean;
}

/**
 * ディベートグラフ全体を表す型定義
 */
export interface DebateGraph {
  /** グラフ内のすべてのノード */
  nodes: DebateGraphNode[];
  /** グラフ内のすべてのエッジ */
  edges: DebateGraphEdge[];
  /** ノードに対する反論 */
  node_rebuttals?: unknown[];
  /** エッジに対する反論 */
  edge_rebuttals?: unknown[];
  /** カウンター引数に対する反論 */
  counter_argument_rebuttals?: unknown[];
  /** ターン引数に対する反論 */
  turn_argument_rebuttals?: unknown[];
}
