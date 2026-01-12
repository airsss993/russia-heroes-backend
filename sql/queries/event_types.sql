-- name: GetEventTypeByID :one
SELECT *
FROM event_types
WHERE id = $1
LIMIT 1;

-- name: GetEventTypeByName :one
SELECT *
FROM event_types
WHERE name = $1
LIMIT 1;

-- name: ListEventTypes :many
SELECT *
FROM event_types
ORDER BY name;

-- name: CreateEventType :one
INSERT INTO event_types (name, color)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateEventType :exec
UPDATE event_types
SET name  = $2,
    color = $3
WHERE id = $1;

-- name: DeleteEventType :exec
DELETE
FROM event_types
WHERE id = $1;