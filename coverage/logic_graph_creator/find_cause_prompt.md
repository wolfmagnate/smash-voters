# 議論の内容分析

## 役割
あなたは文章分析アシスタントです。

## タスク
与えられた分析対象の文章と、それに関する議論構造分析結果のJSONに基づき、指定された文章中の主張の直接の原因になっている主張をすべて列挙してください。

## 入力
1. 分析対象の文章：メリット・デメリットを分析する元の文章
2. 議論構造分析結果：次の構造を持ったJSON

```json
{
  "is_argument": true,
  "status_quo": "現状維持策の説明",
  "affirmative_plan": "提案行動策の説明",
  "position": "status_quo" // または "affirmative_plan"
}
```

positionプロパティは、分析対象の文章がどちらに賛成しているかを表します。

3. 分析対象の主張：分析対象の文章に記述された主張

## 出力形式
以下の形式のJSONで出力してください。
```go
// FoundCausesは分析対象の主張の直接の原因として本文中で述べられている内容です
type FoundCauses struct {
	Causes []string `json:"causes"`
}
```

## 分析のポイント
### 書かれていないことを類推しないこと
原因は、分析対象の文章中に書かれた根拠のある内容を抽出してください。あなたが推論した結果の原因を含めないでください。

例えば分析対象の文章が「死刑は「国家による殺人」であり、文明社会の価値観にそぐわないと考えます。いかなる理由があろうとも、国家が人の生命を奪う権利を持つべきではありません。これは応報感情を満たす以上の意味を持たず、社会全体の倫理観を低下させる恐れがあります。」だとします。

この場合、死刑制度がある世界(Status Quo)の話をしています。したがって、「社会全体の倫理観を低下させる」の原因として「死刑に応報感情を満たす以上の意味がない」とか「国家に殺人を許可することは文明社会の価値観にそぐわない」とかを挙げられます。しかし、「社会全体の倫理観が向上する」原因として「文明社会の価値観に合致する」を挙げてはいけません。なぜなら、この文章では死刑制度がない社会の分析はしていないからです。勝手にAffirmative Plan側の世界にStatus Quoの逆側の主張がなされていると想定してはいけません。

### 論理的に正しい因果関係を抽出すること
因果関係とは、「XをしたからYが発生した」という意味です。したがって、Causesに列挙する原因は、「原因」をしたから「分析対象の主張」が発生した、と論理的に言える必要があります。

例えば「生徒一人ひとりの潜在能力を最大限に引き出し、より専門的で質の高い教育を提供できるのです。」という文章において「専門的で質の高い教育が提供できる」の原因は「生徒一人ひとりの潜在能力を最大限に引き出せる」ことではありません。能力を引き出すことが質の高い教育を引き起こすというのは論理的におかしく、むしろ質の高い教育の結果として生徒の能力が改善されるからです。
もしくは、「個別の生徒の能力を引き出す教育」と「専門的で質の高い教育」が等しい並立の関係と言えるかもしれません。

いずれにしろ、絶対に論理的に原因になる内容だけを出力してください。この点が最重要です。

### 因果関係にある主張のみを抽出すること
直接の原因だけを抽出してください。例えば、次のようなものが直接の原因です。

分析対象の主張：「製品の需要が高まった。」
文中の記述：「インフルエンサーによる紹介がきっかけで、製品の需要が高まった。」
抽出される原因：「インフルエンサーによる紹介があった。」

#### 具体例は原因ではない
逆に、主張を強めるための記述でも直接の原因でないものがあります。例えば、具体例は原因ではありません。

分析対象の主張：「運動は健康に良い。」
文中の記述：「例えば、定期的なジョギングは心肺機能を高める。」
この場合、「定期的なジョギングは心肺機能を高める」は「運動は健康に良い」という主張を具体的に示す例であり、直接的な原因ではありません。

#### 因果関係を強める主張は原因ではない
他にも、原因と混同しやすいものとして、原因と結果を補強するための主張が挙げられます。

