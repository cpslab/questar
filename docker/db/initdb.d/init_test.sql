CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    username VARCHAR(20) NOT NULL,
    password VARCHAR(30) NOT NULL
);
INSERT INTO users
VALUES(1, testuser1, test1);
INSERT INTO users
VALUES(2, testuser2, test2);
INSERT INTO users
VALUES(3, testuser3, test3);