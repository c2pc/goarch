CREATE TABLE IF NOT EXISTS role_has_users
(
    user_id INT NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    role_id       INT REFERENCES roles (id) ON DELETE RESTRICT ,

    PRIMARY KEY (user_id, role_id)
)