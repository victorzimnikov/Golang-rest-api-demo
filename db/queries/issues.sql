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