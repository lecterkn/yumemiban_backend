
-- +migrate Up
CREATE TABLE posts(
    id BINARY(16) PRIMARY KEY COMMENT "投稿ID",
    user_id BINARY(16) NOT NULL COMMENT "投稿者ユーザーID",
    nickname VARCHAR(255) NOT NULL COMMENT "ニックネーム",
    content VARCHAR(511) NOT NULL COMMENT "投稿内容",
    novel VARCHAR(511) NOT NULL COMMENT "小説内容",
    created_at DATETIME NOT NULL COMMENT "投稿日時",
    updated_at DATETIME NOT NULL COMMENT "更新日時",
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE posts;