分析対象の主張：「新製品Xは市場で大きなシェアを獲得する」
文中の記述「A社は新製品Xを発売した。この製品は、従来の製品Yと比較して処理速度が2倍であり、消費電力も半減している。日常の業務の処理速度が遅いと非常に苛立たしいため、新製品Xは市場で大きなシェアを獲得するだろう」
この場合、「処理速度が遅いと非常に苛立たしい」ことはXのシェアが高いことの直接の原因ではない。直接の原因は「従来製品に比べて処理速度が2倍であり、消費電力も半減している」ことであり、クレームは「性能改善」が「シェアの獲得」を引き起こす可能性が高いことを説明するための主張である。
ユーザーが怒るだけで製品のシェアが大きくなるはずがないため、論理的に考えれば分かりますが、接続詞「ため」で繋がっているため、原因のように見えてしまいます。
しかし、このような原因が結果を引き起こす確率が高いことを示すための主張は原因ではありません。


### メリット・デメリットの強調は原因ではない
与えられた文章にはメリット・デメリット自体ではなく、そのメリット・デメリットの重要性を強調するための説明が含まれるかもしれません。

- 利点や欠点が多くの主体に影響することの説明
  例：情報は瞬く間にSNS上で広がるため一度でも噂が発生すると誤情報によって誤った行動を取る人が多いです
  なぜデメリットではないか：「噂」が引き起こす「誤った行動をとる」がデメリットです。その影響が多いことを示すの証拠となる議論がSNSです。「SNSでの情報が広がる速度が速い」という事実が直接誤った行動を導くわけではありませんよね。
- 利点や欠点が大きなものであることの説明
  例：希望の大学に進学できない苦しみは非常に大きなもので、その後の人生で受験に関して軽いPTSDを持つ人もいます
  なぜデメリットではないか：希望の大学に進学できないことがPTSDを引き起こすと考えると、PTSDは苦しみの強さを強調するための例であり、PTSD自体も苦しみなので、「苦しいから苦しい」という無意味な因果関係になってしまう
- 利点や欠点が長期的に継続して続くことの説明
  例：途上国子供の生活を改善して学習させることで、より給与の高い職に就くことができ、その子孫まで含めて生活改善が予想できる
  なぜメリットではないか：途上国の子供の生活改善自体がメリットであり、子孫の生活改善は子供の生活改善が長期的に続くことの説明だから。メリットが続くことの説明であり、新しいメリットではない

これらのインパクトの強調は議論の重要な要素ですが、直接的なメリット・デメリットそのものや議論の骨組みとなる因果の要素ではありません。非常に判断が難しい点であるので、どれが主要かつ最終的なメリット・デメリットかは慎重に判断してください。

### Status Quo・Affirmative Plan特有であることの説明は原因ではない
Affirmative Planの新規性やStatus Quoの固有の事情であることの説明は議論において重要です。なぜなら、二種類の世界の差分が説得力を生み出すからです。片方のプランのときにのみ発生する主張であることを示すことは重要ですが、これは因果関係とは直接関係ありません。

例：死刑制度の導入時と誤審による被害は、それ以外の刑罰と違って取り返しがつきません。金銭的な補償をしようにも死んでしまえば意味がないからです

「死んだら取り返しがつかない」ことが直接誤審による被害を引き起こすわけではありません。これは死刑制度というStatus Quo特有の事情であることを示すだけです。

### 直接的な因果関係のみを抽出すること
主張Aが主張Bを直接引き起こす場合にのみ、主張Aを原因として抽出してください。
ある主張が別の主張を介して間接的に影響を与えている場合（例：A → C → B）、その中間にある主張CがBの直接の原因となります。

例：
文章：「昨夜、台風が接近し、強風が吹いた。その結果、多くの電柱が倒れた。電柱が倒れたため、広範囲で停電が発生した。」
分析対象の主張：「広範囲で停電が発生した。」
抽出される直接の原因：「電柱が倒れた。」
「昨夜、台風が接近し、強風が吹いた」は間接的な原因であり、このタスクでは抽出しません。

### 複数の因果関係が存在する可能性がある

ある主張に対して、本文中で複数の直接的な原因が述べられている場合は、そのすべてを列挙してください。

例：
分析対象の主張：「そのレストランは閉店した。」
文中の記述：「そのレストランが閉店したのは、近隣に強力な競合店が出現したこと、そして、食材の価格が高騰し続けたことの双方が影響している。」
抽出される原因：「近隣に強力な競合店が出現した」、「食材の価格が高騰し続けた」

### 原因が記述されていない可能性がある
分析対象の文章中に、その主張の直接的な原因が明示的に記述されていない場合もあります。これは、その原因が自明の理（一般常識）であると筆者が判断した場合や、主張自体が単に観察された事実を述べているに過ぎない場合などです。

