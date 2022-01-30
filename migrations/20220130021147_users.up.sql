CREATE TABLE IF NOT EXISTS users
(
    id                SERIAL       NOT NULL UNIQUE,
    username          VARCHAR(256) NOT NULL UNIQUE,
    name              VARCHAR(256) NOT NULL,
    password          VARCHAR(256) NOT NULL,
    email             VARCHAR(256) NOT NULL,
    token             VARCHAR(256),
    email_verified_at timestamp    NULL,
    updated_at        timestamp DEFAULT now(),
    created_at        timestamp DEFAULT now()
)