CREATE TABLE IF NOT EXISTS users (
    id          uuid        PRIMARY KEY DEFAULT gen_random_uuid(),
    username    VARCHAR(50) NOT NULL UNIQUE,
    fullname    VARCHAR(150),
    password    TEXT,
    email       VARCHAR(150) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS sessions_manager (
    user_id uuid NOT NULL,
    user_name    VARCHAR(50) NOT NULL ,
    user_agent VARCHAR(300) NOT NULL,
    token   VARCHAR(300) NOT NULL,
    CONSTRAINT users_session_fk FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT users_name_session_fk FOREIGN KEY(user_name) REFERENCES users(username) ON DELETE CASCADE ON UPDATE CASCADE
);