このような場合は、出力JSONのcausesフィールドを空の配列（[]）としてください。

例：自明の理や議論の前提として扱われる場合
    分析対象の主張「死刑制度が存続する」
    文中の記述「死刑制度を存続させれば、犯罪を犯すときに死刑の危険を感じるため、凶悪犯罪を減らせる」
    これは文章の前提となることであり、なぜ死刑制度が存続しているのかは記述されない。

例：単に文章中に記述がない場合
    分析対象の主張「犯罪の抑止には建設的なアプローチが必要」
    文中の記述「死刑制度が凶悪犯罪の抑止力になるという明確な証拠はありません。犯罪の抑止には、社会環境の改善や教育、再犯防止プログラムの充実など、より建設的なアプローチが必要です。」
    「死刑が犯罪を抑止する明確な証拠がないからより建設的にするべきだ」という文章も考えられます。しかし、「抑止力になる証拠がない」ことが直接「建設的なアプローチの必要性」を引き起こすわけではありません。

絶対に出力する原因は分析対象の主張と因果関係を持つようにしてください。（原因）により（分析対象の主張）が直接的に発生すると言えない場合は出力しないでください。

### 議論の構造を意識する
分析対象の文章は何らかの対立構造を対象としています。Status Quoという現状維持の選択と、Affirmative Planという積極的な改善策を比較して、どちらかを行うべきであるという内容の文章です。そのため、2つの世界（Status QuoとAffirmative Plan）それぞれの分析がなされています。

したがって、Status Quoで現状維持を行うこと、Affirmative Planで現状変更することは議論の前提です。したがって、原因はありません。

例えば「現在、日本のエネルギー供給は、東日本大震災以降の原発停止により、化石燃料への依存度が極めて高まっています。原発の再稼働によって複数の電力供給方法を確保し、安定的なベースロード電源が得られます」という文章では、Status Quoで「原発停止」が「化石燃料への依存度の高まり」を引き起こし、Affirmative Planでは「原発再稼働」が「複数の電力供給方法の確保」と「安定したベースロード電源」を引き起こしています。この2種類の世界において、「原発停止」と「原発再稼働」に理由はありません。

ここで重要なのは「原発を再稼働する」と「原発を再稼働するべき」は全く違うことです。この点については、次節の「因果関係と目的手段関係を区別する」を参考にしてください。

分析対象の文章は説得を行うために「Xをするべきだ」という結論を作り出そうとします。この結論を導くためには、「Xをした世界」と「Xをしない世界」でそれぞれXをすること・しないことによって何らかの因果関係が発生し、メリット・デメリットが引き起こされることを分析します。そして、2つの世界の差分を比べることで、よりよい方がXなのだからXをするべきだと説得できるのです。

例えば、原発再稼働の例では「安定供給できる世界」と「化石燃料依存で不安定な世界」を比較して、「再稼働するべきだ」という結論を導き出しています。Affirmative Planでは「再稼働する」は原因のない分析の前提であり、Status Quoでは「再稼働しない」が前提です。

「再稼働する」は因果関係の起点となる事実ですが、「再稼働するべき」という主張自体は単なる文章全体の結論であって何かの結果を導き出すわけではありません。

### 因果関係と目的手段関係を区別する
「なぜ」とか「～だから」「～のため」というような文字だけを見ていると、論理構造の因果を誤って逆に捉えてしまうことがあります。

例えば「駅前の美観を向上させるため、再開発計画を推進する」という文章では、論理構造グラフ上で「駅前の美観を向上させる」から「再開発を行う」に対してエッジを引きたくなる。なぜなら、「～するため」という接続詞で繋がれているからである。しかし、これは誤りである。なぜなら、どれだけ美しくしたからと言って駅前は再開発されないからである。この文章において、再開発は手段であり、美観向上が目的である。
逆に、「駅前を再開発したため、美観が向上した」の場合は論理構造グラフに適合する。なぜなら、再開発で区画を整理したり、公園を作ったり、古い建物を新しくしたりすることによって美観向上を引き起こせるからである。この場合、再開発が原因であり、美観向上は結果である。

ここで定義されている論理構造グラフにおいて、エッジは原因から結果に対して引かれるため、全く同じ理由で手段から目的に対して引かれる。手段の結果目的を達成するということは、手段が現実世界で何らかの作用を行って目的という結果が生まれるからである。

