SET
statement_timeout = 60000; -- 60 seconds
SET
lock_timeout = 60000; -- 60 seconds

--gopg:split

CREATE TABLE notes
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    description TEXT,
    is_checked  BOOLEAN DEFAULT true
);
