package test

import (
	"os"
	"testing"

	"github.com/wolfmagnate/smash-voters/coverage/domain"
)

func TestRegenerateSampleJson(t *testing.T) {
	dg := domain.NewDebateGraph()

	// Helper function to create assertions from simple strings
	makeTextAssertion := func(text string, evidences ...*domain.Evidence) *domain.Assertion {
		return &domain.Assertion{Statement: text, Evidence: evidences}
	}

	// Node definitions from sample.json
	nodeArgs := []struct {
		Arg        string
		IsRebuttal bool
		Importance []*domain.Assertion
		Uniqueness []*domain.Assertion
		Certainty  []*domain.Assertion // For edges, but we can use it for nodes if needed
	}{
		{Arg: "脱原発政策を推進し再生可能エネルギーへ転換する", IsRebuttal: false},
		{Arg: "原子力発電を基幹電源として維持・再稼働する", IsRebuttal: false},
		{Arg: "原子力発電所を段階的に停止する", IsRebuttal: false},
		{Arg: "再生可能エネルギー導入への公的投資を拡大する", IsRebuttal: false},
		{Arg: "省エネルギー技術の開発・普及を国家戦略として推進する", IsRebuttal: false},
		{Arg: "再生可能エネルギー関連産業が国内で成長する", IsRebuttal: false},
		{Arg: "スマートグリッドと高性能蓄電技術が社会に普及する", IsRebuttal: false},
		{Arg: "電力需要の全体量が効率化により抑制される", IsRebuttal: false, Importance: []*domain.Assertion{makeTextAssertion("国際エネルギー機関(IEA)は、エネルギー効率の改善を『第一の燃料(First Fuel)』と位置づけており、最もクリーンで安価なエネルギー供給源であると評価している。需要の抑制は、新たな発電所建設の必要性を低減させる。", &domain.Evidence{URL: "https://www.iea.org/reports/energy-efficiency-2023", Title: "Energy Efficiency 2023 - IEA"})}},
		{Arg: "グリーン分野で質の高い新たな雇用が創出される", IsRebuttal: false, Certainty: []*domain.Assertion{makeTextAssertion("IRENAの予測では、野心的なエネルギー転換シナリオの下で、2050年までに世界のエネルギー分野の雇用は1億3900万人に達し、そのうち4300万人が再生可能エネルギー分野とされている。この世界的な潮流に乗ることができる。", &domain.Evidence{URL: "https://www.irena.org/publications/2023/Sep/Renewable-Energy-and-Jobs-Annual-Review-2023", Title: "Renewable Energy and Jobs: Annual Review 2023"})}, Uniqueness: []*domain.Assertion{makeTextAssertion("原子力産業の雇用が特定の地域・専門分野に集中するのに対し、再生可能エネルギー関連の雇用は、建設、製造、メンテナンス、ソフトウェア開発など多岐にわたり、全国に分散して創出される。この雇用の裾野の広さと成長ポテンシャルは、縮小・固定化傾向にある原子力産業では得られない独自のメリットである。")}},
		{Arg: "エネルギー供給システムが分散化・強靭化する", IsRebuttal: false},
		{Arg: "大規模原発事故の壊滅的リスクが根本的に解消される", IsRebuttal: false, Uniqueness: []*domain.Assertion{makeTextAssertion("火力発電所の爆発やダムの決壊も甚大な被害をもたらすが、放射性物質による広範囲かつ数世代にわたる国土汚染、強制避難、風評被害という複合的で長期的なダメージは、原子力事故にのみ固有の脅威である。このリスクは、他のいかなるエネルギー源でも代替できない。")}},
		{Arg: "高レベル放射性廃棄物の新規発生が停止し、将来世代への負担増を食い止める", IsRebuttal: false, Uniqueness: []*domain.Assertion{makeTextAssertion("太陽光パネルや風力タービンの廃棄も課題だが、これらはリサイクル技術が確立されつつあり、有害性も管理可能である。一方で、高レベル放射性廃棄物は数万年にわたり致死的な放射線を放ち続け、現在の科学技術では無害化できない。この質的・時間的スケールの違いは決定的であり、原発にのみ固有の解決不能な問題である。")}},
		{Arg: "大規模災害時におけるエネルギー供給のレジリエンスが向上する", IsRebuttal: false, Importance: []*domain.Assertion{makeTextAssertion("地震や台風による大規模停電時においても、地域コミュニティ内の太陽光発電や蓄電池が機能すれば、避難所、医療機関、通信基地局などの生命維持に不可欠なインフラを維持できる。")}, Uniqueness: []*domain.Assertion{makeTextAssertion("原子力のような中央集権型の大規模電源は、送電網のハブや発電所自体が被災すると連鎖的に広範囲のブラックアウトを引き起こす。対照的に、地域に分散した再エネと蓄電池は、個々の施設が被災してもシステム全体が停止するリスクが低く、自律的な電力供給が可能。この災害耐性は、脱原発・分散型システムへの移行によって初めて本格的に実現される。")}},
		{Arg: "エネルギー自給率が向上しエネルギー安全保障が強化される", IsRebuttal: false},
		{Arg: "運転開始後40年を超える老朽原発の運転が延長される", IsRebuttal: false},
		{Arg: "使用済み核燃料が継続的に発生し国内に蓄積され続ける", IsRebuttal: false},
		{Arg: "核燃料サイクル事業への巨額の公的資金投入が継続される", IsRebuttal: false, Certainty: []*domain.Assertion{makeTextAssertion("青森県六ヶ所村の再処理工場は、建設開始から約30年が経過し、総事業費は当初計画の4倍近い約16兆円に膨れ上がっているが、本格操業には至っておらず、今後も多額の維持費・追加投資が必要となる。", &domain.Evidence{URL: "https://www.japantimes.co.jp/news/2023/12/26/japan/science-health/rokkasho-plant-delay/", Title: "Japan's Rokkasho nuclear fuel reprocessing plant faces further delay"})}},
		{Arg: "経年劣化による重大事故・故障のリスクが増加する", IsRebuttal: false, Certainty: []*domain.Assertion{makeTextAssertion("米国の原子力規制委員会(NRC)の研究では、運転年数が長くなるにつれて、圧力容器の脆化やケーブルの絶縁劣化など、交換が困難な重要機器の故障確率が統計的に有意に上昇することが示されている。", &domain.Evidence{URL: "https://www.nrc.gov/reading-rm/doc-collections/nuregs/staff/sr1801/", Title: "NUREG-1801, Rev. 2: Generic Aging Lessons Learned (GALL) Report"})}},
		{Arg: "将来の廃炉コストがさらに高騰し、国民負担が増大する", IsRebuttal: false, Certainty: []*domain.Assertion{makeTextAssertion("会計検査院の2023年の報告によると、国内の商業用原発24基の廃炉費用は総額で約8兆円に達すると試算されており、これは電力会社の当初の見積もりを大幅に上回っている。運転を延長すれば、放射化する範囲が広がり、廃炉はさらに複雑かつ高コストになる。")}},
		{Arg: "「核のゴミ」問題が解決不能なまま次世代へ先送りされる", IsRebuttal: false, Importance: []*domain.Assertion{makeTextAssertion("高レベル放射性廃棄物の最終処分地の選定は、国内で一向に進んでいない。行き場のない核のゴミを増やし続けることは、将来世代から安全な国土と問題解決の選択肢を奪う、深刻な倫理的問題である。")}, Uniqueness: []*domain.Assertion{makeTextAssertion("脱原発政策は、この無限に続く負の遺産の連鎖を断ち切る唯一の方法である。原発維持は問題を悪化させるだけだが、脱原発は『これ以上増やさない』という最低限の倫理的責任を果たすための、具体的かつ唯一の選択肢となる。")}},
		{Arg: "より有望な次世代エネルギー技術への投資機会を逸失する", IsRebuttal: false, Importance: []*domain.Assertion{makeTextAssertion("原子力関連に固定化される巨額の資金と優秀な人材を、蓄電池、水素エネルギー、次世代送電網、CCUS（二酸化炭素回収・利用・貯留）といった未来の成長分野に振り向けることができず、国際競争に遅れをとる。")}, Uniqueness: []*domain.Assertion{makeTextAssertion("原子力関連事業（維持費、安全対策、核燃料サイクル）に投じられる年間数千億円〜数兆円規模の資金は、機会費用そのものである。脱原発を選択することで初めて、この莫大なリソース（資金・人材）を、世界が競い合う未来の成長産業へ戦略的に再配分することが可能になるという、明確なトレードオフの関係にある。")}},
		{Arg: "再生可能エネルギー設備も将来的に大量の廃棄物を生み、環境負荷となる", IsRebuttal: true, Importance: []*domain.Assertion{makeTextAssertion("太陽光パネルの寿命は約20〜30年であり、経済産業省の試算では2040年頃に廃棄のピークを迎える。パネルには鉛やセレンなどの有害物質が含まれる場合があり、不適切な処理は土壌汚染を引き起こすため、新たな社会問題となる。", &domain.Evidence{URL: "https://www.meti.go.jp/shingikai/energy_environment/solar_panel_recycle/pdf/001_03_00.pdf", Title: "太陽光発電設備のリサイクル等の推進に向けたガイドライン（第一版）"})}},
		{Arg: "再生可能エネルギーは出力が天候に左右され、安定供給には課題がある", IsRebuttal: true, Importance: []*domain.Assertion{makeTextAssertion("太陽光は夜間や曇天時には発電できず、風力も風がなければ発電できない。電力需要のピーク時に安定して供給するには、大規模な蓄電池か、結局は調整力として火力発電によるバックアップが不可欠となり、コスト増と化石燃料依存を招く。")}},
		{Arg: "急進的なエネルギー転換は既存産業の安定した雇用を破壊する", IsRebuttal: true, Importance: []*domain.Assertion{makeTextAssertion("原子力や火力発電所の閉鎖は、そこで働く高度な専門性を持つ技術者や、関連するサプライチェーン企業の従業員の職を直接的に奪う。再エne分野で創出される雇用が、これらの失業者を全て吸収できる保証はなく、スキルのミスマッチによる構造的な失業は地域経済に深刻な打撃を与える。")}},
		{Arg: "原子力への継続的投資が次世代革新炉（SMR等）の研究開発を促進する", IsRebuttal: true, Importance: []*domain.Assertion{makeTextAssertion("SMR（小型モジュール炉）や高温ガス炉などの次世代炉は、従来型原発よりも安全性が高く、多様な用途に利用できる可能性がある。これはカーボンニュートラル達成のための重要な選択肢であり、日本の技術的優位性を確保する上でも不可欠である。")}},
		{Arg: "省エネ努力による需要削減効果は限定的であり、経済成長が需要を上回る", IsRebuttal: true, Certainty: []*domain.Assertion{makeTextAssertion("エネルギー効率が向上すると、その分機器の使用頻度が増えたり、新たな電力需要（例：データセンター、EV充電）が生まれたりするため、社会全体のエネルギー消費量が必ずしも減るとは限らない。これを『ジェボンズのパラドックス』またはリバウンド効果と呼び、需要抑制の有効性を減じる。")}},
		{Arg: "分散型電源も広域災害に対して脆弱であり、そのレジリエンスは過大評価されている", IsRebuttal: true, Certainty: []*domain.Assertion{makeTextAssertion("2018年の北海道胆振東部地震では、大規模な停電（ブラックアウト）が発生し、多くの太陽光発電設備も系統から解列され機能しなかった。広範囲にわたる送配電網の復旧が不可欠であり、分散型電源だけでは都市機能の維持は困難である。")}},
		{Arg: "原発停止は、火力発電への過度な依存や系統の不安定化を招き、別種の大規模停電リスクを増大させる", IsRebuttal: true, Certainty: []*domain.Assertion{makeTextAssertion("ベースロード電源である原子力を失うと、天候に左右される再エネの出力変動を埋めるため、火力発電の急な出力調整が頻発する。これは発電設備の負担を増やし、故障リスクを高める。また、需給バランスが崩れやすくなり、周波数変動による大規模停電の引き金となりうる。")}},
		{Arg: "電力システムの過度な分散化は管理コストを増大させ、電気料金を高騰させる", IsRebuttal: true, Importance: []*domain.Assertion{makeTextAssertion("無数の小規模電源を束ね、常に電力の品質（周波数・電圧）を維持するための制御システムや送配電網の増強には莫大なコストがかかる。このコストは託送料金などを通じて消費者に転嫁され、国民生活や企業の国際競争力を圧迫する要因となる。")}},
	}

	for _, n := range nodeArgs {
		node := domain.NewDebateGraphNode(makeTextAssertion(n.Arg), n.IsRebuttal)
		node.Importance = append(node.Importance, n.Importance...)
		node.Uniqueness = append(node.Uniqueness, n.Uniqueness...)
		// sample.jsonではCertaintyがNodeにないので、Certaintyはここでは追加しない
		err := dg.AddNode(node)
		if err != nil {
			// 既存ノードのエラーは無視して良い
		}
	}

	// Edge definitions
	edgeDefs := []struct {
		Cause      string
		Effect     string
		IsRebuttal bool
		Uniqueness []*domain.Assertion
	}{
		{"脱原発政策を推進し再生可能エネルギーへ転換する", "原子力発電所を段階的に停止する", false, nil},
		{"脱原発政策を推進し再生可能エネルギーへ転換する", "再生可能エネルギー導入への公的投資を拡大する", false, nil},
		{"脱原発政策を推進し再生可能エネルギーへ転換する", "省エネルギー技術の開発・普及を国家戦略として推進する", false, nil},
		{"再生可能エネルギー導入への公的投資を拡大する", "再生可能エネルギー関連産業が国内で成長する", false, nil},
		{"再生可能エネルギー導入への公的投資を拡大する", "スマートグリッドと高性能蓄電技術が社会に普及する", false, nil},
		{"省エネルギー技術の開発・普及を国家戦略として推進する", "電力需要の全体量が効率化により抑制される", false, nil},
		{"再生可能エネルギー関連産業が国内で成長する", "グリーン分野で質の高い新たな雇用が創出される", false, nil},
		{"再生可能エネルギー関連産業が国内で成長する", "エネルギー供給システムが分散化・強靭化する", false, nil},
		{"スマートグリッドと高性能蓄電技術が社会に普及する", "エネルギー供給システムが分散化・強靭化する", false, nil},
		{"原子力発電所を段階的に停止する", "大規模原発事故の壊滅的リスクが根本的に解消される", false, nil},
		{"原子力発電所を段階的に停止する", "高レベル放射性廃棄物の新規発生が停止し、将来世代への負担増を食い止める", false, nil},
		{"エネルギー供給システムが分散化・強靭化する", "大規模災害時におけるエネルギー供給のレジリエンスが向上する", false, nil},
		{"再生可能エネルギー関連産業が国内で成長する", "エネルギー自給率が向上しエネルギー安全保障が強化される", false, []*domain.Assertion{makeTextAssertion("原子力の燃料ウランも、化石燃料と同様に100%輸入に依存しており、価格変動や供給国の地政学リスクから逃れられない。国内の自然資源（太陽、風、地熱）を最大限活用する再エネと、需要自体を減らす省エネの組み合わせだけが、外部要因に左右されない真のエネルギー安全保障を構築する道である。")}},
		{"電力需要の全体量が効率化により抑制される", "エネルギー自給率が向上しエネルギー安全保障が強化される", false, nil},
		{"原子力発電を基幹電源として維持・再稼働する", "運転開始後40年を超える老朽原発の運転が延長される", false, nil},
		{"原子力発電を基幹電源として維持・再稼働する", "使用済み核燃料が継続的に発生し国内に蓄積され続ける", false, nil},
		{"原子力発電を基幹電源として維持・再稼働する", "核燃料サイクル事業への巨額の公的資金投入が継続される", false, nil},
		{"運転開始後40年を超える老朽原発の運転が延長される", "経年劣化による重大事故・故障のリスクが増加する", false, nil},
		{"運転開始後40年を超える老朽原発の運転が延長される", "将来の廃炉コストがさらに高騰し、国民負担が増大する", false, nil},
		{"経年劣化による重大事故・故障のリスクが増加する", "大規模原発事故の壊滅的リスクが根本的に解消される", false, nil},
		{"使用済み核燃料が継続的に発生し国内に蓄積され続ける", "「核のゴミ」問題が解決不能なまま次世代へ先送りされる", false, nil},
		{"核燃料サイクル事業への巨額の公的資金投入が継続される", "「核のゴミ」問題が解決不能なまま次世代へ先送りされる", false, nil},
		{"核燃料サイクル事業への巨額の公的資金投入が継続される", "より有望な次世代エネルギー技術への投資機会を逸失する", false, nil},
		{"核燃料サイクル事業への巨額の公的資金投入が継続される", "原子力への継続的投資が次世代革新炉（SMR等）の研究開発を促進する", true, nil},
		{"エネルギー供給システムが分散化・強靭化する", "電力システムの過度な分散化は管理コストを増大させ、電気料金を高騰させる", true, nil},
	}

	for _, e := range edgeDefs {
		causeNode, _ := dg.GetNode(e.Cause)
		effectNode, _ := dg.GetNode(e.Effect)
		edge := domain.NewDebateGraphEdge(causeNode, effectNode, e.IsRebuttal)
		edge.Uniqueness = append(edge.Uniqueness, e.Uniqueness...)
		err := dg.AddEdge(edge)
		if err != nil {
			// ignore error
		}
	}

	// Rebuttal definitions
	nodeRebuttals := []struct {
		Target   string
		Type     string
		Rebuttal string
	}{
		{"高レベル放射性廃棄物の新規発生が停止し、将来世代への負担増を食い止める", "uniqueness", "再生可能エネルギー設備も将来的に大量の廃棄物を生み、環境負荷となる"},
		{"電力需要の全体量が効率化により抑制される", "importance", "省エネ努力による需要削減効果は限定的であり、経済成長が需要を上回る"},
		{"大規模災害時におけるエネルギー供給のレジリエンスが向上する", "importance", "分散型電源も広域災害に対して脆弱であり、そのレジリエンスは過大評価されている"},
	}
	for _, r := range nodeRebuttals {
		rebuttal, err := domain.NewDebateGraphNodeRebuttal(dg, r.Target, r.Type, r.Rebuttal)
		if err == nil {
			dg.NodeRebuttals = append(dg.NodeRebuttals, rebuttal)
		}
	}

	edgeRebuttals := []struct {
		Cause    string
		Effect   string
		Type     string
		Rebuttal string
	}{
		{"再生可能エネルギー関連産業が国内で成長する", "エネルギー自給率が向上しエネルギー安全保障が強化される", "certainty", "再生可能エネルギーは出力が天候に左右され、安定供給には課題がある"},
		{"原子力発電所を段階的に停止する", "大規模原発事故の壊滅的リスクが根本的に解消される", "uniqueness", "原発停止は、火力発電への過度な依存や系統の不安定化を招き、別種の大規模停電リスクを増大させる"},
	}
	for _, r := range edgeRebuttals {
		rebuttal, err := domain.NewDebateGraphEdgeRebuttal(dg, r.Cause, r.Effect, r.Type, r.Rebuttal)
		if err == nil {
			dg.EdgeRebuttals = append(dg.EdgeRebuttals, rebuttal)
		}
	}

	counterArgRebuttals := []struct {
		Target   string
		Rebuttal string
	}{
		{"グリーン分野で質の高い新たな雇用が創出される", "急進的なエネルギー転換は既存産業の安定した雇用を破壊する"},
	}
	for _, r := range counterArgRebuttals {
		rebuttal, err := domain.NewCounterArgumentRebuttal(dg, r.Target, r.Rebuttal)
		if err == nil {
			dg.CounterArgumentRebuttals = append(dg.CounterArgumentRebuttals, rebuttal)
		}
	}

	turnArgRebuttals := []string{
		"原子力への継続的投資が次世代革新炉（SMR等）の研究開発を促進する",
		"電力システムの過度な分散化は管理コストを増大させ、電気料金を高騰させる",
	}
	for _, r := range turnArgRebuttals {
		rebuttal, err := domain.NewTurnArgumentRebuttal(dg, r)
		if err == nil {
			dg.TurnArgumentRebuttals = append(dg.TurnArgumentRebuttals, rebuttal)
		}
	}

	// Generate JSON
	jsonStr, err := dg.ToIDJson()
	if err != nil {
		t.Fatalf("Failed to generate JSON: %v", err)
	}

	err = os.WriteFile("../new_sample.json", []byte(jsonStr), 0644)
	if err != nil {
		t.Fatalf("failed to write file: %v", err)
	}
}
