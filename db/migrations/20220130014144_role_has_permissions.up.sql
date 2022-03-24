CREATE TABLE IF NOT EXISTS role_has_permissions
(
    role_id       INT NOT NULL REFERENCES roles (id) ON DELETE CASCADE,
    permission_id INT NOT NULL REFERENCES permissions (id) ON DELETE CASCADE,

    PRIMARY KEY (permission_id, role_id)
)