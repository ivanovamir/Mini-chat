BEGIN;

CREATE TABLE IF NOT EXISTS "user" (
    id SERIAL PRIMARY KEY,
    login VARCHAR(255) NOT NULL UNIQUE ,
    name VARCHAR(255) NOT NULL,
    ip CIDR NOT NULL,
    last_visit TIMESTAMP NOT NULL
);

CREATE INDEX login_usr_idx ON "user"(login);

CREATE TABLE IF NOT EXISTS message (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES "user" (id),
    message TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX user_id_msg_idx ON message(user_id);

COMMIT;