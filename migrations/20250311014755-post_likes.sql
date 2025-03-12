
-- +migrate Up
CREATE TABLE post_likes(
    post_id BINARY(16) PRIMARY KEY COMMENT "投稿ID",
    user_id BINARY(16) NOT NULL COMMENT "ユーザーID",
    UNIQUE (post_id, user_id),
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE post_likes;
