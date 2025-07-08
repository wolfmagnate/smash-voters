/**
 * 主張の内容を表す型定義
 */
export interface ArgumentContent {
  /** 主張の文 */
  statement: string;
}

/**
 * エビデンスを表す型定義
 */
export interface Evidence {
  /** エビデンスのURL */
  url: string;
  /** エビデンスのタイトル */
  title: string;
}

/**
 * 重要性や固有性の論拠を表す型定義
 */
export interface Reasoning {
  /** 論拠の文 */
  statement: string;
  /** 論拠を裏付けるエビデンス */
  evidence?: Evidence[];
}

/**
 * ディベートグラフのノード（主張）を表す型定義
 */
export interface DebateGraphNode {
  /** ノードの一意識別子 */
  id: string;
  /** 主張の内容 */
  argument: ArgumentContent;
  /** 主張がなぜ重要なのかを補強する論拠 */
  importance?: Reasoning[];
  /** その主張が、なぜ自分たちのプランで固有に発生するのかを補強する論拠 */
  uniqueness?: Reasoning[];
  /** このノード自体が、相手への反論として提示されたものかを示す */
  is_rebuttal: boolean;
}

/**
 * ディベートグラフのエッジ（因果関係）を表す型定義
 */
export interface DebateGraphEdge {
  /** エッジの一意識別子 */
  id: string;
  /** 因果関係における「原因」となるノードのID */
  cause_id: string;
  /** 因果関係における「結果」となるノードのID */
  effect_id: string;
  /** なぜその因果関係が確実に発生するのかを補強する論拠 */
  certainty?: Reasoning[];
  /** この因果関係自体が、なぜ特定のプランでのみ強く成立するのかを補強する論拠 */
  uniqueness?: Reasoning[];
  /** このエッジ自体が相手の議論に対する反論の一部として提示されたものかを示す */
  is_rebuttal: boolean;
}

/**
 * ノードに対する反論を表す型定義
 */
export interface NodeRebuttal {
  /** 反論対象のノードID */
  target_node_id: string;
  /** 反論の種類（importance/uniqueness） */
  rebuttal_type: "importance" | "uniqueness";
  /** 反論を含むノードのID */
  rebuttal_node_id: string;
}

/**
 * エッジに対する反論を表す型定義
 */
export interface EdgeRebuttal {
  /** 反論対象のエッジID */
  target_edge_id: string;
  /** 反論の種類（certainty/uniqueness） */
  rebuttal_type: "certainty" | "uniqueness";
  /** 反論を含むノードのID */
  rebuttal_node_id: string;
}

/**
 * カウンター引数に対する反論を表す型定義
 */
export interface CounterArgumentRebuttal {
  /** 反論を含むノードのID */
  rebuttal_node_id: string;
  /** 反論対象のノードID */
  target_node_id: string;
}

/**
 * ターン引数に対する反論を表す型定義
 */
export interface TurnArgumentRebuttal {
  /** 反論を含むノードのID */
  rebuttal_node_id: string;
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
  node_rebuttals: NodeRebuttal[];
  /** エッジに対する反論 */
  edge_rebuttals: EdgeRebuttal[];
  /** カウンター引数に対する反論 */
  counter_argument_rebuttals: CounterArgumentRebuttal[];
  /** ターン引数に対する反論 */
  turn_argument_rebuttals: TurnArgumentRebuttal[];
}
