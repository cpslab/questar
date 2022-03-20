CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    -- 後でランダムな数値に設定する
    user_name VARCHAR(20) NOT NULL,
    password VARCHAR(30) NOT NULL
);
INSERT INTO users(user_name, password)
VALUES('testuser1', 'test1');
INSERT INTO users(user_name, password)
VALUES('testuser2', 'test2');
INSERT INTO users(user_name, password)
VALUES('testuser3', 'test3');