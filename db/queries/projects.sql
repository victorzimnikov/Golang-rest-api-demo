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

-- name: GetProject :one
SELECT
  id,
  name,
  description,
  created_at,
  updated_at
FROM projects
WHERE id = $1;