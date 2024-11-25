-- name: CreateGrid :one
INSERT INTO grids
    (id, created_at, updated_at, grid)
    values($1, $2, $3, $4)
RETURNING *;

-- name: DeleteAllGrids :exec
DELETE FROM grids;

-- name: GetGridById :one
SELECT * FROM grids WHERE id=$1;

-- name: DeleteGridById :exec
DELETE FROM grids WHERE id=$1;