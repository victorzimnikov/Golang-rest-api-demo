-- +goose Up
CREATE UNIQUE INDEX issue_title_unique_idx
  ON issues (lower(btrim(title)));

-- +goose Down
DROP INDEX issue_title_unique_idx;