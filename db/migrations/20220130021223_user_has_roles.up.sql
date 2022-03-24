CREATE TABLE IF NOT EXISTS user_has_roles
(
    user_id INT NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    role_id INT NOT NULL REFERENCES roles (id) ON DELETE RESTRICT,

    PRIMARY KEY (user_id, role_id)
)