CREATE TABLE IF NOT EXISTS roles
(
    id           SERIAL      NOT NULL UNIQUE,
    name         VARCHAR(256) NOT NULL UNIQUE,
    display_name VARCHAR(256),
    description  TEXT
)