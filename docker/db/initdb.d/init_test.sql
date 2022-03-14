CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    -- 後でランダムな数値に設定する
    username VARCHAR(20) NOT NULL,
    password VARCHAR(30) NOT NULL
);
INSERT INTO users(username, password)
VALUES('testuser1', 'test1');
INSERT INTO users(username, password)
VALUES('testuser2', 'test2');
INSERT INTO users(username, password)
VALUES('testuser3', 'test3');