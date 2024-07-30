-- name: FindById :one
SELECT 
    *
FROM
    users
WHERE 
    id = $1;

-- name: FindAll :many
SELECT
    *
FROM
    users;

-- name: Create :one
INSERT INTO
    users (username, phone, email, balance)
VALUES
    ($1, $2, $3, $4)
RETURNING
    *;

-- name: Update :exec
UPDATE
    users
SET
    username = $2,
    phone = $3,
    email = $4,
    balance = $5
WHERE 
    id = $1
RETURNING
    *;

-- name: DeleteById :exec
DELETE FROM 
    users
WHERE
    id = $1;