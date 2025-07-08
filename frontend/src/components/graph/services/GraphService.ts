import { DebateGraph } from "../../../types/debate-graph";

const sampleData: DebateGraph = {
  nodes: [
    {
      id: "node-1",
      argument: {
        statement: "脱原発政策を推進し再生可能エネルギーへ転換する",
      },
      is_rebuttal: false,
      importance: [
        {
          statement:
            "気候変動対策として国際的に求められる脱炭素化の実現に不可欠な政策である。",
        },
        {
          statement:
            "福島第一原発事故の教訓を踏まえ、国民の安全と安心を最優先にしたエネルギー政策の転換である。",
        },
      ],
      uniqueness: [
        {
          statement:
            "日本特有の地震・津波リスクを考慮した、世界に先駆けた脱原発モデルの構築が可能である。",
        },
      ],
    },
    {
      id: "node-2",
      argument: {
        statement: "原子力発電を基幹電源として維持・再稼働する",
      },
      is_rebuttal: false,
      importance: [
        {
          statement:
            "安定した電力供給と経済性を両立させる現実的な選択肢である。",
        },
        {
          statement:
            "エネルギー安全保障の観点から、輸入依存度を下げる重要な手段である。",
        },
      ],
      uniqueness: [
        {
          statement:
            "日本の高度な原子力技術と安全規制は世界最高水準であり、安全な運転が可能である。",
        },
      ],
    },
    {
      id: "node-3",
      argument: {
        statement: "原子力発電所を段階的に停止する",
      },
      is_rebuttal: false,
      importance: [
        {
          statement:
            "大規模事故リスクの段階的軽減により、国民の安全を確保できる。",
        },
        {
          statement:
            "廃炉技術の開発と実証を並行して進めることで、将来の完全脱原発への道筋を作れる。",
          evidence: [
            {
              url: "https://www.jpx.co.jp/english/news/2021/03/20210325_01.html",
              title: "J-PARCの廃炉技術開発",
            },
          ],
        },
      ],
    },
    {
      id: "node-4",
      argument: {
        statement: "再生可能エネルギー導入への公的投資を拡大する",
      },
      is_rebuttal: false,
      importance: [
        {
          statement:
            "グリーン成長戦略の核となる投資で、経済成長と環境保護を両立させる。",
        },
        {
          statement:
            "技術革新を加速させ、再生可能エネルギーのコスト競争力を向上させる。",
        },
      ],
      uniqueness: [
        {
          statement:
            "日本の地理的条件を活かした地熱、海洋エネルギー、バイオマスなど多様な再生可能エネルギーの開発が可能である。",
        },
      ],
    },
    {
      id: "node-5",
      argument: {
        statement: "省エネルギー技術の開発・普及を国家戦略として推進する",
      },
      is_rebuttal: false,
      importance: [
        {
          statement:
            "エネルギー需要の抑制は、供給側の負荷を軽減し、脱原発を加速させる。",
        },
        {
          statement: "産業競争力の向上と環境負荷の低減を同時に実現できる。",
        },
      ],
      uniqueness: [
        {
          statement:
            "日本の製造業の技術力を活かした省エネ技術の輸出により、国際的な環境貢献と経済的利益を両立できる。",
        },
      ],
    },
    {
      id: "node-6",
      argument: {
        statement: "再生可能エネルギー関連産業が国内で成長する",
      },
      is_rebuttal: false,
      importance: [
        {
          statement: "新たな雇用創出と地域経済の活性化につながる。",
        },
        {
          statement: "エネルギー自給率の向上と産業構造の転換を実現する。",
        },
      ],
      uniqueness: [
        {
          statement:
            "太陽光パネル、蓄電池、水素技術など、日本の強みを活かした産業育成が可能である。",
        },
      ],
    },
    {
      id: "node-7",
      argument: {
        statement: "スマートグリッドと高性能蓄電技術が社会に普及する",
      },
      is_rebuttal: false,
      importance: [
        {
          statement:
            "再生可能エネルギーの不安定性を克服し、安定供給を実現する。",
        },
        {
          statement: "電力システムの効率化とコスト削減を実現する。",
        },
      ],
      uniqueness: [
        {
          statement:
            "日本のIT技術と電力技術の融合により、世界最先端のスマートグリッドシステムを構築できる。",
        },
      ],
    },
    {
      id: "node-8",
      argument: {
        statement: "電力需要の全体量が効率化により抑制される",
      },
      is_rebuttal: false,
      importance: [
        {
          statement:
            "国際エネルギー機関(IEA)は、エネルギー効率の改善を『第一の燃料(First Fuel)』と位置づけており、最もクリーンで安価なエネルギー供給源であると評価している。需要の抑制は、新たな発電所建設の必要性を低減させる。",
        },
      ],
      uniqueness: [
        {
          statement:
            "日本の省エネ技術と国民の省エネ意識の高さを活かした、世界に類を見ない効率的なエネルギー社会の実現が可能である。",
        },
      ],
    },
    {
      id: "node-9",
      argument: {
        statement: "グリーン分野で質の高い新たな雇用が創出される",
      },
      is_rebuttal: false,
      importance: [
        {
          statement:
            "IRENAの予測では、野心的なエネルギー転換シナリオの下で、2050年までに世界のエネルギー分野の雇用は1億3900万人に達し、そのうち4300万人が再生可能エネルギー分野とされている。",
        },
      ],
      uniqueness: [
        {
          statement:
            "原子力産業の雇用が特定の地域・専門分野に集中するのに対し、再生可能エネルギー関連の雇用は、建設、製造、メンテナンス、ソフトウェア開発など多岐にわたり、全国に分散して創出される。",
        },
      ],
    },
    {
      id: "node-10",
      argument: {
        statement: "エネルギー供給システムが分散化・強靭化する",
      },
      is_rebuttal: false,
      importance: [
        {
          statement:
            "災害時の電力供給の安定性が向上し、社会インフラの強靭性が高まる。",
        },
        {
          statement: "地域特性に応じた最適なエネルギー供給が可能になる。",
        },
      ],
      uniqueness: [
        {
          statement:
            "日本の多様な地理的条件と地域特性を活かした、世界に類を見ない分散型エネルギーシステムの構築が可能である。",
        },
      ],
    },
    {
      id: "node-11",
      argument: {
        statement: "大規模原発事故の壊滅的リスクが根本的に解消される",
      },
      is_rebuttal: false,
      importance: [
        {
          statement: "国民の生命と財産を守る最も重要な政策効果である。",
        },
        {
          statement: "事故後の復旧コストと社会的損失を完全に回避できる。",
        },
      ],
      uniqueness: [
        {
          statement:
            "火力発電所の爆発やダムの決壊も甚大な被害をもたらすが、放射性物質による広範囲かつ数世代にわたる国土汚染、強制避難、風評被害という複合的で長期的なダメージは、原子力事故にのみ固有の脅威である。",
        },
      ],
    },
    {
      id: "node-12",
      argument: {
        statement: "放射性廃棄物の最終処分問題が解決される",
      },
      is_rebuttal: false,
      importance: [
        {
          statement:
            "将来世代への負担を軽減し、持続可能な社会の実現に貢献する。",
        },
        {
          statement: "処分場選定の社会的対立を根本的に解決できる。",
        },
      ],
      uniqueness: [
        {
          statement:
            "日本特有の地震・火山活動の活発さを考慮すると、安全な最終処分場の選定は極めて困難であり、脱原発によりこの問題を根本的に解決できる。",
        },
      ],
    },
    {
      id: "node-13",
      argument: {
        statement: "国際的な気候変動対策への貢献が強化される",
      },
      is_rebuttal: false,
      importance: [
        {
          statement:
            "パリ協定の目標達成に向けた国際社会の取り組みに大きく貢献する。",
        },
        {
          statement: "日本の環境技術の国際展開を加速させる。",
        },
      ],
      uniqueness: [
        {
          statement:
            "日本の脱原発モデルは、他の先進国や発展途上国にとって参考となる先例となり、世界的な脱炭素化を促進できる。",
        },
      ],
    },
    {
      id: "node-14",
      argument: {
        statement: "エネルギー価格の変動リスクが軽減される",
      },
      is_rebuttal: false,
      importance: [
        {
          statement:
            "化石燃料価格の変動に左右されない安定したエネルギー供給が実現する。",
        },
        {
          statement: "家計と企業のエネルギーコストの予見可能性が向上する。",
        },
      ],
      uniqueness: [
        {
          statement:
            "日本のエネルギー輸入依存度の高さを考慮すると、再生可能エネルギーの自給率向上は特に重要な経済効果を持つ。",
        },
      ],
    },
    {
      id: "node-15",
      argument: {
        statement: "地域コミュニティの自立性が向上する",
      },
      is_rebuttal: false,
      importance: [
        {
          statement:
            "地域主導のエネルギー事業により、地域経済の活性化と雇用創出が実現する。",
        },
        {
          statement:
            "地域の意思決定権が強化され、持続可能な地域社会の構築につながる。",
        },
      ],
      uniqueness: [
        {
          statement:
            "日本の地域特性を活かした、小規模分散型の再生可能エネルギー事業モデルは、世界の地域社会にとって参考となる先進事例である。",
        },
      ],
    },
    {
      id: "node-16",
      argument: {
        statement: "原子力発電の経済性は実際には高くない",
      },
      is_rebuttal: false,
      importance: [
        {
          statement:
            "廃炉費用、放射性廃棄物処分費用、事故リスクの保険費用を考慮すると、原子力発電の真のコストは非常に高い。",
        },
        {
          statement:
            "福島第一原発事故の賠償・除染費用は数十兆円規模に達し、国民負担は莫大である。",
        },
      ],
      uniqueness: [
        {
          statement:
            "日本の原子力発電のコスト計算には、事故リスクの社会的コストが適切に反映されておらず、真の経済性が過大評価されている。",
        },
      ],
    },
    {
      id: "node-17",
      argument: {
        statement: "再生可能エネルギーの不安定性は技術革新で克服可能",
      },
      is_rebuttal: false,
      importance: [
        {
          statement:
            "蓄電池技術の進歩により、再生可能エネルギーの出力変動は大幅に軽減されている。",
        },
        {
          statement:
            "スマートグリッド技術により、需要と供給の最適化が実現されている。",
        },
      ],
      uniqueness: [
        {
          statement:
            "日本の蓄電池技術は世界トップレベルであり、再生可能エネルギーの安定化に必要な技術的基盤が整っている。",
        },
      ],
    },
  ],
  edges: [
    {
      id: "edge-1",
      cause_id: "node-1",
      effect_id: "node-2",
      is_rebuttal: false,
      certainty: [
        {
          statement:
            "脱原発政策の推進により、原子力発電の代替として再生可能エネルギーへの投資が加速される",
        },
        {
          statement:
            "政府の政策転換により、原子力発電の維持・再稼働が現実的な選択肢として検討される",
        },
      ],
    },
    {
      id: "edge-2",
      cause_id: "node-1",
      effect_id: "node-3",
      is_rebuttal: false,
      certainty: [
        {
          statement:
            "脱原発政策の明確な方針により、段階的な原子力発電所の停止が計画される",
        },
        {
          statement:
            "安全性の観点から、既存の原子力発電所の段階的廃止が決定される",
        },
      ],
    },
    {
      id: "edge-3",
      cause_id: "node-1",
      effect_id: "node-4",
      is_rebuttal: false,
      certainty: [
        {
          statement:
            "脱原発政策の実現には再生可能エネルギーへの大規模投資が不可欠",
        },
        {
          statement:
            "政府のグリーン成長戦略により、公的投資の拡大が確実に実行される",
        },
      ],
    },
    {
      id: "edge-4",
      cause_id: "node-1",
      effect_id: "node-5",
      is_rebuttal: false,
      certainty: [
        {
          statement: "脱原発政策の成功には省エネ技術の開発・普及が必須",
        },
        {
          statement:
            "エネルギー安全保障の観点から、省エネ技術の国家戦略化が推進される",
        },
      ],
    },
    {
      id: "edge-5",
      cause_id: "node-2",
      effect_id: "node-16",
      is_rebuttal: false,
      certainty: [
        {
          statement: "原子力発電の真のコストには事故リスクと廃炉費用が含まれる",
        },
        {
          statement:
            "福島事故の教訓により、原子力発電の経済性の過大評価が明らかになった",
        },
      ],
    },
    {
      id: "edge-6",
      cause_id: "node-4",
      effect_id: "node-6",
      is_rebuttal: false,
      certainty: [
        {
          statement:
            "再生可能エネルギーへの投資拡大により、関連産業の成長が促進される",
        },
        {
          statement:
            "政府の産業育成政策により、国内産業の成長が確実に実現される",
        },
      ],
    },
    {
      id: "edge-7",
      cause_id: "node-4",
      effect_id: "node-7",
      is_rebuttal: false,
      certainty: [
        {
          statement:
            "再生可能エネルギーの不安定性を克服するため、スマートグリッド技術の開発が加速される",
        },
        {
          statement:
            "蓄電技術の進歩により、再生可能エネルギーの安定供給が実現される",
        },
      ],
    },
    {
      id: "edge-8",
      cause_id: "node-5",
      effect_id: "node-8",
      is_rebuttal: false,
      certainty: [
        {
          statement:
            "省エネ技術の開発・普及により、電力需要の抑制が確実に実現される",
        },
        {
          statement:
            "IEAの推奨により、エネルギー効率改善が最優先政策として実行される",
        },
      ],
    },
    {
      id: "edge-9",
      cause_id: "node-6",
      effect_id: "node-9",
      is_rebuttal: false,
      certainty: [
        {
          statement:
            "再生可能エネルギー産業の成長により、新たな雇用創出が必然的に発生する",
        },
        {
          statement:
            "IRENAの予測に基づき、グリーン分野での雇用創出が確実に実現される",
        },
      ],
    },
    {
      id: "edge-10",
      cause_id: "node-6",
      effect_id: "node-10",
      is_rebuttal: false,
      certainty: [
        {
          statement:
            "再生可能エネルギー産業の分散型特性により、エネルギー供給システムの分散化が促進される",
        },
        {
          statement:
            "地域特性を活かしたエネルギー事業により、システムの強靭性が向上する",
        },
      ],
    },
    {
      id: "edge-11",
      cause_id: "node-7",
      effect_id: "node-10",
      is_rebuttal: false,
      certainty: [
        {
          statement:
            "スマートグリッド技術により、分散型エネルギーシステムの管理が可能になる",
        },
        {
          statement:
            "高性能蓄電技術により、地域間の電力融通が実現され、システム全体の強靭性が向上する",
        },
      ],
    },
    {
      id: "edge-12",
      cause_id: "node-3",
      effect_id: "node-11",
      is_rebuttal: false,
      certainty: [
        {
          statement:
            "原子力発電所の段階的停止により、大規模原発事故のリスクが確実に軽減される",
        },
        {
          statement:
            "福島事故の教訓により、原発事故の壊滅的リスクの解消が最優先課題として認識される",
        },
      ],
    },
    {
      id: "edge-13",
      cause_id: "node-3",
      effect_id: "node-12",
      is_rebuttal: false,
      certainty: [
        {
          statement:
            "原子力発電所の停止により、新たな放射性廃棄物の発生が停止する",
        },
        {
          statement: "既存の廃棄物処分問題の解決に向けた取り組みが加速される",
        },
      ],
    },
    {
      id: "edge-14",
      cause_id: "node-1",
      effect_id: "node-13",
      is_rebuttal: false,
      certainty: [
        {
          statement:
            "脱原発政策の実現により、日本の気候変動対策への貢献が国際的に評価される",
        },
        {
          statement:
            "パリ協定の目標達成に向けた具体的な行動として、国際社会から高い評価を受ける",
        },
      ],
    },
    {
      id: "edge-15",
      cause_id: "node-10",
      effect_id: "node-14",
      is_rebuttal: false,
      certainty: [
        {
          statement:
            "分散化・強靭化されたエネルギー供給システムにより、価格変動リスクが軽減される",
        },
        {
          statement:
            "地域特性を活かしたエネルギー供給により、化石燃料への依存度が低下する",
        },
      ],
    },
    {
      id: "edge-16",
      cause_id: "node-6",
      effect_id: "node-15",
      is_rebuttal: false,
      certainty: [
        {
          statement:
            "再生可能エネルギー産業の地域分散により、地域コミュニティの自立性が向上する",
        },
        {
          statement:
            "地域主導のエネルギー事業により、地域の意思決定権が強化される",
        },
      ],
    },
    {
      id: "edge-17",
      cause_id: "node-7",
      effect_id: "node-17",
      is_rebuttal: false,
      certainty: [
        {
          statement:
            "スマートグリッドと蓄電技術の進歩により、再生可能エネルギーの不安定性が克服される",
        },
        {
          statement:
            "日本の技術力により、再生可能エネルギーの安定化が確実に実現される",
        },
      ],
    },
  ],
  node_rebuttals: [],
  edge_rebuttals: [],
  counter_argument_rebuttals: [],
  turn_argument_rebuttals: [],
};

export const debateGraphService = {
  async getDebateGraph(): Promise<DebateGraph> {
    // 実際のAPIコールをシミュレート
    return new Promise((resolve) => {
      setTimeout(() => {
        resolve(sampleData);
      }, 500);
    });
  },
};
