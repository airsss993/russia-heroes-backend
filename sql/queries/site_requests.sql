-- name: GetSiteRequestByID :one
SELECT *
FROM site_requests
WHERE id = $1
LIMIT 1;

-- name: ListSiteRequests :many
SELECT *
FROM site_requests
ORDER BY submitted_at DESC;

-- name: ListSiteRequestsByStatus :many
SELECT *
FROM site_requests
WHERE status = $1
ORDER BY submitted_at DESC;

-- name: ListSiteRequestsByEmail :many
SELECT *
FROM site_requests
WHERE user_email = $1
ORDER BY submitted_at DESC;

-- name: ListSiteRequestsByEventType :many
SELECT *
FROM site_requests
WHERE event_type_id = $1
ORDER BY submitted_at DESC;

-- name: CreateSiteRequest :one
INSERT INTO site_requests (user_email, user_name, site_title, site_description, location, event_date, event_type_id,
                            archive_path, extracted_path, preview_image_url)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: UpdateSiteRequest :exec
UPDATE site_requests
SET user_email        = $2,
    user_name         = $3,
    site_title        = $4,
    site_description  = $5,
    location          = $6,
    event_date        = $7,
    event_type_id     = $8,
    archive_path      = $9,
    extracted_path    = $10,
    preview_image_url = $11
WHERE id = $1;

-- name: UpdateSiteRequestStatus :exec
UPDATE site_requests
SET status          = $2,
    rejection_reason = $3,
    reviewed_at     = $4,
    reviewed_by     = $5
WHERE id = $1;

-- name: ApproveSiteRequest :exec
UPDATE site_requests
SET status      = 'approved',
    reviewed_at = NOW(),
    reviewed_by = $2
WHERE id = $1;

-- name: RejectSiteRequest :exec
UPDATE site_requests
SET status           = 'rejected',
    rejection_reason = $2,
    reviewed_at      = NOW(),
    reviewed_by      = $3
WHERE id = $1;

-- name: DeleteSiteRequest :exec
DELETE
FROM site_requests
WHERE id = $1;

-- name: CountSiteRequestsByStatus :one
SELECT COUNT(*)
FROM site_requests
WHERE status = $1;

-- name: GetPendingSiteRequests :many
SELECT *
FROM site_requests
WHERE status = 'pending'
ORDER BY submitted_at ASC;

-- name: GetSiteRequestsWithinRadius :many
SELECT *,
       ST_Distance(location, ST_SetSRID(ST_MakePoint($1, $2), 4326)) as distance
FROM site_requests
WHERE ST_DWithin(location, ST_SetSRID(ST_MakePoint($1, $2), 4326), $3)
ORDER BY distance;