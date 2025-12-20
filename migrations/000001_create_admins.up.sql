CREATE TYPE admin_role AS ENUM ('super_admin', 'admin');

CREATE TABLE admins
(
    id            SERIAL PRIMARY KEY,
    username      VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255)        NOT NULL,
    role          admin_role          NOT NULL DEFAULT 'admin',
    created_at    TIMESTAMP                    DEFAULT NOW(),
    created_by    INTEGER REFERENCES admins (id)
);

CREATE INDEX idx_admins_username ON admins (username);