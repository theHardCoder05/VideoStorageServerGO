-- name: createContent :one
INSERT INTO contents (
  fileid,
  name,
  size,
  created_at
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetContent :one
SELECT * FROM contents
WHERE id = $1 LIMIT 1;

-- name: GetContentUpdate :one
SELECT * FROM contents
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListContents :many
SELECT * FROM contents
WHERE fileid = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateContent :one
UPDATE contents
SET name = $2
WHERE id = $1
RETURNING *;



-- name: DeleteContent :exec
DELETE FROM contents
WHERE id = $1;
