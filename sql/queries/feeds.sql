-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeed :one
SELECT *
FROM feeds
WHERE feeds.url = $1;

-- name: GetFeeds :many
SELECT feeds.name, feeds.url, users.name AS created_by
FROM feeds
JOIN users ON feeds.user_id = users.id;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET updated_at = $1, last_fetched_at = $1
WHERE $2 = id;

-- name: GetNextFeedToFetch :one
SELECT *
FROM feeds
ORDER BY last_fetched_at NULLS FIRST
LIMIT 1;
