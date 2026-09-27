-- name: CreateComment :one
WITH inserted AS (
  INSERT INTO comments (
    issue_id,
    text
  )
  SELECT
    sqlc.arg('issue_id'),
    sqlc.arg('text')
  WHERE EXISTS (
    SELECT 1
    FROM issues
    WHERE id = sqlc.arg('issue_id')
  )
  RETURNING
    id,
    issue_id,
    text,
    created_at,
    updated_at
)
SELECT
  i.id,
  i.text,
  i.created_at,
  i.updated_at,
  p.id          AS issue_id,
  p.title        AS issue_title
FROM inserted i
JOIN issues p ON p.id = i.issue_id;