-- +goose Up
CREATE UNIQUE INDEX project_name_unique_idx
  ON projects (lower(btrim(name)));

-- +goose Down
DROP INDEX project_name_unique_idx;