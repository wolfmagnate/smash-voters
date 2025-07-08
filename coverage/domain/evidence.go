package domain

// Evidence は、特定の主張や論拠を裏付ける具体的な証拠情報を表します。
type Evidence struct {
	URL   string `json:"url"`   // 証拠となる情報源のURL
	Title string `json:"title"` // 情報源のタイトル（例：「経産省のエネルギー白書」）
}

// Assertion は、特定の性質（重要性、独自性、確実性など）に関する
// 一つのまとまった主張（Statement）と、それを裏付ける証拠（Evidence）のセットです。
type Assertion struct {
	Statement string      `json:"statement"`          // 主張の内容（例：「安定供給により産業競争力が向上する」）
	Evidence  []*Evidence `json:"evidence,omitempty"` // その主張を裏付ける証拠のリスト
}
