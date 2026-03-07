-- name: CreateChirp :one
INSERT INTO chirps (id, created_at, updated_at, user_id)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1
)
RETURNING *;

-- name: GetChirpByUserId :many
SELECT * FROM chirps WHERE user_id=$1;

-- name: GetChirpById :one
SELECT * FROM chirps WHERE id=$1;

-- name: DeleteChirp :exec
DELETE FROM chirps WHERE id=$1;

-- name: GetChirp :many
SELECT * FROM chirps;