したがって、次のような論理構造は誤りです。

```json
{
  "argument": "駅周辺の再開発計画を推進すべきである",
  "causes": [
    "歩行者の安全性が懸念される",
    "駅前の美観を向上させる必要がある",
    "地域の商業活動が不調である"
  ]
}
```

すべての`causes`が「なぜ再開発をするべきなのか」という目的に対応しているからである。安易に「なぜ」とか「～だから」といった言葉だけで判断すると、目的手段の関係と、原因結果の関係を見誤るため、十分に議論の内容を理解したうえで判断してください。


### 究極的な原因がStatus QuoまたはAfter Planになるようにする

すでに説明したように、説得力はStatus Quoの世界とAffirmative Planの世界の違いから生まれます。したがって、一見原因に見える主張でも、それがStatus QuoまたはAffirmative Planの仮定から導き出せそうにないならば、原因ではありません。

例えば「原発を再稼働する」ことによって「安定した電力供給が得られる」が引き起こされるという議論において、「原発は天候に左右されない」」という主張は、すでに説明した"因果関係を強める主張"であって、原因ではありません。なぜなら、「原発の再稼働」というAffirmative Planにおいて、「原発を再稼働する」ことが「原発による電力供給が天候に左右されない」を引き起こすというのは論理的に不自然だからです。これは単なる原発の性質です。

したがって、「原発は天候に左右されない」から「安定した電力共有が得られる」を引き起こすという因果関係を出力してはなりません。これは非常に重要です。

あなたが出力する原因は、「Status QuoまたはAffirmative Plan」だから「あなたが出力する原因」が引き起こされ、「あなたが出力する原因」だから「分析対象の主張」が引き起こされる、という両方で自然な論理的な因果関係を持つ必要があります。

### 議論の賛成反対に関係なく因果関係を抽出する
分析対象の文章が、抽出した因果関係の流れに対して反論している場合があります。例えば、

「もちろん、制服廃止によって私服の購入費用が新たにかかるという意見や、生徒間の経済格差が服装に表れやすくなるという懸念も存在します。しかし、これらの課題は、学校や家庭、地域社会が連携し、適切な指導や支援を行うことで克服可能であると考えます。」

という文章では、「制服廃止」→「私費の購入費用発生」と「制服廃止」→「生徒間の経済格差の顕在化」という因果関係があります。しかし、「学校や家庭、地域社会による支援」によってこの因果関係を弱められると主張しています。しかし、反論が存在することに関わらず、因果関係自体は存在するので、「私費の購入費用発生」の原因が求められたら「制服廃止」を出力してください。反論は無視して、Status QuoおよびAffirmative Planにおけるメリット・デメリットとの因果関係のみに注力してください。

## 具体例
分析対象文章:
「現在の私たちの町の公園は設備が古く、利用者も減少傾向にあります。新しい遊具を導入し、カフェスペースを併設するべきです。小さい子供や若者はおしゃれで楽しい場所を好むため、子供連れの家族や若者が集まる魅力的な場所に生まれ変わるでしょう。初期費用はかかりますが、地域活性化と住民の満足度向上に繋がる投資です。」

議論構造分析結果:
```json
{
  "is_argument": true,
  "status_quo": "公園の設備をそのままにしておく",
  "affirmative_plan": "新しい遊具を導入し、カフェスペースを併設する",
  "position": "affirmative_plan"
}
```

分析対象の主張「公園に子供連れの家族や若者が集まる」

```json
{
  "causes": [
    "公園に新しい遊具を導入し、カフェスペースを併設する"
  ]
}
```

ここで注意するべきなのは、「小さい子供や若者はおしゃれで楽しい場所を好む」は直接的原因ではなく、因果関係を補強するための主張であることです。子供や若者がおしゃれで楽しい場所を好きであるという感性の特徴だけで、他に特定の措置を実行せずとも地元の公園に人が来るというのはおかしいからです。あくまで公園に遊具が増え、カフェが増えるから若者が増えるのです。
このように、因果関係を補強する主張と原因の誤りをしないよう、十分に注意してください。

## 分析対象の文章

{{.Document}}

## 議論構造分析結果

{{.BasicArgumentStructure}}

## 分析対象の主張

{{.TargetArgument}}
