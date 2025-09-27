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

-- name: GetAllFeeds :many
SELECT name, url, (
        SELECT name
        FROM users
        WHERE id = user_id
    ) AS created_by
FROM feeds;

-- name: GetFeed :one
SELECT * FROM feeds
    WHERE url = $1;