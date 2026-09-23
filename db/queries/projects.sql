-- name: CreateProject :one
INSERT INTO
  projects (
    name,
    description
  )
VALUES
  (
    sqlc.arg('name'),
    sqlc.arg('description')
  )
RETURNING
  id,
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

-- name: DeleteProject :one
DELETE FROM projects
WHERE id = $1
RETURNING id;

-- name: UpdateProject :one
UPDATE projects
SET
  name = COALESCE(sqlc.narg('name'), projects.name),
  description = COALESCE(sqlc.narg('description'), projects.description),
  updated_at = NOW()
WHERE id = sqlc.arg('id')
RETURNING id, name, description, created_at, updated_at;
