-- name: CreateProject :one
INSERT INTO
  projects (
    name,
    description,
    created_at,
    updated_at
  )
VALUES
  ($1, $2, $3, $4) RETURNING id,
  name,
  description,
  created_at,
  updated_at;