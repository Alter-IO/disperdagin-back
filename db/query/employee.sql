-- name: FindEmployeeByID :one
SELECT
    id,
    name,
    position,
    address,
    employee_id,
    birthplace,
    birthdate,
    photo,
    status,
    author,
    created_at,
    updated_at
FROM
    employees
WHERE
    id = $1
AND
    deleted_at IS NULL;

-- name: FindAllEmployees :many
SELECT
    id,
    name,
    position,
    employee_id,
    status,
    author,
    created_at
FROM
    employees
WHERE
    deleted_at IS NULL
ORDER BY
    name ASC;

-- name: FindEmployeesByPosition :many
SELECT
    id,
    name,
    position,
    employee_id,
    status,
    author,
    created_at
FROM
    employees
WHERE
    position = $1
AND
    deleted_at IS NULL
ORDER BY
    name ASC;

-- name: FindActiveEmployees :many
SELECT
    id,
    name,
    position,
    employee_id,
    author,
    created_at
FROM
    employees
WHERE
    status = 1
AND
    deleted_at IS NULL
ORDER BY
    name ASC;

-- name: InsertEmployee :exec
INSERT INTO employees(id, name, position, address, employee_id, birthplace, birthdate, photo, status, author, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);

-- name: UpdateEmployee :execrows
UPDATE
    employees
SET
    name = sqlc.arg(name),
    position = sqlc.arg(position),
    address = sqlc.arg(address),
    employee_id = sqlc.arg(employee_id),
    birthplace = sqlc.arg(birthplace),
    birthdate = sqlc.arg(birthdate),
    photo = sqlc.arg(photo),
    status = sqlc.arg(status),
    updated_at = sqlc.arg(updated_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;

-- name: DeleteEmployee :execrows
UPDATE
    employees
SET
    deleted_at = sqlc.arg(deleted_at)
WHERE
    id = sqlc.arg(id)
AND
    deleted_at IS NULL;