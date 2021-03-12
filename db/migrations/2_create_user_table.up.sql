SET
statement_timeout = 60000;
SET
lock_timeout = 60000;

CREATE TABLE user_credentials
(
    id          SERIAL PRIMARY KEY,
    username    VARCHAR(255) NOT NULL,
    password    VARCHAR(255) NOT NULL
);
