-- name: GetLatestElection :one
SELECT id, name FROM elections ORDER BY election_date DESC LIMIT 1;

-- name: GetQuestionsByElectionID :many
SELECT id, title, question_text, description FROM questions WHERE election_id = $1 ORDER BY display_order ASC;

-- name: GetAllParties :many
SELECT id, name FROM parties;

-- name: GetPartyStancesByElectionID :many
SELECT ps.question_id, ps.party_id, ps.stance, p.name as party_name
FROM party_stances ps
JOIN questions q ON ps.question_id = q.id
JOIN parties p ON ps.party_id = p.id
WHERE q.election_id = $1;

-- name: GetElectionByName :one
SELECT id FROM elections WHERE name = $1;

-- name: GetPartyByName :one
SELECT id FROM parties WHERE name = $1;

-- name: GetQuestionByTitleAndElectionID :one
SELECT id FROM questions WHERE title = $1 AND election_id = $2;

-- name: CreateElection :one
INSERT INTO elections (name, election_date) VALUES ($1, $2) RETURNING id;

-- name: CreateParty :one
INSERT INTO parties (name, short_name) VALUES ($1, $2) RETURNING id;

-- name: CreateQuestion :one
INSERT INTO questions (election_id, title, question_text, description, display_order) VALUES ($1, $2, $3, $4, $5) RETURNING id;

-- name: CreatePartyStance :one
INSERT INTO party_stances (party_id, question_id, stance) VALUES ($1, $2, $3) RETURNING id;

-- name: TruncateAllTables :exec
TRUNCATE TABLE party_stances, questions, parties, elections RESTART IDENTITY CASCADE;