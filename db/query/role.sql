-- name: FindRoles :many
SELECT
    id,
    name
FROM
    roles
WHERE
    id != 'superadmin';