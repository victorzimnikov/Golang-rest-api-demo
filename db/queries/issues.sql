-- name: CreateIssue :one
WITH inserted AS (
  INSERT INTO issues (
    project_id,
    title,
    description,
    status,
    priority,
    due_date
  )
  SELECT
    sqlc.arg('project_id'),
    sqlc.arg('title'),
    sqlc.arg('description'),
    sqlc.arg('status'),
    sqlc.arg('priority'),
    sqlc.arg('due_date')
  WHERE EXISTS (
    SELECT 1
    FROM projects
    WHERE id = sqlc.arg('project_id')
  )
  RETURNING
    id,
    project_id,
    title,
    description,
    status,
    priority,
    due_date,
    created_at,
    updated_at
)
SELECT
  i.id,
  i.title,
  i.description,
  i.status,
  i.priority,
  i.due_date,
  i.created_at,
  i.updated_at,
  p.id          AS project_id,
  p.name        AS project_name
FROM inserted i
JOIN projects p ON p.id = i.project_id;

-- name: GetProjectIssuesList :many
SELECT
  id,
  title,
  status,
  priority
FROM issues
WHERE
  project_id = sqlc.arg(project_id)
  AND (
    sqlc.arg(q)::text = ''
    OR title ILIKE '%' || sqlc.arg(q)::text || '%'
    OR description ILIKE '%' || sqlc.arg(q)::text || '%'
  )
  AND (
    sqlc.arg(status)::text = ''
    OR status = sqlc.arg(status)::text
  )
  AND (
    sqlc.arg(priority)::text = ''
    OR priority = sqlc.arg(priority)::text
  )
ORDER BY created_at DESC, id DESC
LIMIT sqlc.arg(limit_count)
OFFSET sqlc.arg(skip)::bigint;

-- name: CountProjectIssues :one
SELECT COUNT(*)
FROM issues
WHERE
  project_id = sqlc.arg(project_id)
  AND (
    sqlc.arg(q)::text = ''
    OR title ILIKE '%' || sqlc.arg(q)::text || '%'
    OR description ILIKE '%' || sqlc.arg(q)::text || '%'
  )
  AND (
    sqlc.arg(status)::text = ''
    OR status = sqlc.arg(status)::text
  )
  AND (
    sqlc.arg(priority)::text = ''
    OR priority = sqlc.arg(priority)::text
  );

-- name: GetIssue :one
SELECT
  i.id,
  i.title,
  i.description,
  i.status,
  i.priority,
  i.created_at,
  i.updated_at,
  i.due_date,
  p.id AS project_id,
  p.name AS project_name
FROM issues i
JOIN projects p ON p.id = i.project_id
WHERE i.id = sqlc.arg(issue_id);

-- name: DeleteIssue :one
DELETE FROM issues
WHERE id = sqlc.arg(issue_id)
RETURNING id;

-- name: UpdateIssue :one
UPDATE issues AS i
SET
  title = COALESCE(sqlc.narg('title'), i.title),
  description = COALESCE(sqlc.narg('description'), i.description),
  status = COALESCE(sqlc.narg('status'), i.status),
  priority = COALESCE(sqlc.narg('priority'), i.priority),
  due_date = CASE
    WHEN sqlc.arg('due_date_set')::boolean
      THEN sqlc.narg('due_date')::date
    ELSE i.due_date
  END,
  updated_at = NOW()
FROM projects AS p
WHERE i.id = sqlc.arg('id') AND p.id = i.project_id
RETURNING
  i.id,
  i.title,
  i.description,
  i.status,
  i.priority,
  i.due_date,
  i.created_at,
  i.updated_at,
  p.id          AS project_id,
  p.name        AS project_name;
