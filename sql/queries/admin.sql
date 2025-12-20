-- name: GetAdminByID :one
SELECT *
FROM admins
WHERE id = $1
LIMIT 1;

-- name: GetAdminByUsername :one
SELECT *
FROM admins
WHERE username = $1
LIMIT 1;

-- name: ListAdmins :many
SELECT *
FROM admins
ORDER BY username;

-- name: CreateAdmin :one
INSERT INTO admins (username, password_hash, role, created_by)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateAdmin :exec
UPDATE admins
SET username = $2,
    password_hash = $3,
    role = $4
WHERE id = $1;

-- name: DeleteAdmin :exec
DELETE
FROM admins
WHERE id = $1;