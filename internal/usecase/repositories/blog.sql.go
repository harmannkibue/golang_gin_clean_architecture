// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: blog.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createBlog = `-- name: CreateBlog :one
INSERT INTO blog (
    descriptions
) VALUES (
             $1
         )
    RETURNING id, descriptions, user_role, created_at, updated_at
`

func (q *Queries) CreateBlog(ctx context.Context, descriptions sql.NullString) (Blog, error) {
	row := q.db.QueryRowContext(ctx, createBlog, descriptions)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.Descriptions,
		&i.UserRole,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteBlog = `-- name: DeleteBlog :exec
DELETE FROM blog WHERE id = $1
`

func (q *Queries) DeleteBlog(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteBlog, id)
	return err
}

const getBlog = `-- name: GetBlog :one
SELECT id, descriptions, user_role, created_at, updated_at FROM blog
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetBlog(ctx context.Context, id uuid.UUID) (Blog, error) {
	row := q.db.QueryRowContext(ctx, getBlog, id)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.Descriptions,
		&i.UserRole,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listBlog = `-- name: ListBlog :many
SELECT id, descriptions, user_role, created_at, updated_at FROM blog
ORDER BY created_at
    LIMIT $1
OFFSET $2
`

type ListBlogParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListBlog(ctx context.Context, arg ListBlogParams) ([]Blog, error) {
	rows, err := q.db.QueryContext(ctx, listBlog, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Blog{}
	for rows.Next() {
		var i Blog
		if err := rows.Scan(
			&i.ID,
			&i.Descriptions,
			&i.UserRole,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
