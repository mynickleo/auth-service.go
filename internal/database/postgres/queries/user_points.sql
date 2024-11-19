-- name: CreateUserPoint :exec
INSERT INTO user_points (user_id, points)
VALUES ($1, $2);

-- name: GetUserPoints :many
SELECT 
    up.id AS user_points_id, 
    u.id AS user_id,
    u.full_name AS user_full_name,
    up.points AS user_points_count 
FROM user_points up
LEFT JOIN users u ON u.id = up.user_id;

-- name: GetUserPointByUserID :one
SELECT 
    up.id AS user_points_id, 
    u.id AS user_id,
    u.full_name AS user_full_name,
    up.points AS user_points_count 
FROM user_points up
LEFT JOIN users u ON u.id = up.user_id
WHERE up.user_id = $1;

-- name: UpdateUserPointsByUserId :exec
UPDATE user_points
SET points = points + $1
WHERE user_id = $2;

-- name: DeleteUserPoints :exec
DELETE FROM user_points
WHERE id = $1;