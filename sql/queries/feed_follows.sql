-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
        VALUES (
            $1,
            $2,
            $3,
            $4,
            $5
        ) RETURNING *
)

SELECT inserted_feed_follow.*, users.name AS user_name, feeds.name as feed_name
    FROM inserted_feed_follow
    JOIN users ON users.id = inserted_feed_follow.user_id
    JOIN feeds ON feeds.id = inserted_feed_follow.feed_id;

-- name: GetFeedFollowsForUser :many
SELECT feeds.name AS feed_name, users.name AS user_name
FROM feed_follows
JOIN users ON users.id = feed_follows.user_id
JOIN feeds ON feeds.id = feed_follows.feed_id
WHERE users.id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
WHERE $1 = feed_follows.user_id AND $2 = feed_follows.feed_id;