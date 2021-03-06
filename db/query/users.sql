-- name: CreateUser :one
INSERT INTO users (
  id,
  full_name,
  email,
  password,
  contact,
  dog,
  address,
  city,
  post_code,
  longitude,
  latitude
) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11 ) RETURNING *;


-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: GetWalkers :many
SELECT * FROM users ORDER BY name;

-- name: UpdateUser :one
UPDATE users SET
contact=$2,
address=$3,
city=$4,
post_code=$5,
dog=$6
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id=$1;