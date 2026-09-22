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

-- name: GetProjectsList :many
SELECT
  id,
  name,
  description
FROM projects
WHERE (
  name ILIKE '%' || sqlc.arg(q)::text || '%'
  OR description ILIKE '%' || sqlc.arg(q)::text || '%'
)
ORDER BY created_at DESC, id DESC
LIMIT sqlc.arg(limit_count)
OFFSET sqlc.arg(skip)::bigint;

-- name: CountProjects :one
SELECT COUNT(*)
FROM projects
WHERE (
  name ILIKE '%' || sqlc.arg(q)::text || '%'
  OR description ILIKE '%' || sqlc.arg(q)::text || '%'
);