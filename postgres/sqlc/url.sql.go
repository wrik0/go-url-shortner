// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: url.sql

package postgres

import (
	"context"
)

const addUrl = `-- name: AddUrl :one
INSERT INTO urls (
    url, hash
) VALUES (
    $1, $2
)
RETURNING id, url, hash, created_at, updated_at
`

type AddUrlParams struct {
	Url  string `json:"url"`
	Hash string `json:"hash"`
}

func (q *Queries) AddUrl(ctx context.Context, arg AddUrlParams) (Url, error) {
	row := q.queryRow(ctx, q.addUrlStmt, addUrl, arg.Url, arg.Hash)
	var i Url
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.Hash,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUrl = `-- name: GetUrl :one
SELECT id, url, hash, created_at, updated_at FROM urls
WHERE hash = $1
LIMIT 1
`

func (q *Queries) GetUrl(ctx context.Context, hash string) (Url, error) {
	row := q.queryRow(ctx, q.getUrlStmt, getUrl, hash)
	var i Url
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.Hash,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUrlHash = `-- name: GetUrlHash :one
SELECT id, url, hash, created_at, updated_at FROM urls
WHERE url = $1
LIMIT 1
`

func (q *Queries) GetUrlHash(ctx context.Context, url string) (Url, error) {
	row := q.queryRow(ctx, q.getUrlHashStmt, getUrlHash, url)
	var i Url
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.Hash,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
