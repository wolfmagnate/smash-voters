-- 選挙情報を管理するテーブル
-- 各マッチング企画はこのテーブルを親とする
CREATE TABLE elections (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    election_date DATE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    -- 例: '第27回参議院議員通常選挙'
    UNIQUE(name)
);

-- 政党の基本情報を管理するテーブル
-- このデータは選挙をまたいで共通で利用される
CREATE TABLE parties (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    short_name VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    -- 例: '自由民主党', '自民党'
    UNIQUE(name)
);

-- 設問情報を管理するテーブル
-- 設問は選挙(elections)に紐づく
CREATE TABLE questions (
    id SERIAL PRIMARY KEY,
    election_id INTEGER NOT NULL,
    title VARCHAR(100) NOT NULL,
    question_text TEXT NOT NULL,
    description TEXT,
    display_order INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    -- 外部キー制約: electionsテーブルのIDを参照
    FOREIGN KEY (election_id) REFERENCES elections(id) ON DELETE CASCADE,
    -- 同じ選挙内での表示順はユニークであるべき
    UNIQUE(election_id, display_order)
);

-- 各政党の各設問に対するスタンス（回答）を管理するテーブル
-- どの政党が、どの選挙の、どの質問にどう答えたかを記録
CREATE TABLE party_stances (
    id SERIAL PRIMARY KEY,
    party_id INTEGER NOT NULL,
    question_id INTEGER NOT NULL,
    -- -2: 反対, -1: やや反対, 0: 中立, 1: やや賛成, 2: 賛成
    stance INTEGER NOT NULL CHECK (stance BETWEEN -2 AND 2),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    -- 外部キー制約
    FOREIGN KEY (party_id) REFERENCES parties(id) ON DELETE CASCADE,
    FOREIGN KEY (question_id) REFERENCES questions(id) ON DELETE CASCADE,
    -- 同じ政党が同じ質問に複数回答できないようにする
    UNIQUE(party_id, question_id)
);

-- Indexを作成して検索パフォーマンスを向上
CREATE INDEX idx_questions_election_id ON questions(election_id);
CREATE INDEX idx_party_stances_question_id ON party_stances(question_id);
