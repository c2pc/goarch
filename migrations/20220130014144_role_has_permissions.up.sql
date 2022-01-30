CREATE TABLE IF NOT EXISTS role_has_permissions
(
    permission_id INT NOT NULL REFERENCES permissions (id) ON DELETE CASCADE,
    role_id       INT REFERENCES roles (id) ON DELETE CASCADE,

    PRIMARY KEY (permission_id, role_id)
)