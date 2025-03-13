
-- +migrate Up
ALTER TABLE posts ADD COLUMN title VARCHAR(255) NOT NULL COMMENT "投稿タイトル";

-- +migrate Down
