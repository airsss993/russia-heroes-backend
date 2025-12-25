-- TODO: убрать возвращение полных данных из БД, возвращать только id при создании админа

-- name: GetAdminByID :one
SELECT *
FROM admins
WHERE id = $1
LIMIT 1;

-- name: GetAdminByUsername :one
SELECT *
FROM admins
WHERE user_name = $1
LIMIT 1;

-- name: ListAdmins :many
SELECT *
FROM admins
ORDER BY user_name;

-- name: CreateAdmin :one
INSERT INTO admins (user_name, password_hash, role, created_by)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: CreateSuperAdmin :one
INSERT INTO admins (user_name, password_hash, role, created_by)
VALUES ($1, $2, 'super-admin', $3)
RETURNING *;

-- name: DeleteAdmin :exec
DELETE
FROM admins
WHERE id = $1;