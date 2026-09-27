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

-- name: GetIssueCommentsList :many
SELECT
  id,
  text
FROM comments
WHERE
  issue_id = sqlc.arg(issue_id)
  AND (
    sqlc.arg(q)::text = ''
    OR text ILIKE '%' || sqlc.arg(q)::text || '%'
  )
ORDER BY created_at DESC, id DESC
LIMIT sqlc.arg(limit_count)
OFFSET sqlc.arg(skip)::bigint;

-- name: CountIssueComments :one
SELECT COUNT(c.id)
FROM issues i
LEFT JOIN comments c
  ON c.issue_id = i.id
  AND (
    sqlc.arg(q)::text = ''
    OR c.text ILIKE '%' || sqlc.arg(q)::text || '%'
  )
WHERE i.id = sqlc.arg(issue_id)
GROUP BY i.id;

-- name: GetComment :one
SELECT
  c.id,
  c.text,
  c.created_at,
  c.updated_at,
  i.id AS issue_id,
  i.title AS issue_title
FROM comments c
JOIN issues i ON i.id = c.issue_id
WHERE c.id = sqlc.arg(comment_id);