CREATE TYPE admin AS ENUM ('super_admin', 'admin');
CREATE TABLE admins(
    id uuid,
    username varchar(50),
    pass_hash varchar(200),
    role admin,
    created_at timestamp default now(),
    created_by varchar(50)
);
CREATE INDEX index_admin_id ON admins (id);