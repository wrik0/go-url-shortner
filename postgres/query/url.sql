-- name: AddUrl :one
INSERT INTO urls (
    url, hash
) VALUES (
    $1, $2
)
RETURNING *;

-- name: GetUrl :one
SELECT * FROM urls
WHERE hash = $1
LIMIT 1;

-- name: GetUrlHash :one
SELECT * FROM urls
WHERE url = $1
LIMIT 1;

